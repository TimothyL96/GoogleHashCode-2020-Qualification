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

		if endlessRun {
			go runEndless(filePath)
		} else if bruteForceRun {
			go runBruteForce(filePath)
		} else {
			go runDataSet(filePath)
		}
	}
}

// Write to file with name "output_datasetFileName"
//
// We now first check if this is the highest score recorded,
// if yes write to output_best and output_last and update score at score_filename,txt,
// or else just write to output_last
//
func (p *problem) writeFile() {
	outputLast := prefixFilePath + prefixLastOutputFolderPath

	// Write the best score
	p.writeBest()

	// Write to last output folder:
	// Write submission file
	writer := NewWriter(outputLast + "output_" + p.filePath)

	err := writer.WriteLine(p.writeFirstLine(), writeFirstLine)
	errorCheck(err)

	for k := range p.answers {
		err = writer.WriteLine(p.answers[k].writeData(), writeOtherLines)
		errorCheck(err)
	}

	// Write score to file
	writerScore := NewWriter(outputLast + "score_" + p.filePath)
	err = writerScore.WriteLine(IntToString(p.score), writeFirstLine)
	errorCheck(err)

	writer.CloseFile()
	writerScore.CloseFile()
}

// Write only the best score ever recorded
func (p *problem) writeBest() {
	outputBest := prefixFilePath + outputFolder

	// Write to best output folder if better than previous recorded score:
	if p.score > p.previousBestScore {
		// Write submission file
		writer := NewWriter(outputBest + "output_" + p.filePath)

		err := writer.WriteLine(p.writeFirstLine(), writeFirstLine)
		errorCheck(err)

		for k := range p.answers {
			err = writer.WriteLine(p.answers[k].writeData(), writeOtherLines)
			errorCheck(err)
		}

		// Write score to file
		writerScore := NewWriter(outputBest + "score_" + p.filePath)
		err = writerScore.WriteLine(IntToString(p.score), writeFirstLine)
		errorCheck(err)

		writer.CloseFile()
		writerScore.CloseFile()

		// Update new previous best score
		p.previousBestScore = p.score

		fmt.Println("Written to best output folder:", p.filePath, "Score:", p.score)
	}
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
	reader.ID = startID

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

	// Read previous best score
	p.readPreviousBest()

	return p
}

// Read previous best score
func (p *problem) readPreviousBest() {
	// Read the highest score from best output folder:
	// Create a new reader
	reader, err := NewReader(prefixFilePath + outputFolder + "score_" + p.filePath)

	p.previousBestScore = -1

	if err == nil {
		reader.ReadFirstLine(readFirstLine[0])
		errorCheck(reader.Err)
		p.previousBestScore = reader.Data[0].GetInt()
	}
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
func readFileSpecial() {
	for _, s := range getAllFileName() {
		wg.Add(1)
		runDataSet(s)
	}
}
