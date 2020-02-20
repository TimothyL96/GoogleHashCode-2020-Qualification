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

	// Sort book by score
	for k := range p.libraries {
		sort.Slice(p.libraries[k].books, func(i, j int) bool {
			return p.libraries[k].books[i].score > p.libraries[k].books[j].score
		})
	}

	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(p.libraries), func(i, j int) {
		p.libraries[i], p.libraries[j] = p.libraries[j], p.libraries[i]
	})

	curLibrary := 0
	lastDay := 0
	if len(p.answers) > 0 {
		lastDay = p.answers[len(p.answers)-1].signUpEndDay
	}

	for i := lastDay; i < p.nrOfDays && curLibrary < len(p.libraries); {
		if i+p.libraries[curLibrary].signUpDuration <= p.nrOfDays && !p.libraries[curLibrary].assigned {
			i += p.libraries[curLibrary].signUpDuration
			p.answers = append(p.answers, answer{library: &p.libraries[curLibrary], signUpEndDay: i})
			p.libraries[curLibrary].assigned = true
		}
		curLibrary++

		rand.Seed(time.Now().Unix())
		rand.Shuffle(len(p.libraries), func(i, j int) {
			p.libraries[i], p.libraries[j] = p.libraries[j], p.libraries[i]
		})
	}

	for k := range p.answers {
		maxBooks := (p.nrOfDays - p.answers[k].signUpEndDay) * p.answers[k].shipPerDay

		for j := 0; j < len(p.answers[k].books) && j < maxBooks && !p.answers[k].books[j].assigned; j++ {
			p.answers[k].booksAns = append(p.answers[k].booksAns, p.answers[k].books[j])
			p.answers[k].books[j].assigned = true
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

	// Get first day max score
	for k := range p.libraries {
		p.libraries[k].maxScore = 0
		days := p.nrOfDays - p.libraries[k].signUpDuration
		books := days * p.libraries[k].shipPerDay
		if books > len(p.libraries[k].books) {
			books = len(p.libraries[k].books)
		}

		for i := 0; i < len(p.libraries[k].books) && i < books; i += p.libraries[k].shipPerDay {
			for j := i; j < p.libraries[k].shipPerDay; j++ {
				p.libraries[k].maxScore += p.libraries[k].books[i].score
			}
		}
	}

	// Sort by max score
	sort.Slice(p.libraries, func(i, j int) bool {
		return p.libraries[i].maxScore > p.libraries[j].maxScore
	})

	// assign libraries
	last := 0
	for k := range p.libraries {
		if last+p.libraries[k].signUpDuration <= p.nrOfDays && !p.libraries[k].assigned {
			p.libraries[k].assigned = true
			last += p.libraries[k].signUpDuration
			p.answers = append(p.answers, answer{library: &p.libraries[k], signUpEndDay: last})
		}

		// Re-get max score
		for k := range p.libraries {
			if !p.libraries[k].assigned {
				p.libraries[k].maxScore = 0
				days := p.nrOfDays - p.libraries[k].signUpDuration - last
				books := days * p.libraries[k].shipPerDay
				if books > len(p.libraries[k].books) {
					books = len(p.libraries[k].books)
				}

				for i := 0; i < len(p.libraries[k].books) && i < books; i += p.libraries[k].shipPerDay {
					for j := i; j < p.libraries[k].shipPerDay; j++ {
						if !p.libraries[k].books[i].assigned {
							p.libraries[k].maxScore += p.libraries[k].books[i].score
						}
					}
				}
			}
		}

		// Sort by max score
		sort.Slice(p.libraries, func(i, j int) bool {
			return p.libraries[i].maxScore > p.libraries[j].maxScore
		})

	}

	// Create books from answer
	for k := range p.answers {
		maxBooks := (p.nrOfDays - p.answers[k].signUpEndDay) * p.answers[k].shipPerDay

		for j := 0; j < len(p.answers[k].books) && j < maxBooks && !p.answers[k].books[j].assigned; j++ {
			p.answers[k].booksAns = append(p.answers[k].booksAns, p.answers[k].books[j])
			p.answers[k].books[j].assigned = true
		}

	}
}

// Default recursive algorithm
//
func (p *problem) recursive(data, curData []problemData, curPD problemData, maxData []answer, maxScore *int, currentScore int) []answer {
	// Return if max reached
	if true { // *maxScore == p.maxPizzaSlices
		return maxData
	}

	return maxData
}

// Endless algorithm till max reached or interrupt signalled
func (p *problem) algorithmEndless() {
	p.algorithm2()
}

// Run recursive algorithm
func (p *problem) algorithmBruteForce() {
	// p.recursive()
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
