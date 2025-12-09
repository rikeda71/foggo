package basic

// ServerConfig demonstrates basic usage of functional options pattern
// Run: foggo fop -struct=ServerConfig -output=config_gen.go config.go
type ServerConfig struct {
	// Host is the server hostname
	Host string `fog:"default:localhost"`

	// Port is the server port
	Port int `fog:"default:8080"`

	// TLS enables HTTPS
	TLS bool `fog:"default:false"`

	// Timeout in seconds
	Timeout int `fog:"default:30"`
}

// DatabaseConfig shows more complex configuration
type DatabaseConfig struct {
	// DSN is the database connection string
	DSN string `fog:"required"`

	// MaxConnections limits concurrent connections
	MaxConnections int `fog:"default:10"`

	// MaxIdleTime in seconds
	MaxIdleTime int `fog:"default:300"`

	// EnableLogging turns on query logging
	EnableLogging bool
}