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
	prefixFilePath          = "C:\\Users\\Timothy\\go\\src\\github.com\\ttimt\\GoogleHashCode\\2020\\practice\\"
	prefixDatasetFolderPath = "datasets\\"
	prefixOutputFolderPath  = "output\\"
)

var wg sync.WaitGroup

type problem struct {
	data   []problemData
	answer []answer
	score  int
}

type problemData struct {
}

type answer struct {
}

func main() {
	var datasets string

	datasets += "A"
	// datasets += "B"
	// datasets += "C"
	// datasets += "D"
	// datasets += "E"

	runDataSets(datasets)

	wg.Wait()
}

func runDataSet(filePath string) {
	p := readFile(prefixFilePath + prefixDatasetFolderPath + filePath)

	p.algorithm1()
	p.calcScore()
	p.printScore(filePath)

	p.writeFile(filePath)
	wg.Done()
}

func (p *problem) readFirstLine(dataInput []InputString) {
	// p.nrOfPhotos = dataInput[0].GetInt()
}

func (d *problemData) readData(dataInput []InputString) {
	// d.orientation = dataInput[0].GetString()
}

func (p *problem) writeFirstLine() string {
	str := IntToString(p.score)

	return str
}

func (a *answer) writeData() string {
	str := ""

	return str
}

func (p *problem) writeFile(filePath string) {
	writer := NewWriter(prefixFilePath + prefixOutputFolderPath + "output_" + filePath)

	err := writer.WriteLine(p.writeFirstLine())
	errorCheck(err)

	for k := range p.answer {
		err = writer.WriteLine(p.answer[k].writeData())
		errorCheck(err)
	}

	writer.CloseFile()
}

func runDataSets(datasets string) {
	wg.Add(len(datasets))

	for k := range datasets {
		filePath := getFileName(string(datasets[k]))

		go runDataSet(filePath)
	}
}

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

func (p *problem) printScore(filePath string) {
	fmt.Println("Score of", filePath, ":", p.score)
}

// test
func errorCheck(err error) {
	if err != nil && err != io.EOF {
		panic(err)
	}
}
