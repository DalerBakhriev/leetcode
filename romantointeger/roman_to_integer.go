package romantointeger

var romanToArabicInts map[string]int = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// RomanToInt ...
func RomanToInt(roman string) int {

	var result int
	for ind := 0; ind < len(roman)-1; ind++ {

		multiplier := 1
		if (string(roman[ind]) == "I") && (string(roman[ind+1]) == "V" || string(roman[ind+1]) == "X") {
			multiplier = -1
		} else if (string(roman[ind]) == "X") && (string(roman[ind+1]) == "L" || string(roman[ind+1]) == "C") {
			multiplier = -1
		} else if (string(roman[ind]) == "C") && (string(roman[ind+1]) == "D" || string(roman[ind+1]) == "M") {
			multiplier = -1
		}

		result += multiplier * romanToArabicInts[string(roman[ind])]

	}
	result += romanToArabicInts[string(roman[len(roman)-1])]
	return result
}
