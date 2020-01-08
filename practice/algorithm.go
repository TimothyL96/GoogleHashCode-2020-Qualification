package main

import (
	"fmt"
	"sort"
)

// Main algorithm
//
// To sort:
// sort.Slice(p.data, func(i, j int) bool {
// 	return p.data[i].ID < p.data[j].ID
// })
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

	newCur := p.data[1:]
	maxScore, ans := p.recursive(0, 0, &p.data[0], newCur, make([]answer, 0), make([]answer, 0))
	fmt.Println("Max score:", maxScore)
	p.answers = ans
}

func (p *problem) recursive(currentScore int, maxScore int, pd *problemData, cur []problemData, curans, ans []answer) (int, []answer) {
	if currentScore+pd.nrOfSlices <= p.maxPizzaSlices {
		currentScore += pd.nrOfSlices
		curans = append(ans, answer{pd})

		if currentScore > maxScore {
			ans = curans
			maxScore = currentScore

			fmt.Println("Current max score", maxScore)
			if maxScore == p.maxPizzaSlices {
				return maxScore, ans
			}
		}
	} else {
		return maxScore, ans
	}

	if cur == nil {
		return maxScore, ans
	}

	for k := range cur {
		newCur := []problemData(nil)
		if len(cur) <= k+1 {
			break
		}
		newCur = (cur)[k+1:]
		maxScore, ans = p.recursive(currentScore, maxScore, &(cur)[k], newCur, curans, ans)

		if currentScore > maxScore {
			ans = curans
			maxScore = currentScore

			fmt.Println("Current max score 1", maxScore)
		}

		if maxScore == p.maxPizzaSlices {
			break
		}
	}

	return maxScore, ans
}

// Calculate answers score and store result in p.score
// Access answer struct with p.answers (type is a slice of answer)
func (p *problem) calcScore() {
	p.score = 0

	for k := range p.answers {
		p.score += p.answers[k].nrOfSlices
	}
}
