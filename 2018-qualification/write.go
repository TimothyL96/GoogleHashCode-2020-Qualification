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
	str = IntToString(len(p.answers))

	return
}

// Write remaining data to the file
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
	for k := range a.vehicles {
		str += IntToString(a.vehicles[k].ID) + " "
	}

	if len(str) > 0 {
		str = str[:len(str)-1]
	}

	return
}
