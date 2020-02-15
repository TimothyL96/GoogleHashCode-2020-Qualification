package main

import (
	"sync"
)

const (
	// Path for source code and such, should be updated before competition start
	prefixFilePath = "C:\\Users\\Timothy\\go\\src\\github.com\\ttimt\\GoogleHashCode-2020-Qualification\\qualification\\"

	// All folder names below should resides inside prefixFilePath above
	//
	// Folder containing all the datasets
	prefixDatasetFolderPath = "datasets\\"

	// Scoring submission output folder
	prefixOutputFolderPath = "output\\"

	// Constants to used for reading write new line or space
	rwNewLine = "\n"
	rwSpace   = " "

	// What happens when you read/write first line, usually first line is used so no changes needed
	readFirstLine  = rwNewLine
	writeFirstLine = rwNewLine

	// What happens when you read other lines
	// If each data is separated to new line, use rwNewLine
	// If there's only 1 line for all data use rwSpace
	// Like in 2020 pizza practice problem is using rwSpace
	// Ex: 2 3 4 5 6
	// In 2019 slide show qualification problem is using rwNewLine.
	// Ex:
	// H 2 cat beach
	// V 3 dog green garden
	readOtherLines = rwNewLine

	// What is the delimiter for separating data when writing to new line
	// Use rwNewLine if data is separated by new line
	// Like in 2020 pizza practice problem is using rwSpace. Ex: 2 3 4 5 6 7
	writeOtherLines = rwNewLine

	// Starting ID for problem data struct
	// Ex: Photos in 2019 qualification problem, type of pizzas in 2020 practice problem
	startID = 0
)

var wg sync.WaitGroup

// The initial struct for the problem
// Ex: var nrOfPhotos int - Number of photos in the dataset file
type problem struct {
	// DEFAULT
	data     []problemData
	answers  []answer
	score    int
	filePath string

	// PROBLEM SPECIFIC FIELDS

}

// Struct for the data
// Ex: var nrOfTags - Number of tags in photo with ID 3
type problemData struct {
	// DEFAULT
	ID int

	// PROBLEM SPECIFIC FIELDS

}

// Struct to store per data for the final answer
// Ex: type answer struct {
//          data problemData
// 	   }
//
// *Note - This is a slice in the problem struct above
// Slice inside this answer struct should be avoided if unnecessary
type answer struct {
	*problemData
}

func main() {
	var datasets string

	// Uncomment any dataset that you'll want to run concurrently and vice versa
	// **************** //

	datasets += "A"
	datasets += "B"
	datasets += "C"
	datasets += "D"
	datasets += "E"

	// **************** //
	// For more datasets, simply add new line above as needed, according to first character of dataset file
	// Ex:  datasets += "F"
	//
	// If it doesn't work / in case datasets are not named a_xxx, b_xxx ....
	// Use (Uncomment line below):
	// readFileSpecial()
	// This will read ALL files in prefixDatasetFolderPath and run them as a dataset respectively
	// Comment out line:
	// "var datasets string"
	// "datasets += "A"
	// ...
	// "runDataSets(datasets)"
	//
	// To run single/few specific datasets if above doesn't work, manually enter file names:
	// Add file path below as required (Uncomment lines below):
	//
	// filePath := []string{
	// 	"xxx.txt",
	// 	"yyy.txt", // Add a comma at the last line
	// }
	// for k := range filePath {
	// 	wg.Add(1)
	// 	runDataSet(prefixFilePath + prefixDatasetFolderPath + filePath[k])
	// }
	//
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
	p := readFile(filePath)

	// Run the main algorithm - code it in algorithm.go
	// Call and comment other algorithms as needed
	p.algorithm1()
	// p.algorithm2()
	// p.runRecursive()

	// Calculate the score  - code it in algorithm.go
	p.calcScore()

	// Print the score out
	p.printScore()

	// Print ID of answers out in sequence
	// Use it for DEBUGGING
	// p.printAnswer()

	// Write to file:
	// Remember to update writeFirstLine() and writeData()
	p.writeFile()

	// Indicate the goroutine has finished its task
	wg.Done()
}
