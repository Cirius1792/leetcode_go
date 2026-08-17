// Given an integer x, return true if x is a palindrome, and false otherwise.
//
// Example 1:
//
// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.
// Example 2:
//
// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
// Example 3:
//
// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
package palindrome_number

import "testing"

func Test_isPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		i        int
		expected bool
	}{
		{"121", 121, true},
		{"-121", -121, false},
		{"10", 10, false},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := isPalindrome(tc.i)
			if actual != tc.expected {
				t.Errorf("For input %d expected %t", tc.i, tc.expected)
			}
		})
	}

}
