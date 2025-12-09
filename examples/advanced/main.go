package advanced

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Example usage after running:
// foggo afop -struct=HTTPClient -output=client_gen.go client.go

func ExampleAFOPUsage() {
	ctx := context.Background()

	// Basic client with required field
	basicClient, err := NewHTTPClient(
		ctx,
		WithBaseURL("https://api.example.com"),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Basic Client: %+v\n", basicClient)

	// Advanced client with multiple options
	headers := map[string]string{
		"X-API-Key": "secret",
		"Accept":    "application/json",
	}

	advancedClient, err := NewHTTPClient(
		ctx,
		WithBaseURL("https://api.production.com"),
		WithHeaders(headers),
		WithTimeout(60*time.Second),
		WithRetryCount(5),
		WithRetryDelay(2*time.Second),
		WithUserAgent("MyApp/2.0"),
		WithMaxIdleConns(200),
		WithMaxConnsPerHost(20),
		WithBearerToken("jwt-token-here"),
		WithEnableDebug(true),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Advanced Client: %+v\n", advancedClient)

	// Client for development with proxy
	devClient, err := NewHTTPClient(
		ctx,
		WithBaseURL("https://api.dev.com"),
		WithProxyURL("http://localhost:8888"),
		WithTLSSkipVerify(true), // Only for development!
		WithEnableDebug(true),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Dev Client: %+v\n", devClient)
}

// ExampleClientBuilder shows how to create a builder pattern
func ExampleClientBuilder() {
	ctx := context.Background()

	// Chain multiple configurations
	client, err := NewHTTPClient(
		ctx,
		WithBaseURL("https://api.example.com"),
		ConfigureForProduction(),
		ConfigureAuthentication("my-token"),
		ConfigureDebugging(true),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Built Client: %+v\n", client)
}

// ConfigureForProduction returns a set of production-ready options
func ConfigureForProduction() HTTPClientOption {
	return func(ctx context.Context, c *HTTPClient) error {
		c.Timeout = 30 * time.Second
		c.RetryCount = 3
		c.MaxIdleConns = 100
		c.MaxConnsPerHost = 10
		c.TLSSkipVerify = false
		return nil
	}
}

// ConfigureAuthentication sets up authentication
func ConfigureAuthentication(token string) HTTPClientOption {
	return func(ctx context.Context, c *HTTPClient) error {
		c.BearerToken = token
		if c.Headers == nil {
			c.Headers = make(map[string]string)
		}
		c.Headers["Authorization"] = "Bearer " + token
		return nil
	}
}

// ConfigureDebugging enables debugging features
func ConfigureDebugging(enabled bool) HTTPClientOption {
	return func(ctx context.Context, c *HTTPClient) error {
		c.EnableDebug = enabled
		if enabled {
			c.Headers["X-Debug"] = "true"
		}
		return nil
	}
}