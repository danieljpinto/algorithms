package arrays

func HasUniqueChars(str string) bool {
	if len(str) > 128 {
		return false
	}

	// Consider ASCII
	charSet := make([]bool, 128)

	for i := 0; i < len(str); i++ {
		val := str[i]

		if charSet[val] {
			return false
		}

		charSet[val] = true
	}

	return true
}
