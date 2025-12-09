// Package logger provides logging utilities for foggo.
package logger

import (
	"io"
	"log"
)

// InitializeLogger creates a new logger with the specified output and prefix.
func InitializeLogger(out io.Writer, prefix string) *log.Logger {
	logger := log.New(out, prefix, log.Ldate|log.Ltime)
	return logger
}
