package main

import (
	. "github.com/ttimt/GoogleHashCode-2020-Qualification/stdlib"
)

// Write to file by returning the intended line
// *Note - New line is automatically inserted

// Write first line to the submission file
// Example: Number of photos in a slide show (str = IntToString(len(p.answers)))
//
// If more than 1 variables per line:
// Ex: str := IntToString(a.ID) + " " + IntToString(a.Orientation)
func (p *problem) writeFirstLine() (str string) {
	count := 0
	for k := range p.answers {
		if len(p.answers[k].booksAns) > 0 {
			count++
		}
	}
	str = IntToString(count)

	return
}

// Write remaining books to the file
//
// *Note - p.answers will be automatically traversed
// Just indicate what to output per dataset
// Example: Photos ID in a slide in sequence (str = IntToString(len(p.answer.ID)))
//
// If more than 1 variables per line:
// Ex: str := IntToString(a.ID) + " " + IntToString(a.Orientation)
//
// Use IntToString to convert integers to string
func (a *answer) writeData() (str string) {
	if len(a.booksAns) == 0 {
		return ""
	}

	str = IntToString(a.ID) + " " + IntToString(len(a.booksAns)) + "\n"

	for k := range a.booksAns {
		str += IntToString(a.booksAns[k].ID) + " "
	}

	str = str[:len(str)-1]

	return
}
