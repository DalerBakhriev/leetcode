package lettercombinations

var numbers map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var allSolutions []string

func cleaArr()  {
	allSolutions = make([]string, 0)
}

// LetterCombinations ...
func LetterCombinations(digits string) []string {

	defer cleaArr()
	if len(digits) != 0 {
		backtrack("", digits)
	}

	return allSolutions
}

func backtrack(combination string, currDigits string) []string {

	if len(currDigits) == 0 {
		allSolutions = append(allSolutions, combination)
	} else {
		digit := currDigits[0:1]
		candidates := numbers[digit]

		for i := 0; i < len(candidates); i++ {
			letter := candidates[i : i+1]
			backtrack(combination+letter, currDigits[1:])
		}
	}

	return allSolutions
}
