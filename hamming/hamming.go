package hamming

import "errors"

func Distance(a, b string) (int, error) {
	var (
		distance = 0
		runeA    = []rune(a)
		runeB    = []rune(b)
	)
	if len(runeA) != len(runeB) {
		return 0, errors.New("Strands don't have equal length")
	}

	for i, rune := range runeA {
		if rune != runeB[i] {
			distance++
		}
	}
	return distance, nil
}
