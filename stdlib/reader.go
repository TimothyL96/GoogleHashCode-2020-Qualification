package stdlib

import (
	"bufio"
	"os"
)

// Reader store reader struct for write and pause reading by storing pointer of file and io reader
// Read data and error will be stored in the struct
type Reader struct {
	file     *os.File
	ioReader *bufio.Reader
	Data     []InputString
	Err      error
}

// NewReader returns a new reader for reading dataset file
func NewReader(filePath string) (*Reader, error) {
	// Open file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	// Create a reader
	ioReader := bufio.NewReader(file)

	return &Reader{
		file:     file,
		ioReader: ioReader,
	}, nil
}

// ReadFirstLine read first line of file that usually initialize the dataset problem
func (r *Reader) ReadFirstLine() bool {
	firstLine, err := r.ioReader.ReadString('\n')
	if err != nil {
		r.file.Close()
		return false
	}

	r.Data = DataSplit(firstLine, " ")

	return true
}

// ReadNextData will read next line of data
func (r *Reader) ReadNextData() bool {
	line, err := r.ioReader.ReadString('\n')
	if err != nil {
		r.file.Close()
		return false
	}

	r.Data = DataSplit(line, " ")

	return true
}
