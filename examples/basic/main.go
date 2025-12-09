package basic

import (
	"fmt"
	"log"
)

// Example usage after running:
// foggo fop -struct=ServerConfig -output=config_gen.go config.go
// foggo fop -struct=DatabaseConfig -output=config_gen.go config.go

func ExampleUsage() {
	// Basic usage with default values
	server, err := NewServerConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server: %+v\n", server)

	// Custom configuration
	customServer, err := NewServerConfig(
		WithHost("api.example.com"),
		WithPort(443),
		WithTLS(true),
		WithTimeout(60),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Custom Server: %+v\n", customServer)

	// Database config with required field
	db, err := NewDatabaseConfig(
		WithDSN("postgres://user:pass@localhost/db"),
		WithMaxConnections(50),
		WithEnableLogging(true),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Database: %+v\n", db)
}