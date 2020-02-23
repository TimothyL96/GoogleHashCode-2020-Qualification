package main

import (
	"sort"
)

// Main algorithm
//
// To sort p.books ID ascending :
// sort.Slice(p.books, func(i, j int) bool {
// 	return p.books[i].ID < p.books[j].ID
// })
//
func (p *problem) algorithm1(libs []library, maxScore int, ans []answer, days int, uniqueBooks map[int]struct{}) {
	if len(libs) == 0 {
		if maxScore > p.score {
			p.answers = ans
			p.calcScore()
		}

		return
	}

	for k := range libs {
		if libs[k].signUpDuration+days <= p.nrOfDays {
			newAns := answer{library: &libs[k]}
			maxBooks := (p.nrOfDays - (libs[k].signUpDuration + days)) * libs[k].shipPerDay
			for j := range libs[k].books {
				if _, ok := uniqueBooks[libs[k].books[j].ID]; !ok && maxBooks > 0 {
					maxScore += libs[k].books[j].score
					newAns.booksAns = append(newAns.booksAns, libs[k].books[j])
					uniqueBooks[libs[k].books[j].ID] = struct{}{}
					maxBooks--
				}
			}
			ans = append(ans, newAns)

			uniqueBooksNew := make(map[int]struct{})
			for j := range uniqueBooks {
				uniqueBooksNew[j] = struct{}{}
			}

			p.algorithm1(libs[k+1:], maxScore, ans, days+libs[k].signUpDuration, uniqueBooksNew)
		}
		p.algorithm1(libs[k+1:], maxScore, ans, days, uniqueBooks)
	}
}

// Secondary algorithm
//
func (p *problem) algorithm2() {
	for k := range p.libraries {
		sort.Slice(p.libraries[k].books, func(i, j int) bool {
			return p.libraries[k].books[i].score > p.libraries[k].books[j].score
		})
	}
	p.algorithm1(p.libraries, 0, []answer{}, 0, make(map[int]struct{}))
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
