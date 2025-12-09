# Advanced Example with AFOP

This example demonstrates the Advanced Functional Options Pattern (AFOP) which includes context support.

## Files

- `client.go` - HTTP client struct with various configuration options
- `main.go` - Examples of using AFOP with context and custom option builders

## Usage

1. Generate the AFOP code:
```bash
foggo afop -struct=HTTPClient -output=client_gen.go client.go
```

2. Use the generated code with context:
```go
ctx := context.Background()

client, err := NewHTTPClient(
    ctx,
    WithBaseURL("https://api.example.com"),
    WithTimeout(60*time.Second),
    WithBearerToken("token"),
)
```

## Features Demonstrated

- Context support in AFOP
- Complex types: time.Duration, map[string]string
- Custom option builders for common configurations
- Chaining multiple options
- Production vs development configurations
- Authentication setup
- Debug mode configuration

## Best Practices

1. **Use custom option builders** for common configurations:
   ```go
   func ConfigureForProduction() HTTPClientOption {
       // Return a function that sets multiple options
   }
   ```

2. **Context cancellation** can be used to abort initialization:
   ```go
   ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
   defer cancel()
   client, err := NewHTTPClient(ctx, options...)
   ```

3. **Validate in option functions** when needed:
   ```go
   func WithCustomOption(value string) HTTPClientOption {
       return func(ctx context.Context, c *HTTPClient) error {
           if value == "" {
               return fmt.Errorf("value cannot be empty")
           }
           c.Value = value
           return nil
       }
   }
   ```