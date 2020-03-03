package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Main algorithm
//
// To sort p.books ID ascending :
// sort.Slice(p.books, func(i, j int) bool {
// 	return p.books[i].ID < p.books[j].ID
// })
//
func (p *problem) algorithm1(libs []library, maxScore int, ans []answer, days int, uniqueBooks map[int]struct{}) {

}

// Secondary algorithm
// Simulated annealing
func (p *problem) algorithm2() {
	// Set initial temperature
	// Set cooling rate
	// Generate initial generation

	// while temperature still larger than 1
	// Duplicate population, and perform random swap
	// Calculate energy/score
	// Decide if we should accept with acceptance probability
	// Record new population if score better than previous
	// Cool the temperature
	p.answers = p.generateIndividual()
	fmt.Println("New individual score:", p.calcScoreBase(p.answers))
}

// Calculate acceptance probability
func acceptanceProbability(energy, newEnergy int, temperature float64) float64 {
	if newEnergy > energy {
		return 1
	}

	return math.Exp2(float64(newEnergy-energy) / temperature)
}

// Generate individuals
func (p *problem) generateIndividual() (newAns []answer) {
	// Shuffle libraries
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(p.libraries), func(i, j int) {
		p.libraries[i], p.libraries[j] = p.libraries[j], p.libraries[i]
	})

	// Assign libraries
	lastDay := 0

	for k := range p.libraries {
		if lastDay+p.libraries[k].signUpDuration <= p.nrOfDays {
			lastDay += p.libraries[k].signUpDuration
			newAns = append(newAns, answer{
				library:      p.libraries[k],
				signUpEndDay: lastDay})
		}
	}

	// Store ID of assigned books
	plannedBooks := make(map[int]struct{})

	// Assign books for the assigned libraries
	for k := range newAns {
		// Get number of books allowed to be shipped by this library
		allowedBooks := (p.nrOfDays - newAns[k].signUpEndDay) * newAns[k].shipPerDay

		// Shuffle books as well
		rand.Shuffle(len(newAns[k].books), func(i, j int) {
			newAns[k].books[i], newAns[k].books[j] = newAns[k].books[j], newAns[k].books[i]
		})

		// Assign books
		for j := range newAns[k].books {
			if _, ok := plannedBooks[newAns[k].books[j].ID]; !ok && allowedBooks > 0 {
				plannedBooks[newAns[k].books[j].ID] = struct{}{}
				newAns[k].booksAns = append(newAns[k].booksAns, newAns[k].books[j])
				allowedBooks--
			}
		}
	}

	return
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
