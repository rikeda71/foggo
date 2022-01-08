package writer

import (
	"fmt"
	"os"
	"strings"
)

type Writer struct{}

func InitializeWriter() *Writer {
	return &Writer{}
}

func (w *Writer) Write(code string, baseFileName string) error {
	fn := w.createFileName(baseFileName)
	f, err := os.Create(fn)
	if err != nil {
		return fmt.Errorf("file open error: %w", err)
	}
	_, err = f.Write([]byte(code))
	return err
}

func (w *Writer) createFileName(baseFileName string) string {
	return strings.Replace(baseFileName, ".go", "_gen.go", 1)
}
