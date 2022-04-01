package exercise

func FirstRepeatedChar(str string) (char byte) {
	mapper := make(map[byte]bool)

	for i := 0; i < len(str); i++ {
		character := str[i]
		if _, ok := mapper[character]; ok {
			char = character
			return
		} else {
			mapper[character] = true
		}
	}

	return
}
