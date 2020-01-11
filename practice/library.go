package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	. "github.com/ttimt/GoogleHashCode-2020-Qualification/stdlib"
)

/* DON'T HAVE TO TOUCH ANYTHING BELOW UNLESS MODIFICATION REQUIRED */

// Print the score of the problem.answers with file path
func (p *problem) printScore() {
	fmt.Println("Score of", p.filePath, ":", p.score)
}

// Calculate answers score and store result in p.score
// Access answer struct with p.answers (type is a slice of answer)
func (p *problem) calcScore() {
	p.score = calcScore(p.answers)
}

// Run all datasets according to the input string
func runDataSets(datasets string) {
	wg.Add(len(datasets))

	for k := range datasets {
		filePath := getFileName(string(datasets[k]))

		go runDataSet(filePath)
	}
}

// Write to file with name "output_datasetFileName"
func (p *problem) writeFile() {
	writer := NewWriter(prefixFilePath + prefixOutputFolderPath + "output_" + p.filePath)

	err := writer.WriteLine(p.writeFirstLine(), writeFirstLine)
	errorCheck(err)

	for k := range p.answers {
		err = writer.WriteLine(p.answers[k].writeData(), writeOtherLines)
		errorCheck(err)
	}

	writer.CloseFile()
}

// Read first line to problem struct and remaining lines of first to problemData struct
func readFile(filePath string) *problem {
	// Create a new problem instance
	p := &problem{}
	p.filePath = filePath

	// Update the file path
	filePath = prefixFilePath + prefixDatasetFolderPath + filePath

	// Create a new reader
	reader, err := NewReader(filePath)
	errorCheck(err)

	// Set starting ID
	reader.ID = StartID

	reader.ReadFirstLine(readFirstLine[0])
	errorCheck(reader.Err)
	p.readFirstLine(reader.Data)

	for reader.ReadNextData(readOtherLines[0]) {
		errorCheck(reader.Err)

		var d problemData

		d.readData(reader.Data, reader)
		d.ID = reader.GetNewID()

		p.data = append(p.data, d)
	}

	return p
}

func getAllFileName() []string {
	files, err := ioutil.ReadDir(prefixFilePath + prefixDatasetFolderPath)
	errorCheck(err)

	var datasetFilesName []string

	for k := range files {
		datasetFilesName = append(datasetFilesName, strings.ToLower(files[k].Name()))
	}

	return datasetFilesName
}

// Get file name according to the first character of file (A, B, C etc...)
func getFileName(datasetID string) string {
	files, err := ioutil.ReadDir(prefixFilePath + prefixDatasetFolderPath)
	errorCheck(err)

	for k := range files {
		if strings.ToLower(files[k].Name())[:1] == strings.ToLower(datasetID) {
			return strings.ToLower(files[k].Name())
		}
	}

	return ""
}

// Check for non nil error and panic
func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

// In case auto file name retrieval does not work
func ReadFileSpecial() {
	for _, s := range getAllFileName() {
		wg.Add(1)
		runDataSet(s)
	}
}
