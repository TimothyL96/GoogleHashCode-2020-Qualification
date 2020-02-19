package main

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
	// Generate vehicles
	for i := 0; i < p.nrOfVehicles; i++ {
		p.answers = append(p.answers, answer{vehicles: []*problemData{}})
	}
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
		for j := range answers[k].vehicles {
			score += (answers[k].vehicles[j].rowEnd - answers[k].vehicles[j].rowStart) + (answers[k].vehicles[j].columnEnd - answers[k].vehicles[j].columnStart)

			if answers[k].vehicles[k].start == answers[k].vehicles[j].earliestStart {
				score += p.onTimeBonus
			}
		}
	}

	return score
}
