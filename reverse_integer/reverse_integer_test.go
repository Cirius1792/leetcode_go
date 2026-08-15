package reverse_integer

import "testing"

// Input: x = 123
// Output: 321
//
// Example 2:
//
// Input: x = -123
// Output: -321
//
// Example 3:
//
// Input: x = 120
// Output: 21
func Test_reverseInteger(t *testing.T) {
	testCases := []struct {
		i        int
		expected int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{9999999999, 0},
	}
	for _, tc := range testCases {
		actual := reverse(tc.i)
		if actual != tc.expected{
			t.Errorf("Expected %d, found %d for input %d", tc.expected, actual, tc.i)
		}
	}
	
}
