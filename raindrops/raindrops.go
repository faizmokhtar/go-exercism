package raindrops

import (
	"bytes"
	"strconv"
)

// Convert a number to a fun string based on its factorial
func Convert(n int) string {
	var b bytes.Buffer
	nums := factorial(n)
	for _, num := range nums {
		switch num {
		case 3:
			b.WriteString("Pling")
		case 5:
			b.WriteString("Plang")
		case 7:
			b.WriteString("Plong")
		}
	}

	if s := b.String(); s != "" {
		return s
	}

	return strconv.Itoa(nums[len(nums)-1])
}

func factorial(n int) (s []int) {
	s = make([]int, 0)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			s = append(s, i)
		}
	}
	return
}
