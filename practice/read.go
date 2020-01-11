package main

import (
	. "github.com/ttimt/GoogleHashCode-2020-Qualification/stdlib"
)

// Example data in line: 1 H cat
// dataInput[0] represent 1
// dataInput[1] represent H
// dataInput[2] represent cat
//
// Use GetInt() if expecting an integer and use GetString() vice versa

// Read first line gets the first line data from the file
func (p *problem) readFirstLine(dataInput []InputString) {
	// Store the data from dataInput to p of type problem accordingly
	// Ex: p.nrOfPhotos = dataInput[0].GetInt()
	p.maxPizzaSlices = dataInput[0].GetInt()
	p.typesOfPizza = dataInput[1].GetInt()
}

// Read lines of data excluding first line from the file
func (d *problemData) readData(dataInput []InputString, reader *Reader) {
	// Store the data from dataInput to d of type problemData
	// d will be stored to p.data[]
	// Ex:
	// d.nrOfTags = dataInput[0].GetInt()
	// d.orientation = dataInput[1].GetString()
	d.nrOfSlices = dataInput[0].GetInt()
}
