package main

import (
	. "github.com/ttimt/GoogleHashCode-2020-Qualification/stdlib"
)

// Write to file by returning the intended line
// *Note - New line is automatically inserted

// Write first line to the submission file
// Example: Number of photos in a slide show (str = IntToString(len(p.answers)))
func (p *problem) writeFirstLine() string {
	str := IntToString(len(p.answers))

	return str
}

// Write remaining data to the file
//
// *Note - p.answers will be automatically traversed
// Just indicate what to output per dataset
// Example: Photos ID in a slide in sequence (str = IntToString(len(p.answer.ID)))
func (a *answer) writeData() string {
	str := IntToString(a.ID)

	return str
}
