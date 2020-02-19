package main

import (
	"math"
	"sort"

	"github.com/ttimt/GoogleHashCode-2020-Qualification/stdlib"
)

// Main algorithm
//
// To sort p.data ID ascending :
// sort.Slice(p.data, func(i, j int) bool {
// 	return p.data[i].ID < p.data[j].ID
// })
//
func (p *problem) algorithm1() {
	// Find bonus
	for k := range p.data {
		if !p.data[k].assigned {
			// If empty car
			for j := range p.answers {
				if len(p.answers[j].rides) == 0 {
					if p.data[k].rowStart+p.data[k].columnStart+p.data[k].earliestStart <= p.onTimeBonus {
						p.answers[j].rides = append(p.answers[j].rides, &p.data[k])
						p.data[k].assigned = true
						p.data[k].start = stdlib.MaxInt(p.data[k].earliestStart, p.data[k].rowStart+p.data[k].columnStart)
						p.data[k].end = p.data[k].start + p.data[k].getDistance()
						break
					}
				} else if p.answers[j].rides[len(p.answers[j].rides)-1].end+p.data[k].getDistance() <= p.data[k].latestFinish &&
					int(math.Abs(float64(p.data[k].rowStart-p.answers[j].rides[len(p.answers[j].rides)-1].rowEnd))+float64(p.data[k].columnStart-p.answers[j].rides[len(p.answers[j].rides)-1].columnEnd)+float64(p.data[k].earliestStart)) <= p.onTimeBonus {
					// If does not exceed latest finish
					p.answers[j].rides = append(p.answers[j].rides, &p.data[k])
					p.data[k].assigned = true
					p.data[k].start = stdlib.MaxInt(p.data[k].earliestStart,
						(p.data[k].rowStart-p.answers[j].rides[len(p.answers[j].rides)-1].rowEnd)+(p.data[k].columnStart-p.answers[j].rides[len(p.answers[j].rides)-1].columnEnd)+p.answers[j].rides[len(p.answers[j].rides)-1].end)
					p.data[k].end = p.data[k].start + p.data[k].getDistance()
					break
				}
			}
		}

		// Sort vehicle
		sort.Slice(p.answers, func(i, j int) bool {
			if len(p.answers[i].rides) == 0 {
				return true
			} else if len(p.answers[j].rides) == 0 {
				return false
			}

			return p.answers[i].rides[len(p.answers[i].rides)-1].end < p.answers[j].rides[len(p.answers[j].rides)-1].end
		})
	}

	for k := range p.data {
		if !p.data[k].assigned {

			// If empty car
			for j := range p.answers {
				if len(p.answers[j].rides) == 0 {
					p.answers[j].rides = append(p.answers[j].rides, &p.data[k])
					p.data[k].assigned = true
					p.data[k].start = stdlib.MaxInt(p.data[k].earliestStart, p.data[k].rowStart+p.data[k].columnStart)
					p.data[k].end = p.data[k].start + p.data[k].getDistance()
					break
				} else if p.answers[j].rides[len(p.answers[j].rides)-1].end+p.data[k].getDistance() <= p.data[k].latestFinish {
					// If does not exceed latest finish
					p.answers[j].rides = append(p.answers[j].rides, &p.data[k])
					p.data[k].assigned = true
					p.data[k].start = stdlib.MaxInt(p.data[k].earliestStart,
						int(math.Abs(float64(p.data[k].rowStart-p.answers[j].rides[len(p.answers[j].rides)-1].rowEnd))+float64(p.data[k].columnStart-p.answers[j].rides[len(p.answers[j].rides)-1].columnEnd+p.answers[j].rides[len(p.answers[j].rides)-1].end)))
					p.data[k].end = p.data[k].start + p.data[k].getDistance()
					break
				}
			}

			// Sort vehicle
			sort.Slice(p.answers, func(i, j int) bool {
				if len(p.answers[i].rides) == 0 {
					return true
				} else if len(p.answers[j].rides) == 0 {
					return false
				}

				return p.answers[i].rides[len(p.answers[i].rides)-1].end < p.answers[j].rides[len(p.answers[j].rides)-1].end
			})
		}
	}
}

// Secondary algorithm
//
func (p *problem) algorithm2() {
	// Generate rides
	for i := 0; i < p.nrOfVehicles; i++ {
		p.answers = append(p.answers, answer{rides: []*problemData{}})
	}

	// Sort by earliest start
	sort.Slice(p.data, func(i, j int) bool {
		return p.data[i].earliestStart < p.data[j].earliestStart
	})

	p.algorithm1()
	p.printAnswer()
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
		for j := range answers[k].rides {
			score += int(math.Abs(float64(answers[k].rides[j].rowEnd-answers[k].rides[j].rowStart)) + math.Abs(float64(answers[k].rides[j].columnEnd-answers[k].rides[j].columnStart)))

			if answers[k].rides[j].start == answers[k].rides[j].earliestStart {
				score += p.onTimeBonus
			}
		}
	}

	return score
}

func (d *problemData) getDistance() int {
	return getDistance(d.rowStart, d.columnStart, d.rowEnd, d.columnEnd)
}

func getDistance(fromRow, fromColumn, toRow, toColumn int) int {
	return int(math.Abs(float64(toRow-fromRow)) + float64(toColumn-fromColumn))
}
