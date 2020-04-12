package generateparentheses

func insertInside(str string, leftIndex int) string {

	leftPart := str[:leftIndex+1]
	rightPart := str[leftIndex+1:]

	return leftPart + "()" + rightPart
}

func generateParenthesis(n int) map[string]int {

	solutions := make(map[string]int)
	if n == 0 {
		solutions[""] = 1
	} else {
		prevSolution := generateParenthesis(n - 1)

		for str := range prevSolution {
			for ind := range str {
				if str[ind] == '(' {
					solution := insertInside(str, ind)
					solutions[solution] = 1
				}
			}
			solutions["()"+str] = 1
		}
	}

	return solutions
}

// GenerateParenthesis ...
func GenerateParenthesis(n int) []string {
	solutions := generateParenthesis(n)
	var finalSolutions []string

	for solution := range solutions {
		if solution != "" {
			finalSolutions = append(finalSolutions, solution)
		}
	}

	return finalSolutions
}
