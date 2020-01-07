package stdlib

import (
	"bufio"
	"os"
)

// Writer struct store pointer of file and io writer
type Writer struct {
	file     *os.File
	ioWriter *bufio.Writer
}

// NewWriter returns a new writer with file path from the input given
func NewWriter(filePath string) *Writer {
	// Open file for writing
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	// Create a writer
	ioWriter := bufio.NewWriter(file)

	return &Writer{
		file:     file,
		ioWriter: ioWriter,
	}
}

// WriteLine write a single line to the file
func (w *Writer) WriteLine(str string) error {
	count, err := w.ioWriter.WriteString(str + "\n")

	if count < len(str+"\n") || err != nil {
		w.CloseFile()
		return err
	}

	_ = w.ioWriter.Flush()

	return nil
}

// CloseFile close the file at the end of writing
func (w *Writer) CloseFile() {
	w.file.Close()
}
