// Package vcap Service-specific credentials structs
package vcap

// Postgres is a struct representing a Postgres credentials block
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

// UAA is a struct representing a UAA credentials block
type UAA struct {
	IssuerID  string `json:"ID"`
	Subdomain string `json:"subdomain"`
	URI       string `json:"uri"`
	Zone      struct {
		HeaderName  string `json:"http-header-name"`
		HeaderValue string `json:"http-header-value"`
	} `json:"zone"`
}

// RabbitMQ is a struct representing a rabbit MQ credentials block
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

// Protocols is a map of Rabbit MQ protocol types to configuration
type Protocols map[string]Protocol

// Protocol contains protocol-specific connection credentials.
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
