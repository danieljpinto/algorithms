package arrays

import (
	"sort"
	"strings"
)

func Permutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	return sorted(s1) == sorted(s2)
}

func sorted(str string) string {
	s := strings.Split(str, "")

	sort.Strings(s)

	return strings.Join(s, " ")
}

func Permutation2(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	letters := make([]int, 128) // Consider ASCII

	for i := 0; i < len(s1); i++ {
		letters[s1[i]]++
	}

	for i := 0; i < len(s2); i++ {
		chr := s2[i]

		letters[chr]--

		if letters[chr] < 0 {
			return false
		}
	}

	return true
}
