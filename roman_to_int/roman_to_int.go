package roman_to_int

func romanToInt(s string) int {
	conversionMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	runes := []rune(s)
	// In a roman number, more than one character is used for a single int digit
	// therefore we need to understand when a single digit is "complete"
	// Example:
	// 	- III -> 3
	// 	- IV  -> 4
	// 	- VI  -> 6
	//  - MCMXCIV -> 1994
	ret := 0
	prev := conversionMap[runes[len(runes)-1]]
	current := 0
	for i := len(runes) - 1; i >= 0; i-- {
		current = conversionMap[runes[i]]
		if i == len(runes)-1 {
			ret += current
			prev = current
		} else if prev > current {
			ret -= current
		} else {
			ret += current
			prev = current
		}
	}

	return ret
}
