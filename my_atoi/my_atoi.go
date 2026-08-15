package my_atoi

// import "fmt"

func myAtoi(s string) int {
	var acc int
	isNegative, isFirstDigitFound := false, false
	for _, c := range s {
		if !isFirstDigitFound && c == ' ' {
			// skip leading spaces: they are leading if we have not encountered a digit yet and thus if our accumulator is zero
			continue
		} else if !isFirstDigitFound && c == '-' {
			// Consider '-' as sign only if no digits have already been found.
			// moreover, from now on only digits will be allowed
			isNegative = true
			isFirstDigitFound = true
		} else if !isFirstDigitFound && c == '+' {
			// Consider '+' as sign only if no digits have already been found.
			// moreover, from now on only digits will be allowed
			isFirstDigitFound = true
		} else if c >= '0' && c <= '9' {
			isFirstDigitFound = true
			// If the character is a digit, evaluate it only if it does not make you overflow
			digit := int(c - '0')
			if isNegative {
				digit *= -1
			}
			// fmt.Printf("acc %d | digit %d\r\n", acc, digit)
			//       2147483646
			if acc >= 214748365 || (acc >= 214748364 && digit > 7) {
				acc = 2147483647
				break
			}
			if acc <= -214748365 || (acc <= -214748364 && digit < -8) {
				acc = -2147483648
				isNegative = false // Prevent the flag to flip the sign
				break
			}
			acc = acc*10 + digit
		} else {
			// Stop if you encounter a character
			break
		}
	}
	return acc
}
