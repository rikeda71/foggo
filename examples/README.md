# Foggo Examples

This directory contains examples demonstrating various use cases for Foggo.

## Examples

### 1. [Basic](./basic/)
Simple examples showing the fundamental usage of Foggo with the Functional Options Pattern (FOP).

**Key concepts:**
- Default values with `fog:"default:value"`
- Required fields with `fog:"required"`
- Basic types (string, int, bool)

### 2. [Advanced](./advanced/)
Advanced examples using the Advanced Functional Options Pattern (AFOP) with context support.

**Key concepts:**
- Context support
- Complex types (time.Duration, maps)
- Custom option builders
- Configuration chaining

## Running the Examples

1. Install Foggo:
```bash
go install github.com/rikeda71/foggo@latest
```

2. Navigate to an example directory:
```bash
cd basic
```

3. Generate the code:
```bash
# For FOP
foggo fop -struct=YourStruct -output=generated.go source.go

# For AFOP (with context)
foggo afop -struct=YourStruct -output=generated.go source.go
```

4. Use the generated code in your application.

## Creating Your Own Examples

1. Define a struct with fog tags:
```go
type Config struct {
    Field1 string `fog:"default:value"`
    Field2 int    `fog:"required"`
}
```

2. Run Foggo to generate the functional options
3. Use the generated `NewConfig` function with option functions

## Tips

- Use `fog:"required"` for fields that must be set
- Use `fog:"default:value"` for fields with sensible defaults
- AFOP is recommended when you need context support or async initialization
- Create custom option builders for common configuration patterns