package stdlib

import (
	"bufio"
	"io"
	"os"
)

// Reader store reader struct for write and pause reading by storing pointer of file and io reader
// Read data and error will be stored in the struct
type Reader struct {
	file     *os.File
	ioReader *bufio.Reader
	Data     []InputString
	Err      error
	ID       int
	lastRead bool
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
func (r *Reader) ReadFirstLine(del byte) bool {
	firstLine, err := r.ioReader.ReadString(del)
	if err != nil {
		_ = r.file.Close()
		r.Err = err
		return false
	}

	r.Data = DataSplit(firstLine, " ")

	return true
}

// ReadNextData will read next line of data
func (r *Reader) ReadNextData(del byte) bool {
	if r.lastRead {
		_ = r.file.Close()
		return false
	}

	line, err := r.ioReader.ReadString(del)

	if err != nil {
		// if err != io.EOF || len(line) == 0 {
		// 	_ = r.file.Close()
		// 	r.Err = err
		// 	return false
		// } else {
		// 	r.lastRead = true
		// }
		if err == io.EOF && line != "" {
			r.lastRead = true
		} else {
			_ = r.file.Close()
			r.Err = err
			return false
		}
	}

	r.Data = DataSplit(line, " ")

	return true
}

// GetNewID retrieves a new ID for the problem data
func (r *Reader) GetNewID() int {
	r.ID++

	return r.ID - 1
}
