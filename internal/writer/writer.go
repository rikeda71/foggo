package writer

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Writer struct {
	logger *log.Logger
}

func InitializeWriter(logger *log.Logger) (*Writer, error) {
	if logger == nil {
		return nil, fmt.Errorf("initialize writer error: logger is nil")
	}
	return &Writer{logger: logger}, nil
}

func (w *Writer) Write(code string, baseFileName string) error {
	fn := w.createFileName(baseFileName)
	f, err := os.Create(fn)
	if err != nil {
		return fmt.Errorf("file open error: %w", err)
	}

	_, err = f.Write([]byte(code))
	if err != nil {
		return err
	}

	w.logger.Printf("success to write functional option pattern code to %s", fn)
	return nil
}

func (w *Writer) createFileName(baseFileName string) string {
	return strings.Replace(baseFileName, ".go", "_gen.go", 1)
}
