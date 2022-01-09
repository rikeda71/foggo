package logger

import (
	"io"
	"log"
)

func InitializeLogger(out io.Writer, prefix string) *log.Logger {
	logger := log.New(out, prefix, log.Ldate|log.Ltime)
	return logger
}
