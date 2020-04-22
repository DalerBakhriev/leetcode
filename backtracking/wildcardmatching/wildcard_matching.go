package wildcardmatching

func isMatch(s string, p string) bool {

	if len(p) == 0 {
		return len(s) == 0
	}

	dp := make([]bool, len(p)+1)
	dp[0] = true
	for i := 1; i <= len(p); i++ {
		if p[i-1] == '*' {
			dp[i] = dp[i-1]
		}
	}

	last := dp[0]
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				temp := dp[j]
				dp[j] = dp[j-1] || dp[j]
				last = temp
			} else {
				temp := dp[j]
				dp[j] = last && (p[j-1] == '?' || s[i-1] == p[j-1])
				last = temp
			}
		}

		last = false
	}

	return dp[len(dp)-1]
}
