# vcap
--
    import "github.build.ge.com/aviation-predix-common/vcap-support"

Package vcap provides structures and functions to simplify access to Cloud
Foundry VCAP (VMware Cloud Application Platform) environment variables.

## Usage

```go
const AppKey = "VCAP_APPLICATION"
```
Cloud Foundry application environment variable

```go
const ServiceKey = "VCAP_SERVICES"
```
Cloud Foundry services environment variable

#### type Application

```go
type Application struct {
	ID         string   `json:"application_id"`
	Name       string   `json:"application_name"`
	URIs       []string `json:"application_uris"`
	Version    string   `json:"application_version"`
	InstanceID string   `json:"instance_id"`
	Limits     struct {
		Mem  int `json:"mem"`
		Disk int `json:"disk"`
		FDs  int `json:"fds"`
	} `json:"limits"`
	SpaceID            string `json:"space_id"`
	Start              string `json:"start"`
	StartedAtTimestamp int64  `json:"started_at_timestamp"`
}
```

App is a struct representing VCAP_APPLICATION. We are ignoring the following
deprecated and/or redundant attributes: host, instance_index, name, port,
started_at, state_timestamp, uris, users, version

#### func  LoadApplication

```go
func LoadApplication() (Application, error)
```
LoadApplication loads VCAP application configuration

#### type Service

```go
type Service struct {
	Credentials map[string]interface{} `json:"credentials"`
	Label       string                 `json:"label"`
	Name        string                 `json:"name"`
	Plan        string                 `json:"plan"`
	Tags        []string               `json:"tags"`
}
```

Service is a structure that holds configuration for a single service instance.
The content and structure of the credentials is service dependent, so is just
handled as a map of keys to values of unknown type.

#### type Services

```go
type Services map[string][]Service
```

Services is a map that can hold all the service configurations from
VCAP_SERVICES. Each key (bound service name) maps to a slice of service
configurations, one per bound instance of that service.

#### func  LoadServices

```go
func LoadServices() (Services, error)
```
LoadServices loads VCAP services configuration

#### func (Services) GetCredential

```go
func (s Services) GetCredential(name string, key string) (string, error)
```
GetCredential searches for the credential key with the given service name and
returns the credential's value.

#### func (Services) FindServiceByLabel

```go
func (s Services) FindServiceByLabel(label string) []Service
```
FindServiceByLabel searches through all entries under each service key to find
entries with a given label.

Returns a slice of Service entries with matching label

#### func (Services) FindServiceByName

```go
func (s Services) FindServiceByName(name string) []Service
```
FindServiceByName searches through all entries under each service key to find
entries with a given name.

Returns a slice of Service entries with matching name

#### type Postgres

```go
type Postgres struct {
	ID         int    `json:"ID"`
	BindingID  string `json:"binding_id"`
	Database   string `json:"database"`
	DSN        string `json:"dsn"`
	Host       string `json:"host"`
	InstanceID string `json:"instance_id"`
	JDBCURI    string `json:"jdbc_uri"`
	Password   string `json:"password"`
	Port       string `json:"port"`
	URI        string `json:"uri"`
	Username   string `json:"username"`
}
```

Postgres is a struct representing a Postgres credentials block

#### type RabbitMQ

```go
type RabbitMQ struct {
	DashboardURL string    `json:"dashboard_url"`
	Hostname     string    `json:"hostname"`
	Hostnames    []string  `json:"hostnames"`
	APIURI       string    `json:"http_api_uri"`
	APIURIs      []string  `json:"http_api_uris"`
	Password     string    `json:"password"`
	Protocols    Protocols `json:"protocols"`
	SSL          bool      `json:"ssl"`
	URI          string    `json:"uri"`
	URIs         []string  `json:"uris"`
	Username     string    `json:"username"`
	VHost        string    `json:"vhost"`
}
```

RabbitMQ is a struct representing a rabbit MQ credentials block

#### type Protocols

```go
type Protocols map[string]Protocol
```

Protocols is a map of Rabbit MQ protocol types to configuration

#### type Protocol

```go
type Protocol struct {
	Host     string   `json:"host"`
	Hosts    []string `json:"hosts"`
	Password string   `json:"password"`
	Path     string   `json:"path,omitempty"` // Not expected for AMQP, MQTT, STOMP
	Port     int      `json:"port"`
	SSL      bool     `json:"ssl"`
	URI      string   `json:"uri"`
	URIs     []string `json:"uris"`
	Username string   `json:"username"`
	VHost    string   `json:"vhost,omitempty"` // Not expected for Management, MQTT
}
```

Protocol contains protocol-specific connection credentials.

#### type UAA

```go
type UAA struct {
	IssuerID  string `json:"ID"`
	Subdomain string `json:"subdomain"`
	URI       string `json:"uri"`
	Zone      struct {
		HeaderName  string `json:"http-header-name"`
		HeaderValue string `json:"http-header-value"`
	} `json:"zone"`
}
```

UAA is a struct representing a UAA credentials block
