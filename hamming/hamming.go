package hamming

import "errors"

// Distance calculate the Hamming distance between two strands
func Distance(a, b string) (int, error) {
	runeA := []rune(a)
	runeB := []rune(b)

	if len(runeA) != len(runeB) {
		return 0, errors.New("strands don't have equal length")
	}

	distance := 0
	for i, r := range runeA {
		if r != runeB[i] {
			distance++
		}
	}
	return distance, nil
}
