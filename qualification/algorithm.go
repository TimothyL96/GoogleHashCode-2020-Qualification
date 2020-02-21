package main

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
//
func (p *problem) algorithm2() {

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
