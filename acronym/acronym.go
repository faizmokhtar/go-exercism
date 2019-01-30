// Package acronym implements a simple library to create acronym
package acronym

import (
	"bytes"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/rangetable"
)

// Abbreviate a string into an acronym
func Abbreviate(s string) string {
	var acronym bytes.Buffer
	words := strings.FieldsFunc(s, isSeperator)

	for _, w := range words {
		char := strings.ToUpper(w[0:1])
		acronym.WriteString(char)
	}
	return acronym.String()
}

func isSeperator(c rune) bool {
	r := rangetable.New('-')
	return unicode.IsSpace(c) || unicode.Is(r, c)
}
