package regularexpressionmatching

// IsMatch checks if string s matched pattern p
func IsMatch(s string, p string) bool {

	if len(p) == 0 {
		return len(s) == 0
	}

	currPairMatched := (len(s) > 0) && (p[0] == s[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		return IsMatch(s, p[2:]) || (currPairMatched && IsMatch(s[1:], p))
	}

	return currPairMatched && IsMatch(s[1:], p[1:])
}
