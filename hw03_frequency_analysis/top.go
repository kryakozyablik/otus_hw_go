package hw03_frequency_analysis // nolint:golint,stylecheck

import (
	"regexp"
	"sort"
	"strings"
)

type wordStat struct {
	word  string
	count int
}

func Top10(text string) []string {
	wordsCount := wordCount(text)
	wordsStat := buildSortedWordsStat(wordsCount)
	returnWordCnt := min(len(wordsStat), 10)
	topWords := make([]string, 0, returnWordCnt)

	for i := 0; i < returnWordCnt; i++ {
		topWords = append(topWords, wordsStat[i].word)
	}

	return topWords
}

func wordCount(text string) map[string]int {
	text = strings.ToLower(text)
	text = regexp.MustCompile(`(?i)[^a-zа-я0-9\- ]`).ReplaceAllString(text, " ")
	words := strings.Split(text, " ")
	wordsCount := make(map[string]int, len(words))

	for _, word := range words {
		if word == "" || word == "-" {
			continue
		}
		wordsCount[word]++
	}

	return wordsCount
}

func buildSortedWordsStat(wordsCount map[string]int) []wordStat {
	ws := make([]wordStat, 0, len(wordsCount))

	for w, c := range wordsCount {
		ws = append(ws, wordStat{word: w, count: c})
	}

	sort.Slice(ws, func(i, j int) bool { return ws[i].count > ws[j].count })

	return ws
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
