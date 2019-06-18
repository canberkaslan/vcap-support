// Package vcap provides structures and functions to simplify
// access to Cloud Foundry VCAP (VMware Cloud Application Platform)
// environment variables.
package vcap

import (
	"encoding/json"
	"fmt"
	"os"
)

// AppKey represents the Cloud Foundry application environment variable
const AppKey = "VCAP_APPLICATION"

// ServiceKey represents the Cloud Foundry services environment variable
const ServiceKey = "VCAP_SERVICES"

// Application is a struct representing VCAP_APPLICATION. We are ignoring
// the following deprecated and/or redundant attributes: host,
// instance_index, name, port, started_at, state_timestamp, uris,
// users, version
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

// Services is a map that can hold all the service configurations
// from VCAP_SERVICES. Each key (bound service name) maps to a slice
// of service configurations, one per bound instance of that service.
type Services map[string][]Service

// Service is a structure that holds configuration for a single
// service instance. The content and structure of the credentials
// is service dependent, so is just handled as a map of keys to
// values of unknown type.
type Service struct {
	Credentials map[string]interface{} `json:"credentials"`
	Label       string                 `json:"label"`
	Name        string                 `json:"name"`
	Plan        string                 `json:"plan"`
	Tags        []string               `json:"tags"`
}

// LoadServices loads VCAP services configuration
func LoadServices() (Services, error) {

	var svc Services

	j := os.Getenv(ServiceKey)
	if err := json.Unmarshal([]byte(j), &svc); err != nil {
		return nil, err
	}

	return svc, nil
}

// LoadApplication loads VCAP application configuration
func LoadApplication() (Application, error) {

	var app Application

	j := os.Getenv(AppKey)
	if err := json.Unmarshal([]byte(j), &app); err != nil {
		return Application{}, err
	}

	return app, nil
}

// GetCredential searches for the credential key with the given service name
// and returns the credential's value.
func (s Services) GetCredential(name string, key string) (string, error) {
	svc := s.FindServiceByName(name)
	if len(svc) == 0 {
		svc = s.FindServiceByLabel(name)
		if len(svc) == 0 {
			return "", fmt.Errorf("could not find service with name or label %s", name)
		}
	}

	value, ok := svc[0].Credentials[key].(interface{})
	if !ok {
		return "", fmt.Errorf("could not read credential %s.%s", name, key)
	}

	return fmt.Sprint(value), nil
}

// FindServiceByName searches through all entries under each service
// key to find entries with a given name.
//
// Returns a slice of Service entries with matching name
func (s Services) FindServiceByName(name string) []Service {

	svc := make([]Service, 0)

	for _, v := range s {
		for _, entry := range v {
			if entry.Name == name {
				svc = append(svc, entry)
			}
		}
	}

	return svc
}

// FindServiceByLabel searches through all entries under each service
// key to find entries with a given label.
//
// Returns a slice of Service entries with matching label
func (s Services) FindServiceByLabel(label string) []Service {

	svc := make([]Service, 0)

	for _, v := range s {
		for _, entry := range v {
			if entry.Label == label {
				svc = append(svc, entry)
			}
		}
	}

	return svc
}
