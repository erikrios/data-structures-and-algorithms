package exercise

func FirstNonRepeatingChar(str string) (firstChar byte) {
	mapper := make(map[byte]int)

	for i := 0; i < len(str); i++ {
		char := str[i]
		count := mapper[char]
		mapper[char] = count + 1
	}

	for i := 0; i < len(str); i++ {
		char := str[i]
		if v := mapper[char]; v == 1 {
			firstChar = char
			return
		}
	}

	return
}
