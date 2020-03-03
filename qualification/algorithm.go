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
func (p *problem) algorithm1() {

}

// Secondary algorithm
// Simulated annealing
func (p *problem) algorithm2() {
	// Set initial temperature
	var tempStart float64 = 20000
	var tempStart1 float64 = 20000

	// Set cooling rate
	// coolingRate := 0.005
	coolingRate := 0.003

	// Generate initial generation
	firstGen := p.generateIndividual()

	// Set as best generation
	bestGen := firstGen
	bestGenScore := p.calcScoreBase(firstGen)

	// while temperature still larger than 1
	for tempStart > 1 {
		// Duplicate population
		secondGen := make([]answer, len(firstGen))
		copy(secondGen, firstGen)

		// Perform random swap of libraries
		rand.Seed(time.Now().UnixNano())
		rand1 := rand.Intn(len(secondGen))
		rand2 := rand.Intn(len(secondGen))
		ans := secondGen[rand1]
		secondGen[rand1] = secondGen[rand2]
		secondGen[rand2] = ans

		// Recalculate library end time
		secondGen = p.recalculateLibrary(secondGen)

		// Reassign books
		secondGen = p.assignBooks(secondGen)

		// Calculate energy/score
		firstScore := p.calcScoreBase(firstGen)
		secondScore := p.calcScoreBase(secondGen)

		// Decide if we should accept with acceptance probability
		rand.Seed(time.Now().UnixNano())
		x := (tempStart / tempStart1) / 2
		// x := rand.Float64()
		y := acceptanceProbability(firstScore, secondScore, tempStart)

		if y < x {
			firstGen = secondGen
		}

		// Record new population if score better than previous
		if firstScore > bestGenScore {
			bestGen = secondGen
			bestGenScore = firstScore
		}

		// Cool the temperature
		tempStart *= 1 - coolingRate
	}

	// Print result of best generation
	fmt.Println("Best generation of", p.filePath, "has a score of:", p.calcScoreBase(bestGen))
	p.answers = bestGen
}

// Calculate acceptance probability
func acceptanceProbability(energy, newEnergy int, temperature float64) float64 {
	if newEnergy > energy {
		return 1
	}

	// fmt.Println(energy, newEnergy, temperature)
	// fmt.Println("Acceptance probability:", math.Exp(float64(newEnergy-energy)/temperature), "************")
	// return math.Exp(float64(newEnergy-energy) / temperature)
	return math.Exp(float64(newEnergy-energy)/temperature) / 1
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

	// Assign books in the assigned libraries
	newAns = p.assignBooks(newAns)

	return
}

func (p *problem) assignBooks(newAns []answer) []answer {
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

		// Reset book answer
		newAns[k].booksAns = []book{}

		// Assign books
		for j := range newAns[k].books {
			if _, ok := plannedBooks[newAns[k].books[j].ID]; !ok && allowedBooks > 0 {
				plannedBooks[newAns[k].books[j].ID] = struct{}{}
				newAns[k].booksAns = append(newAns[k].booksAns, newAns[k].books[j])
				allowedBooks--
			}
		}
	}

	return newAns
}

func (p *problem) recalculateLibrary(newAns []answer) []answer {
	lastDay := 0

	for k := range newAns {
		lastDay += newAns[k].signUpDuration
		newAns[k].signUpEndDay = lastDay
	}

	return newAns
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
