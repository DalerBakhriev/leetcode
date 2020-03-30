package regularexpressionmatching

import "regexp"

func isMatch(s string, p string) bool {

	patternForFull := "^" + p + "$"
	matched, _ := regexp.MatchString(patternForFull, s)

	return matched
}
