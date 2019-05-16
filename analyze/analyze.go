package analyze

import (
	"gopkg.in/jdkato/prose.v2"
	"log"
	"sort"
	"strings"
	"time"
)

var stopWords = map[string]bool{
	"Commission": true, "Committee": true, "Committees": true, "Code": true, "II": true, "III": true, "IX": true,
	"United States": true, "Congress": true, "Joint Ad Hoc Committee": true, "House": true, "Senate": true, "Federal": true,
	"State": true, "Board": true}
var stopCharacters = []string{"(", ")"}

type LegislationItem struct {
	Text       string
	Count      int
	Bills      []string
	Sponsors   []string
	TotalCount int
}

func LegislativeText(legislativeText []byte) ([]LegislationItem, error) {
	log.Println("Performing NLP")
	startTime := time.Now()
	nlpResults, err := prose.NewDocument(string(legislativeText))
	if err != nil {
		return nil, err
	}
	executionTime := time.Since(startTime)
	log.Printf("NLP took %s\n", executionTime)

	log.Println("Pivoting and sorting results")
	startTime = time.Now()
	var entities = make(map[string]int)
	for _, ent := range nlpResults.Entities() {
		sanitized := sanitizeText(ent.Text)
		if sanitized != "" {
			if val, ok := entities[sanitized]; ok {
				entities[sanitized] = val + 1
			} else {
				entities[sanitized] = 1
			}
		}
	}

	legislationItems := make([]LegislationItem, 0, 0)
	for k, v := range entities {
		legislationItems = append(legislationItems, LegislationItem{k, v, []string{}, []string{}, -1})
	}
	sort.Slice(legislationItems, func(i, j int) bool {
		return legislationItems[i].Count > legislationItems[j].Count
	})
	executionTime = time.Since(startTime)
	log.Printf("Pivoting and sorting took %s\n", executionTime)

	return legislationItems, nil
}

func sanitizeText(text string) string {
	text = strings.TrimSpace(text)
	if stopWords[text] {
		return ""
	}
	for i := range stopCharacters {
		if strings.Index(text, stopCharacters[i]) > -1 {
			return ""
		}
	}
	text = strings.Replace(text, "Whereas", "", -1)

	return text
}
