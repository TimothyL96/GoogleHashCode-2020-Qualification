package main

import (
	"fmt"
	"sort"
)

// Main algorithm
//
// To sort p.data ID ascending :
// sort.Slice(p.data, func(i, j int) bool {
// 	return p.data[i].ID < p.data[j].ID
// })
//
func (p *problem) algorithm1(filePath string) {
	assignedSlice := 0
	cachePrevious := 0
	cachePreviousPosition := 0

	for k := range p.data {
		if assignedSlice+p.data[k].nrOfSlices <= p.maxPizzaSlices {
			assignedSlice += p.data[k].nrOfSlices
			p.answers = append(p.answers, answer{problemData: &p.data[k]})
			cachePrevious = p.data[k].nrOfSlices
			cachePreviousPosition = len(p.answers) - 1
			p.data[k].assigned = true
		} else if p.data[k].nrOfSlices > cachePrevious && assignedSlice-cachePrevious+p.data[k].nrOfSlices <= p.maxPizzaSlices {
			assignedSlice -= cachePrevious
			assignedSlice += p.data[k].nrOfSlices
			cachePrevious = p.data[k].nrOfSlices
			p.answers[cachePreviousPosition].problemData.assigned = false
			p.answers[cachePreviousPosition] = answer{problemData: &p.data[k]}
			p.data[k].assigned = true
		} else if assignedSlice >= p.maxPizzaSlices {
			break
		}
	}

	currentAssignedSlice := assignedSlice
	for true {
		for k := range p.data {
			if !p.data[k].assigned {
				for j := range p.answers {
					if assignedSlice-p.answers[j].nrOfSlices+p.data[k].nrOfSlices <= p.maxPizzaSlices && p.answers[j].nrOfSlices < p.data[k].nrOfSlices {
						p.data[k].assigned = true
						p.answers[j].problemData.assigned = false
						assignedSlice -= p.answers[j].nrOfSlices
						assignedSlice += p.data[k].nrOfSlices
						p.answers[j].problemData = &p.data[k]
						break
					}
				}
			}
		}

		if currentAssignedSlice == assignedSlice {
			break
		} else {
			currentAssignedSlice = assignedSlice
		}
	}

	var pd []problemData
	for k := range p.data {
		if !p.data[k].assigned {
			pd = append(pd, p.data[k])
		}
	}

	// Sort
	sort.Slice(pd, func(i, j int) bool {
		return pd[i].ID < pd[j].ID
	})
	sort.Slice(p.answers, func(i, j int) bool {
		return p.answers[i].ID < p.answers[j].ID
	})

	// Find sum of at least 2 answers from lowest to be bigger than unassigned pieces
	if len(pd) >= 2 {
		minSlice := p.answers[0].nrOfSlices + p.answers[1].nrOfSlices

		for k := range pd {
			if len(pd) > k && pd[k].nrOfSlices < minSlice {
				if len(pd) > k+1 {
					pd = append(pd[:k], pd[k+1:]...)
				} else {
					pd = pd[:k]
				}
			}
		}

		for k := range pd {
			for j := range p.answers[1:] {
				if len(p.answers) > j+1 &&
					p.answers[j].assigned &&
					p.answers[j+1].assigned &&
					assignedSlice-p.answers[j].nrOfSlices-p.answers[j+1].nrOfSlices+pd[k].nrOfSlices <= p.maxPizzaSlices &&
					pd[k].nrOfSlices >= p.answers[j].nrOfSlices+p.answers[j+1].nrOfSlices {
					assignedSlice += pd[k].nrOfSlices
					assignedSlice -= p.answers[j].nrOfSlices
					assignedSlice -= p.answers[j+1].nrOfSlices

					p.answers[j].assigned = false
					p.answers[j+1].assigned = false
					p.answers[j].problemData = &pd[k]
					p.answers[j].assigned = true
					if len(p.answers) > j+2 {
						p.answers = append(p.answers[:j+1], p.answers[j+2:]...)
					} else {
						p.answers = p.answers[:j+1]
					}
					break
				}
			}
		}

		for k := range p.data {
			if !p.data[k].assigned {
				pd = append(pd, p.data[k])
			}
		}

		// Sort
		sort.Slice(pd, func(i, j int) bool {
			return pd[i].ID < pd[j].ID
		})

		if len(pd) >= 2 && len(p.answers) > 2 {
			minSlice := p.answers[0].nrOfSlices + p.answers[1].nrOfSlices

			for k := range pd {
				if len(pd) > k && pd[k].nrOfSlices < minSlice {
					if len(pd) > k+1 {
						pd = append(pd[:k], pd[k+1:]...)
					} else {
						pd = pd[:k]
					}
				}
			}
			for k := range pd {
				for j := range p.answers[2:] {
					if len(p.answers) > j+2 &&
						p.answers[j].assigned &&
						p.answers[j+1].assigned &&
						p.answers[j+2].assigned &&
						assignedSlice-p.answers[j].nrOfSlices-p.answers[j+1].nrOfSlices-p.answers[j+2].nrOfSlices+pd[k].nrOfSlices <= p.maxPizzaSlices &&
						pd[k].nrOfSlices >= p.answers[j].nrOfSlices+p.answers[j+1].nrOfSlices+p.answers[j+2].nrOfSlices {
						assignedSlice += pd[k].nrOfSlices
						assignedSlice -= p.answers[j].nrOfSlices
						assignedSlice -= p.answers[j+1].nrOfSlices
						assignedSlice -= p.answers[j+2].nrOfSlices

						p.answers[j].assigned = false
						p.answers[j+1].assigned = false
						p.answers[j+2].assigned = false
						p.answers[j].problemData = &pd[k]
						p.answers[j].assigned = true
						if len(p.answers) > j+3 {
							p.answers = append(p.answers[:j+1], p.answers[j+3:]...)
						} else {
							p.answers = p.answers[:j+1]
						}
						break
					}
				}
			}
		}
	}
}

func (p *problem) algorithm2() {
	sort.Slice(p.data, func(i, j int) bool {
		return p.data[i].nrOfSlices >= p.data[j].nrOfSlices
	})

	// Run recursive in a loop
	maxScore := 0
	ans := p.recursive(p.data, make([]problemData, 0), p.data[0], make([]answer, 0), &maxScore, 0)
	p.answers = ans
}

// Default recursive algorithm
func (p *problem) recursive(data, curData []problemData, curPD problemData, maxData []answer, maxScore *int, currentScore int) []answer {
	if *maxScore > 999999995 {
		return maxData
	}

	if curPD.nrOfSlices+currentScore <= p.maxPizzaSlices {
		currentScore += curPD.nrOfSlices
		curData = append(curData, curPD)
	}
	// fmt.Println("Cur data", curPD.ID, "slice:", curPD.nrOfSlices)
	fmt.Println("cur max score", *maxScore)
	if len(data) <= 1 {
		if currentScore > *maxScore {
			*maxScore = currentScore
			var newMax []answer
			for k := range curData {
				newMax = append(newMax, answer{problemData: &curData[k]})
			}
			return newMax
		}

		if *maxScore > 999999995 {
			p.writeFile()
		}
		return maxData
	}

	for k := range data[1:] {
		maxData = p.recursive(data[k+1:], curData, data[k+1], maxData, maxScore, currentScore)
	}

	return maxData
}

// Calculate answers score and store result in p.score
// Access answer struct with p.answers (type is a slice of answer)
func (p *problem) calcScore() {
	p.score = 0

	for k := range p.answers {
		p.score += p.answers[k].nrOfSlices
	}
}
