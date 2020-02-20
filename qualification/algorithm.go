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

}

// Secondary algorithm
//
func (p *problem) algorithm2() {
	// sort.Slice(p.libraries, func(i, j int) bool {
	// 	return p.libraries[i].nrOfBooks > p.libraries[j].nrOfBooks
	// })
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(p.libraries), func(i, j int) {
		p.libraries[i], p.libraries[j] = p.libraries[j], p.libraries[i]
	})

	for k := range p.libraries {
		sort.Slice(p.libraries[k].books, func(i, j int) bool {
			return p.libraries[k].books[i].score > p.libraries[k].books[j].score
		})
	}

	curLibrary := 0
	for i := 0; i < p.nrOfDays && curLibrary < len(p.libraries); {
		if i+p.libraries[curLibrary].signUpDuration <= p.nrOfDays && !p.libraries[curLibrary].assigned {
			i += p.libraries[curLibrary].signUpDuration
			p.answers = append(p.answers, answer{library: &p.libraries[curLibrary], signUpEndDay: i})
			p.libraries[curLibrary].assigned = true
		}
		curLibrary++
	}

	// Remaining libraries
	for i := 0; i < 10; i++ {
		curLibrary = 0
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
		}
	}

	for k := range p.answers {
		maxBooks := (p.nrOfDays - p.answers[k].signUpEndDay) * p.answers[k].shipPerDay

		for j := 0; j < len(p.answers[k].books) && j < maxBooks && !p.answers[k].books[j].assigned; j++ {
			p.answers[k].booksAns = append(p.answers[k].booksAns, p.answers[k].books[j])
			p.answers[k].books[j].assigned = true
		}

		// libraryDays := p.answers[k].signUpEndDay + 1
		// for i := 0; i < p.answers[k].nrOfBooks; i += p.answers[k].shipPerDay {
		// 	libraryDays++
		// 	bookPerDay := 0
		// 	for j := range p.answers[k].books {
		// 		if bookPerDay < p.answers[k].shipPerDay {
		// 			if _, ok := p.uniqueBooks[p.answers[k].books[j].ID]; !ok {
		// 				p.uniqueBooks[p.answers[k].books[j].ID] = struct{}{}
		// 				p.answers[k].booksAns = append(p.answers[k].booksAns, p.answers[k].books[j])
		// 				bookPerDay++
		// 			}
		// 		}
		// 	}
		// }
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
