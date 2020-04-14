package arrays

import (
	"sort"
	"strings"
)

func Permutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	x1 := sorted(s1)
	x2 := sorted(s2)

	return x1 == x2
}

func sorted(str string) string {
	s := strings.Split(str, "")

	sort.Strings(s)

	return strings.Join(s, " ")
}
