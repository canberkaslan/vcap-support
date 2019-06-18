package vcap

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"
	"testing"
)

var validVcap = []byte(`{
  "postgres": [
    {
      "credentials": {
        "ID": 0,
        "database": "notarealpostgresdatabase",
        "host": "postgres-host.io",
        "password": "notarealpostgrespassword",
        "port": "1234",
        "uri": "postgres://notarealpostgresusername:notarealpostgrespassword@postgres-host.io:1234/notarealpostgresdatabase?sslmode=disable",
        "username": "notarealpostgresusername",
        "boolean": true,
        "nothing": null
      },
      "label": "postgres",
      "name": "my-postgres-name",
      "plan": "shared-nr",
      "provider": null,
      "syslog_drain_url": null,
      "tags": [
        "rdpg",
        "postgresql"
      ],
      "volume_mounts": []
    }
  ],
  "redis-1": [
    {
      "credentials": {
        "host": "10.001.10.001",
        "password": "not-a-real-redis-password",
        "port": 12345
      },
      "label": "redis-1",
      "name": "my-redis-name",
      "plan": "shared-vm",
      "provider": null,
      "syslog_drain_url": null,
      "tags": [
        "pivotal",
        "redis"
      ],
      "volume_mounts": []
    }
  ],
  "user-provided": [
    {
      "credentials": {
        "issuerId": "https://my-uaa.predix.io/oauth/token",
        "uri": "https://my-uaa.predix.io",
        "zone": {
          "http-header-name": "X-Identity-Zone-Id",
          "http-header-value": "a-http-header-value"
        }
      },
      "label": "user-provided",
      "name": "my-uaa-name",
      "syslog_drain_url": "",
      "tags": [],
      "volume_mounts": []
    },
    {
      "credentials": {
        "instanceId": "my-views-instance-id",
        "uri": "https://predix-views.predix.io",
        "float": 50.000001
      },
      "label": "user-provided",
      "name": "my-views-name",
      "plan": "Standard",
      "provider": null,
      "syslog_drain_url": null,
      "tags": [],
      "volume_mounts": []
    }
  ]
}`)

type getCredentialTest struct {
	name        string
	vcap        []byte
	serviceName string
	credKey     string
	expdValue   string
	err         error
}

var getCredentialTests = []getCredentialTest{
	{"UsernameFromLabel",
		validVcap, "postgres", "username", "notarealpostgresusername", nil},
	{"PasswordFromName",
		validVcap, "my-redis-name", "password", "not-a-real-redis-password", nil},
	{"UriFromUserProvidedName",
		validVcap, "my-uaa-name", "uri", "https://my-uaa.predix.io", nil},
	{"DifferentUriFromUserProvidedName",
		validVcap, "my-views-name", "uri", "https://predix-views.predix.io", nil},
	{"anIntValue",
		validVcap, "redis-1", "port", "12345", nil},
	{"aFloatValue",
		validVcap, "my-views-name", "float", "50.000001", nil},
	{"aBoolValue",
		validVcap, "postgres", "boolean", "true", nil},
	{"aNullValue",
		validVcap, "my-postgres-name", "nothing", "", errors.New("could not read credential my-postgres-name.nothing")},
}

func TestGetCredential(t *testing.T) {
	for _, tst := range getCredentialTests {
		svc := loadTestServices(tst.vcap)

		cred, err := svc.GetCredential(tst.serviceName, tst.credKey)
		if !reflect.DeepEqual(err, tst.err) {
			t.Errorf(tst.name+": invalid error return - got \"%v\", expected \"%v\"", err, tst.err)
		}

		if err == nil {
			if cred != tst.expdValue {
				t.Errorf(tst.name+": key mismatch - got %+v, expected %+v", cred, tst.expdValue)
			}
		}
	}
}

func loadTestServices(j []byte) Services {
	var svc Services

	if err := json.Unmarshal(j, &svc); err != nil {
		log.Printf("Unmarshalling error: " + err.Error())
		return nil
	}
	return svc
}
