package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type WordCount struct {
	expression string
	count      int
}

func ReplaceSpace(t string) string {
	re := regexp.MustCompile(" - ")
	slug := re.ReplaceAllString(t, " - ")
	return slug
}

func Top10(t string) []string {
	result := []string{}
	tempMap := map[string]int{}
	if len(t) > 0 {
		text := strings.Join(strings.Fields(t), " ")
		splitedText := strings.Split(text, " ")
		w := []WordCount{}
		for word := range splitedText {
			if tempMap[splitedText[word]] == 0 {
				tempMap[splitedText[word]] = 1
			} else {
				tempMap[splitedText[word]]++
			}
		}
		for i, j := range tempMap {
			w = append(w, WordCount{i, j})
		}
		sort.Slice(w, func(i, j int) bool { return w[i].count > w[j].count })
		switch {
		case len(w) < 10:
			w = w[0:]
		default:
			w = w[0:10]
		}
		sort.Slice(w, func(i, j int) bool {
			if w[i].count == w[j].count {
				return w[i].expression < w[j].expression
			}
			return false
		})
		for i := range w {
			result = append(result, w[i].expression)
		}
	} else {
		result = []string{}
	}
	return result
}
