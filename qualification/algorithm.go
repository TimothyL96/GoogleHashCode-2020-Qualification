package main

import (
	"math/rand"
	"sort"
	"time"
)

// Main algorithm
//
// To sort p.data ID ascending :
// sort.Slice(p.data, func(i, j int) bool {
// 	return p.data[i].ID < p.data[j].ID
// })
//
func (p *problem) algorithm1() {
	for _ = range p.libraries {
		// Update max score
		maxLibrary := p.updateMaxScore()

		// Store current answer length
		lastSize := len(p.answers)

		// If no max score exist or no max score library
		if maxLibrary == nil || maxLibrary.maxScore == 0 || maxLibrary.assigned {
			break
		}

		// Set library to assigned
		maxLibrary.assigned = true

		// Append library to answer
		p.answers = append(p.answers, answer{library: maxLibrary, signUpEndDay: p.lastDay})

		// Increment current last day with current library sign up duration
		p.lastDay += maxLibrary.signUpDuration

		// Create books from last signed up library
		lastAnswer := &p.answers[len(p.answers)-1]

		// Break if no library assigned or available books left
		if len(p.answers) == lastSize {
			break
		}

		// Get max number of book that can be shipped from library last day to max days
		maxBooks := (p.nrOfDays - p.lastDay) * p.answers[len(p.answers)-1].shipPerDay

		// Iterate through all the books in the library
		for k := range lastAnswer.books {
			// Process unassigned book and only if there's capacity to assign (maxBooks)
			if !lastAnswer.books[k].assigned && maxBooks > 0 {
				// Append book to book answer
				lastAnswer.booksAns = append(lastAnswer.booksAns, lastAnswer.books[k])

				// Set book to assigned
				lastAnswer.books[k].assigned = true

				// Decrement number of allowed books capacity
				maxBooks--
			}
		}
	}
}

// Secondary algorithm
//
func (p *problem) algorithm2() {
	// Sort book by score
	for k := range p.libraries {
		sort.Slice(p.libraries[k].books, func(i, j int) bool {
			return p.libraries[k].books[i].score > p.libraries[k].books[j].score
		})
	}

	// Get frequency and sort books
	for k := range p.libraries {
		for j := range p.libraries[k].books {
			p.data[p.libraries[k].books[j].ID].frequency++
		}
	}
	sort.Slice(p.data, func(i, j int) bool {
		return p.data[i].frequency > p.data[j].frequency
	})
	// Update score with frequency
	for k := range p.data {
		if p.data[k].frequency == 0 {
			continue
		}
		p.data[k].scoreWFrequency = p.data[k].score / p.data[k].frequency
	}

	// Setup library book hash
	for k := range p.libraries {
		p.libraries[k].booksHash = make(map[int]struct{})

		for j := range p.libraries[k].books {
			p.libraries[k].booksHash[p.libraries[k].books[j].ID] = struct{}{}
		}
	}

	p.algorithm1()
}

func (p *problem) updateMaxScore() *library {
	maxScore1 := 0
	maxScore2 := 0
	maxScore3 := 0
	maxScore4 := 0
	maxScore5 := 0
	var maxLibrary1 *library
	var maxLibrary2 *library
	var maxLibrary3 *library
	var maxLibrary4 *library
	var maxLibrary5 *library

	// Get max score
	for k := range p.libraries {
		// Reset max score
		p.libraries[k].maxScore = 0

		// Only process unassigned libraries and library within max date
		if !p.libraries[k].assigned && p.libraries[k].signUpDuration+p.lastDay < p.nrOfDays {
			// Get number of books that can be assigned
			books := (p.nrOfDays - (p.libraries[k].signUpDuration + p.lastDay)) * p.libraries[k].shipPerDay

			// Add max score from unassigned books where number of books = number of books that can be assigned
			for i := range p.libraries[k].books {
				if !p.libraries[k].books[i].assigned && books > 0 {
					p.libraries[k].maxScore += p.libraries[k].books[i].scoreWFrequency
					books--
				}
			}

			p.libraries[k].maxScore /= p.libraries[k].shipPerDay

			if p.libraries[k].maxScore > maxScore1 {
				maxScore5 = maxScore4
				maxScore4 = maxScore3
				maxScore3 = maxScore2
				maxScore2 = maxScore1
				maxLibrary5 = maxLibrary4
				maxLibrary4 = maxLibrary3
				maxLibrary3 = maxLibrary2
				maxLibrary2 = maxLibrary1

				maxLibrary1 = &p.libraries[k]
				maxScore1 = p.libraries[k].maxScore
			} else if p.libraries[k].maxScore > maxScore2 {
				maxScore5 = maxScore4
				maxScore4 = maxScore3
				maxScore3 = maxScore2
				maxLibrary5 = maxLibrary4
				maxLibrary4 = maxLibrary3
				maxLibrary3 = maxLibrary2

				maxLibrary2 = &p.libraries[k]
				maxScore2 = p.libraries[k].maxScore
			} else if p.libraries[k].maxScore > maxScore3 {
				maxScore5 = maxScore4
				maxScore4 = maxScore3
				maxLibrary5 = maxLibrary4
				maxLibrary4 = maxLibrary3

				maxLibrary3 = &p.libraries[k]
				maxScore3 = p.libraries[k].maxScore
			} else if p.libraries[k].maxScore > maxScore4 {
				maxScore5 = maxScore4
				maxLibrary5 = maxLibrary4

				maxLibrary4 = &p.libraries[k]
				maxScore4 = p.libraries[k].maxScore
			} else if p.libraries[k].maxScore > maxScore5 {
				maxLibrary5 = &p.libraries[k]
				maxScore5 = p.libraries[k].maxScore
			}
		}
	}

	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(20)
	if r == 0 || r == 1 || r == 2 || r == 3 || r == 4 || r == 15 || r == 16 || r == 17 || r == 8 {
		return maxLibrary1
	} else if r == 5 || r == 6 || r == 7 || r == 13 {
		return maxLibrary3
	} else if r == 9 || r == 10 || r == 11 || r == 18 || r == 19 {
		return maxLibrary4
	} else if r == 12 {
		return maxLibrary2
	}

	return maxLibrary5
}

// Endless algorithm till max reached or interrupt signalled
func (p *problem) algorithmEndless() {
	p.algorithm2()
}

// Calculate score from input
// Access answer struct with p.answers (type is a slice of answer)
func (p *problem) calcScoreBase(answers []answer) int {
	score := 0

	for k := range answers {
		for j := range answers[k].booksAns {
			score += answers[k].booksAns[j].score
		}
	}

	return score
}
