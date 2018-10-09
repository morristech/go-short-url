package main

func GenerateHash(seed interface{}) string {

	series := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	}

	seedStr := seed.(string)
	charsCount := len(seedStr)

	if len(seedStr) == 0 {
		return series[0]
	}

	lastChar := string(seedStr[charsCount-1])

	// find last char index in the series
	// this will help us in finding the next char from series
	i := 0
	for k, v := range series {
		if lastChar == v {
			i = k
			break
		}
	}

	// if the last has reached the end of series then append first char
	if lastChar == "z" {
		return seedStr + series[0]
	}

	// append the next char from series to the last hash-1
	return seedStr[:charsCount-1] + series[i+1]
}
