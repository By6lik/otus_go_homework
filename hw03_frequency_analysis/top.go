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
	if len(t) == 0 {
		result = []string{}
	} else {
		text := strings.Join(strings.Fields(t), " ")
		splitedText := strings.Split(text, " ")
		A := []WordCount{}
		for word := range splitedText {
			if tempMap[splitedText[word]] == 0 {
				tempMap[splitedText[word]] = 1
			} else {
				tempMap[splitedText[word]]++
			}
		}
		for i, j := range tempMap {
			A = append(A, WordCount{i, j})
		}
		sort.Slice(A, func(i, j int) bool { return A[i].count > A[j].count })
		A = A[0:10]
		sort.Slice(A, func(i, j int) bool {
			if A[i].count == A[j].count {
				return A[i].expression < A[j].expression
			}
			return false
		})
		for i := range A {
			result = append(result, A[i].expression)
		}
	}

	return result
}
