package main

import (
	"strings"
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
	// Create map of tags and node list
	tagMap := make(map[string]*tagWrap)

	for k := range p.dataWVertical {
		for j := range p.dataWVertical[k].tags {
			if _, ok := tagMap[j]; !ok {
				tagMap[j] = initializeNewTag(&p.dataWVertical[k])
			} else {
				// If tag exist before
				appendNewTag(&p.dataWVertical[k], tagMap[j])
			}
		}
	}
	// END

	// Assign to slide show
	assigned := make(map[int]struct{})

	// Store answer
	var currentPhoto *problemData

	for i := 0; i < len(p.dataWVertical); i++ {
		if _, ok := assigned[p.dataWVertical[i].ID]; !ok {
			currentPhoto = &p.dataWVertical[i]
			p.answers = append(p.answers, answer{&p.dataWVertical[i]})
			assigned[p.dataWVertical[i].ID] = struct{}{}

			p.solve(currentPhoto, assigned, tagMap)
		}
	}

}

func (p *problem) solve(currentPhoto *problemData, assigned map[int]struct{}, tagMap map[string]*tagWrap) {
	// Get max score
	var maxPhoto *problemData
	var maxScore int
	var samePhotos []*problemData
	searchedPhotos := make(map[int]struct{})
	currentPhotoMaxScore := currentPhoto.nrOfTags / 2

	for j := range currentPhoto.tags {
		samePhotos = getPhotosInTag(currentPhoto, tagMap[j])

		for h := range samePhotos {
			if _, ok := assigned[samePhotos[h].ID]; !ok {
				if _, ok := searchedPhotos[samePhotos[h].ID]; !ok {
					newScore := calcScoreBetweenTwo(answer{currentPhoto}, answer{samePhotos[h]})
					if newScore > maxScore {
						maxScore = newScore
						maxPhoto = samePhotos[h]
					}

					if maxScore >= currentPhotoMaxScore {
						break
					}

					searchedPhotos[samePhotos[h].ID] = struct{}{}
				}
			}
		}

		if maxScore >= currentPhotoMaxScore {
			break
		}
	}

	// Assign to the photo with max score
	if maxPhoto != nil {
		assigned[maxPhoto.ID] = struct{}{}
		p.answers = append(p.answers, answer{maxPhoto})

		// Start on the assigned photo
		currentPhoto = maxPhoto
		p.solve(currentPhoto, assigned, tagMap)
	}
}

func getPhotosInTag(photo *problemData, tw *tagWrap) (photos []*problemData) {
	photoTag := tw.start.next
	for photoTag != tw.end {
		if photoTag.photo != photo {
			photos = append(photos, photoTag.photo)
		}

		photoTag = photoTag.next
	}

	return
}

func appendNewTag(p *problemData, tw *tagWrap) {
	tag := &tag{
		photo:    p,
		previous: tw.end.previous,
		next:     tw.end,
	}

	tw.end.previous.next = tag
	tw.end.previous = tag
}

func initializeNewTag(p *problemData) *tagWrap {
	startTag := &tag{
		photo:    nil,
		previous: nil,
		next:     nil,
	}

	endTag := &tag{
		photo:    nil,
		previous: nil,
		next:     nil,
	}

	tag := &tag{
		photo:    p,
		previous: startTag,
		next:     endTag,
	}

	startTag.next = tag
	endTag.previous = tag

	tagWrap := &tagWrap{
		start: startTag,
		end:   endTag,
	}

	return tagWrap
}

// Assign vertical
func (p *problem) assignVertical() {
	// Store unassigned vertical photos
	var singleVertical []problemData

	for _, v := range p.data {
		if strings.ToUpper(v.orientation) != "V" {
			p.dataWVertical = append(p.dataWVertical, v)
		} else {
			singleVertical = append(singleVertical, v)
		}
	}

	// Process vertical
	for i := 0; i < len(singleVertical); i += 2 {
		if i+1 < len(singleVertical) {
			appendVertical(&singleVertical[i], &singleVertical[i+1])
			p.dataWVertical = append(p.dataWVertical, singleVertical[i])
		}
	}

	return
}

func appendVertical(v1, v2 *problemData) {
	for k := range v2.tags {
		if _, ok := v1.tags[k]; !ok {
			v1.tags[k] = struct{}{}
			v1.nrOfTags++
		}
	}

	v1.photosID = append(v1.photosID, v1.ID, v2.ID)
}

// Default recursive algorithm
//
func (p *problem) recursive(data, curData []problemData, curPD problemData, maxData []answer, maxScore *int, currentScore int) []answer {
	// Return if max reached
	if true { // *maxScore == p.maxPizzaSlices
		return maxData
	}

	// Add current curPD value if still within range
	if true { // Ex:curPD.nrOfSlices+currentScore <= p.maxPizzaSlices
		// currentScore += curPD.nrOfSlices
		curData = append(curData, curPD)
	}

	// End if data ends
	if len(data) <= 1 {
		// Update max score
		if currentScore > *maxScore {
			*maxScore = currentScore

			var newMax []answer
			for k := range curData {
				newMax = append(newMax, answer{problemData: &curData[k]})
			}

			return newMax
		}

		// Output to preserve current max score if recursive takes too long time
		// if *maxScore > 999999995 {
		// 	p.writeFile()
		// }

		return maxData
	}

	// Recursive
	for k := range data[1:] {
		maxData = p.recursive(data[k+1:], curData, data[k+1], maxData, maxScore, currentScore)
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
func calcScore(answers []answer) int {
	score := 0

	for k := range answers[1:] {
		score += calcScoreBetweenTwo(answers[k], answers[k+1])
	}

	return score
}

func calcScoreBetweenTwo(p, p1 answer) int {
	pCount := p.nrOfTags
	p1Count := p1.nrOfTags
	overlap := 0

	for k := range p.tags {
		if _, ok := p1.tags[k]; ok {
			overlap++
		}
	}

	return min(pCount, p1Count, overlap)
}

func min(i ...int) int {
	lowest := i[0]

	for _, v := range i[1:] {
		if v < lowest {
			lowest = v
		}
	}

	return lowest
}
