package main

import (
	"flag"
	"sync"
)

// Best known score:
// A - 21
// B - 5,822,900
// C - 5,690,870
// D - 5,107,115
// E - 5,237,345
// F - 5,348,248

const (
	// Path for source code and such, should be updated before competition start
	// Always add a slash '\\' behind directory
	prefixFilePath = ".\\"

	// All folder names below should resides inside prefixFilePath above
	//
	// Folder containing all the datasets
	prefixDatasetFolderPath = "datasets\\"

	// Scoring submission output folders
	prefixOutputFolderPath           = "output_best\\"
	prefixLastOutputFolderPath       = "output_last\\"
	prefixEndlessOutputFolderPath    = "output_endless\\"
	prefixBruteForceOutputFolderPath = "output_brute\\"

	// Constants to used for reading write new line or space
	rwNewLine = "\n"
	rwSpace   = " "

	// What happens when you read/write first line, usually first line is used so no changes needed
	readFirstLine  = rwNewLine
	writeFirstLine = rwNewLine

	// What happens when you read other lines
	// If each books is separated to new line, use rwNewLine
	// If there's only 1 line for all books use rwSpace
	// Like in 2020 pizza practice problem is using rwSpace
	// Ex: 2 3 4 5 6
	// In 2019 slide show qualification problem is using rwNewLine.
	// Ex:
	// H 2 cat beach
	// V 3 dog green garden
	readOtherLines = rwNewLine

	// What is the delimiter for separating books when writing to new line
	// Use rwNewLine if books is separated by new line
	// Like in 2020 pizza practice problem is using rwSpace. Ex: 2 3 4 5 6 7
	writeOtherLines = rwNewLine

	// Starting ID for problem books struct
	// Ex: Photos in 2019 qualification problem, type of pizzas in 2020 practice problem
	startID = 0
)

var wg sync.WaitGroup
var outputFolder string
var endlessRun bool

// The initial struct for the problem
// Ex: var nrOfPhotos int - Number of photos in the dataset file
type problem struct {
	// DEFAULT
	books             []book
	libraries         []library
	answers           []answer
	score             int
	previousBestScore int
	filePath          string

	// PROBLEM SPECIFIC FIELDS
	nrOfBooks     int
	nrOfLibraries int
	nrOfDays      int
	lastDay       int
	maxScore      int
}

// Struct for the books
// Ex: var nrOfTags - Number of tags in photo with ID 3
// books
type book struct {
	// DEFAULT
	ID       int
	assigned bool

	// PROBLEM SPECIFIC FIELDS
	score     int
	frequency int
}

type library struct {
	ID             int
	nrOfBooks      int
	signUpDuration int
	shipPerDay     int
	books          []book
	booksHash      map[int]struct{}
	assigned       bool
	maxScore       int
}

// Struct to store per books for the final answer
// Ex: type answer struct {
//          books book
// 	   }
//
// *Note - This is a slice in the problem struct above
// Slice inside this answer struct should be avoided if unnecessary
type answer struct {
	*library
	signUpEndDay int
	booksAns     []book
}

func init() {
	flag.BoolVar(&endlessRun, "endless", false, "Execute endless run")
}

func main() {
	// Parse CLI flags
	flag.Parse()

	// Set output folder based on flags
	if endlessRun {
		outputFolder = prefixEndlessOutputFolderPath
	} else {
		outputFolder = prefixOutputFolderPath
	}

	var datasets string

	// Uncomment any dataset that you'll want to run concurrently and vice versa
	// **************** //

	// datasets += "A"
	datasets += "B"
	datasets += "C"
	// datasets += "D"
	// datasets += "E"
	datasets += "F"

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
	// Read books from the file path and return to p as problem struct
	// Remember to update readFirstLine() and readData()
	p := readFile(filePath)

	// Run the main algorithm - code it in algorithm.go
	// Call and comment other algorithms as needed
	// p.algorithm1()
	p.algorithm2()

	// Calculate the score  - code it in algorithm.go
	p.calcScore()

	// Print the score out
	p.printScore()

	// Write to file:
	// Remember to update writeFirstLine() and writeData()
	p.writeFile()

	// Indicate the goroutine has finished its task
	wg.Done()
}

// Execute endless run
func runEndless(filePath string) {
	// Read books from the file path and return to p as problem struct
	// Remember to update readFirstLine() and readData()
	p := readFile(filePath)

	for k := range p.books {
		p.maxScore += p.books[k].score
	}

	for p.score != p.maxScore {
		p.answers = nil

		for k := range p.books {
			p.books[k].assigned = false
		}
		for k := range p.libraries {
			p.libraries[k].assigned = false
		}
		p.lastDay = 0
		p.algorithmEndless()

		// Calculate the score  - code it in algorithm.go
		p.calcScore()

		// p.printScore()
		if p.score > p.previousBestScore {

			// Print the score out
			p.printScore()

			// Write to file:
			// Remember to update writeFirstLine() and writeData()
			p.writeBest()
		}
	}

	// Indicate the goroutine has finished its task
	wg.Done()
}
