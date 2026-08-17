package palindrome_number

import (
	"strconv"
)

func isPalindrome(x int) bool {
	// if the number is negative, it can't be palindrome
	if x < 0 {
		return false
	}
	// Convert the int to string
	txt := strconv.Itoa(x)

	// iterate over the string with two poinets h(ead) and t(ail)
	// compare the values of s[t] and s[h] until h <= t
	// return false at first mismatch, otherwise return true
	h, t := 0, len(txt)-1
	for h < t {
		if txt[h] != txt[t] {
			return false
		}
		h++
		t--
	}

	return true
}
