package integer_to_roman

import (
	"strings"
)

func intToRoman(num int) string {
	parametricConverter := func(n int, lowerBound, middleBound, upperBound string) string {
		txt := ""
		if n <= 3 {
			txt = strings.Repeat(lowerBound, n)
		} else if n > 3 && n < 5 {
			txt = strings.Repeat(lowerBound, 5-n) + middleBound
		} else if n >= 5 && n <= 8 {
			txt = middleBound + strings.Repeat(lowerBound, n-5)
		} else {
			txt = strings.Repeat(lowerBound, 10-n) + upperBound
		}
		return txt
	}
	// Handle thousands
	res := num
	thousands := res / 1000
	res %= 1000
	tString := strings.Repeat("M", thousands)

	hundreds := res / 100
	res %= 100
	hString := parametricConverter(hundreds, "C", "D", "M")

	tens := res / 10
	res %= 10
	tensString := parametricConverter(tens, "X", "L", "C")

	unitString := parametricConverter(res, "I", "V", "X")

	return tString + hString + tensString + unitString
}
