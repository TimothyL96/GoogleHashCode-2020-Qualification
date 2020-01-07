package main

// Main algorithm
func (p *problem) algorithm1() {
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
}

// Calculate answers score and store result in p.score
// Access answer struct with p.answers (type is a slice of answer)
func (p *problem) calcScore() {
	p.score = 0

	for k := range p.answers {
		p.score += p.answers[k].nrOfSlices
	}
}
