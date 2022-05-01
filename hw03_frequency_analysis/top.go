package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

const (
	limit = 10
)

func Top10(inStr string) []string {
	words := strings.Fields(inStr)

	freqMap := make(map[string]int64)
	for _, w := range words {
		freqMap[w]++
	}

	uniqWords := make([]string, 0, len(freqMap))
	for w := range freqMap {
		uniqWords = append(uniqWords, w)
	}

	sort.Slice(uniqWords, func(i, j int) bool {
		if freqMap[uniqWords[i]] > freqMap[uniqWords[j]] {
			return true
		}
		if freqMap[uniqWords[i]] < freqMap[uniqWords[j]] {
			return false
		}

		return uniqWords[i] < uniqWords[j]
	})

	count := limit
	if count > len(uniqWords) {
		count = len(uniqWords)
	}

	return uniqWords[:count]
}
