package advanced

import (
	"time"
)

// HTTPClient demonstrates advanced functional options with AFOP
// Run: foggo afop -struct=HTTPClient -output=client_gen.go client.go
type HTTPClient struct {
	// BaseURL is the API base URL
	BaseURL string `fog:"required"`

	// Headers for all requests
	Headers map[string]string

	// Timeout for requests
	Timeout time.Duration `fog:"default:30s"`

	// RetryCount for failed requests
	RetryCount int `fog:"default:3"`

	// RetryDelay between attempts
	RetryDelay time.Duration `fog:"default:1s"`

	// UserAgent string
	UserAgent string `fog:"default:FoggoClient/1.0"`

	// TLSSkipVerify disables TLS verification (not recommended for production)
	TLSSkipVerify bool

	// MaxIdleConns controls the maximum number of idle connections
	MaxIdleConns int `fog:"default:100"`

	// MaxConnsPerHost limits connections per host
	MaxConnsPerHost int `fog:"default:10"`

	// EnableDebug turns on debug logging
	EnableDebug bool

	// ProxyURL for HTTP proxy
	ProxyURL string

	// BearerToken for authentication
	BearerToken string
}