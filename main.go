package main

import (
	"fmt"

	"github.com/danieljpinto/algorithms/arrays"
)

func main() {

	fmt.Printf("arrays.HasUniqueChars for %q is %v\n", "danielzz", arrays.HasUniqueChars("danielzz"))
	fmt.Printf("arrays.HasUniqueChars for %q is %v\n", "daniel", arrays.HasUniqueChars("daniel"))

	fmt.Printf("arrays.Permutation    for %q and %q is %v\n", "god", "dog", arrays.Permutation("god", "dog"))
	fmt.Printf("arrays.Permutation    for %q and %q is %v\n", "neo", "one", arrays.Permutation("neo", "one"))
	fmt.Printf("arrays.Permutation    for %q and %q is %v\n", "orange", "egnaso", arrays.Permutation("orange", "egnaso"))
}
