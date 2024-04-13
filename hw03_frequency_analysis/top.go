package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	words := strings.Fields(text)
	topFrequencies := make(map[string]int)
	for _, word := range words {
		if word == "" {
			continue
		}
		if _, ok := topFrequencies[word]; ok {
			topFrequencies[word]++
		} else {
			topFrequencies[word] = 1
		}
	}

	keys := make([]string, 0, len(topFrequencies))
	for k := range topFrequencies {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		if topFrequencies[keys[i]] == topFrequencies[keys[j]] {
			return keys[i] < keys[j]
		}
		return topFrequencies[keys[i]] > topFrequencies[keys[j]]
	})

	sorted := make([]string, 0, 10)
	for i, word := range keys {
		if i == 10 {
			break
		}
		sorted = append(sorted, word)
	}
	return sorted
}
