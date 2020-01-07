package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"sync"

	. "../stdlib"
)

const (
	// Path for source code and such, should be updated before competition start
	prefixFilePath = "C:\\Users\\Timothy\\go\\src\\github.com\\ttimt\\GoogleHashCode-2020-Qualification\\2020\\qualification\\"

	// All folder names below should resides inside prefixFilePath above
	//
	// Folder containing all the datasets
	prefixDatasetFolderPath = "datasets\\"

	// Scoring submission output folder
	prefixOutputFolderPath = "output\\"
)

var wg sync.WaitGroup

// ******************** INFO ******************** //
// Steps:
// 1. Update const above for directory path and folder name
// 2. Update problem, problemData and answer struct below according to the question
// 3. Update ReadFirstLine and ReadNextData according to the struct defined
// 4. Update WriteFirstLine and WriteData according to the required output
// 5. In algorithm.go :
//      a. Update calcScore with calculation
//      b. Write your algorithm in func algorithm1()
//      c. Optionally, add/remove algorithm and in main.go,
//          update "func runDataSet(string)" to add/change the algorithm
// 6. Update which dataset(s) to run concurrently in "func main()"
// 7. Open terminal in prefixFilePath folder and check algorithm score with "go run ."
// 8. When done, run "zip file.bat" located in prefixFilePath
// 9. Submit to Judge System:
//      a. Source code:
//          source.zip in prefixFilePath
//      b. Output file:
//          output file in prefixFilePath/output
// * Submit early to verify your calcScore method is accurate

// LIBRARY: Hover method name for more information
// func IntToString(int) string - Convert integer to string and not to the ASCII representation
//
// ******************** INFO ******************** //

// The initial struct for the problem
// Ex: var nrOfPhotos int - Number of photos in the dataset file
type problem struct {
	// DEFAULT
	data    []problemData
	answers []answer
	score   int

	// PROBLEM SPECIFIC
}

// Struct for the data
// Ex: var nrOfTags - Number of tags in photo with ID 3
type problemData struct {
}

// Struct to store per data for the final answer
// Ex: type answer struct {
//          data problemData
// 	   }
//
// *Note - This is a slice in the problem struct above
// Slice inside this answer struct should be avoided if unnecessary
type answer struct {
}

func main() {
	var datasets string

	// Uncomment any dataset that you'll want to run concurrently and vice versa
	// **************** //

	datasets += "A"
	// datasets += "B"
	// datasets += "C"
	// datasets += "D"
	// datasets += "E"

	// **************** //

	// Automatically find dataset file name and run the algorithm
	runDataSets(datasets)

	// Wait all concurrent goroutines to finish before exiting
	wg.Wait()
}

// Main flow of program per dataset
func runDataSet(filePath string) {
	// Read data from the file path and return to p as problem struct
	// Remember to update readFirstLine() and readData()
	p := readFile(prefixFilePath + prefixDatasetFolderPath + filePath)

	// Run the main algorithm - code it in algorithm.go
	// Call and comment other algorithms as needed
	p.algorithm1()

	// Calculate the score  - code it in algorithm.go
	p.calcScore()

	// Print the score out
	p.printScore(filePath)

	// Write to file:
	// Remember to update writeFirstLine() and writeData()
	p.writeFile(filePath)

	// Indicate the goroutine has finished its task
	wg.Done()
}

// *************** READ FILE METHOD *************** //
// Example data in line: 1 H cat
// dataInput[0] represent 1
// dataInput[1] represent H
// dataInput[2] represent cat
//
// Use GetInt() if expecting an integer and use GetString() vice versa

// Read first line parameter gets the first line data from the file
func (p *problem) readFirstLine(dataInput []InputString) {
	// Store the data from dataInput to p of type problem accordingly
	// Ex: p.nrOfPhotos = dataInput[0].GetInt()

}

// Read lines of data excluding first line from the file
func (d *problemData) readData(dataInput []InputString) {
	// Store the data from dataInput to d of type problemData
	// d will be stored to p.data[]
	// Ex:
	// d.nrOfTags = dataInput[0].GetInt()
	// d.orientation = dataInput[1].GetString()
}

// *************** END READ FILE METHOD *************** //

// *************** WRITE FILE METHOD *************** //
// Write to file by returning the intended line
// *Note - New line is automatically inserted

// Write first line to the submission file
// Example: Number of photos in a slide show (str = IntToString(len(p.answers)))
func (p *problem) writeFirstLine() string {
	str := IntToString(p.score)

	return str
}

// Write remaining data to the file
//
// *Note - p.answers will be automatically traversed
// Just indicate what to output per dataset
// Example: Photos ID in a slide in sequence (str = IntToString(len(p.answer.ID)))
func (a *answer) writeData() string {
	str := ""

	return str
}

// *************** END WRITE FILE METHOD *************** //

/* DON'T HAVE TO TOUCH ANYTHING BELOW*/
// Print the score of the problem.answers with file path
func (p *problem) printScore(filePath string) {
	fmt.Println("Score of", filePath, ":", p.score)
}

// Write to file with name "output_datasetFileName"
func (p *problem) writeFile(filePath string) {
	writer := NewWriter(prefixFilePath + prefixOutputFolderPath + "output_" + filePath)

	err := writer.WriteLine(p.writeFirstLine())
	errorCheck(err)

	for k := range p.answers {
		err = writer.WriteLine(p.answers[k].writeData())
		errorCheck(err)
	}

	writer.CloseFile()
}

// Run all datasets according to the input string
func runDataSets(datasets string) {
	wg.Add(len(datasets))

	for k := range datasets {
		filePath := getFileName(string(datasets[k]))

		go runDataSet(filePath)
	}
}

// Read first line to problem struct and remaining lines of first to problemData struct
func readFile(filePath string) *problem {
	p := &problem{}

	reader, err := NewReader(filePath)
	errorCheck(err)

	reader.ReadFirstLine()
	errorCheck(reader.Err)
	p.readFirstLine(reader.Data)

	for reader.ReadNextData() {
		errorCheck(reader.Err)

		var d problemData
		d.readData(reader.Data)
		p.data = append(p.data, d)
	}

	return p
}

// Get file name according to the first character of file (A, B, C etc...)
func getFileName(datasetID string) string {
	files, err := ioutil.ReadDir(prefixFilePath + prefixDatasetFolderPath)
	errorCheck(err)

	for k := range files {
		if strings.ToLower(files[k].Name())[:1] == strings.ToLower(datasetID) {
			return files[k].Name()
		}
	}

	return ""
}

// Check for non nil error and panic
func errorCheck(err error) {
	if err != nil && err != io.EOF {
		panic(err)
	}
}
