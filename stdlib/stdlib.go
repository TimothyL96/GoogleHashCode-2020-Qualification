package stdlib

import (
	"strconv"
	"strings"
)

// InputString is an alias of string to add methods GetInt and GetString for faster conversion
type InputString string

// GetInt returns the integer value of an string and not the ASCII value of the string
func (d InputString) GetInt() int {
	value, _ := strconv.Atoi(d.GetString())
	return value
}

// GetString returns the given string trimmed from space
func (d InputString) GetString() string {
	return strings.TrimSpace(string(d))
}

// DataSplit accepts a string and delimiter and split the string by the delimiter and returns a slice of it
func DataSplit(str string, del string) []InputString {
	data := strings.Split(InputString(str).GetString(), del)

	inputData := make([]InputString, len(data))
	for k := range data {
		inputData[k] = InputString(data[k])
	}

	return inputData
}

// IntToString directly converts integer to string and not to ASCII representation of the integer value
func IntToString(i int) string {
	return strconv.Itoa(i)
}
