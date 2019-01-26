package bob

import (
	"strings"
)

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if isFine(remark) {
		return "Fine. Be that way!"
	} else if isQuestion(remark) {
		return "Sure"
	} else if isYelling(remark) {
		return "Whoa, chill out!"
	}

	return "Whatever."
}

func isFine(r string) bool {
	return len(r) == 0
}

func isQuestion(r string) bool {
	return r[:len(r)] == "?"
}

func isYelling(r string) bool {
	return r == strings.ToUpper(r) && r != strings.ToLower(r)
}
