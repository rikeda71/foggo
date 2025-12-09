# Basic Example

This example demonstrates the basic usage of Foggo for generating functional options pattern.

## Files

- `config.go` - Contains the struct definitions with fog tags
- `main.go` - Shows how to use the generated code

## Usage

1. Generate the functional options:
```bash
# For ServerConfig
foggo fop -struct=ServerConfig -output=config_gen.go config.go

# For DatabaseConfig
foggo fop -struct=DatabaseConfig -output=config_gen.go config.go
```

2. Use the generated code:
```go
// Create with defaults
server, err := NewServerConfig()

// Create with custom options
server, err := NewServerConfig(
    WithHost("api.example.com"),
    WithPort(443),
    WithTLS(true),
)
```

## Features Demonstrated

- Default values using `fog:"default:value"`
- Required fields using `fog:"required"`
- Various types: string, int, bool
- Error handling for required fields