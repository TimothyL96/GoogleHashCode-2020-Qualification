package main

import (
	. "github.com/ttimt/GoogleHashCode-2020-Qualification/stdlib"
)

// Example books in line: 1 H cat
// dataInput[0] represent 1
// dataInput[1] represent H
// dataInput[2] represent cat
//
// Use GetInt() if expecting an integer
// and use GetString() vice versa

// Read first line gets the first line books from the file
func (p *problem) readFirstLine(dataInput []InputString) {
	// Store the books from dataInput to p of type problem accordingly
	// Ex: p.nrOfPhotos = dataInput[0].GetInt()
	p.nrOfBooks = dataInput[0].GetInt()
	p.nrOfLibraries = dataInput[1].GetInt()
	p.nrOfDays = dataInput[2].GetInt()
}

// Read first line gets the first line books from the file
func (p *problem) readSecondLine(dataInput []InputString) {
	// Store the books from dataInput to p of type problem accordingly
	// Ex: p.nrOfPhotos = dataInput[0].GetInt()
	for i := 0; i < len(dataInput); i++ {
		book := book{ID: i, score: dataInput[i].GetInt()}
		p.books = append(p.books, book)
	}
}

// Read lines of books excluding first line from the file
func (d *library) readData(dataInput []InputString, reader *Reader, p *problem) {
	// Store the books from dataInput to d of type book
	// d will be stored to p.books[]
	// Ex:
	// d.nrOfTags = dataInput[0].GetInt()
	// d.orientation = dataInput[1].GetString()
	//
	// Ex: To create a map to store a set of tags in a single line in the file
	// d.tags = make(map[int]struct{})
	// for _, v := range dataInput[2:] {
	// 	d.tags[v] = struct {}{}
	// }
	//
	//
	// If there are more than 1 row/line per books, first get the number of rows required:
	// Ex: d.nrOfRows = dataInput[0].GetInt()
	//
	// Then traverse through the rows and read with	reader.ReadNextData(readOtherLines[0])
	// where readOtherLines[0] ('\n') is the delimiter for 1 single line like '\n'
	// And make sure all books from dataInput is retrieved before calling the loop
	//
	// Ex:
	// for i := 0; i < d.nrOfRows; i++ {
	// 	reader.ReadNextData(readOtherLines[0])
	// 	errorCheck(reader.Err)
	// 	d.coordinate[0] = reader.Data[0].GetString() // or create and assign to the proper struct
	// }

	d.nrOfBooks = dataInput[0].GetInt()
	d.signUpDuration = dataInput[1].GetInt()
	d.shipPerDay = dataInput[2].GetInt()

	reader.ReadNextData(readOtherLines[0])
	errorCheck(reader.Err)
	for k := range reader.Data {
		d.books = append(d.books, &p.books[reader.Data[k].GetInt()])
	}
}
