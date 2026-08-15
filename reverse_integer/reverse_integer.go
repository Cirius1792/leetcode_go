package reverse_integer

// Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-2^31, 2^31 - 1], then return 0.
//
// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
//
//
//
// Example 1:
//
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
//

const max = (2147483648 - 1)/10
const min = -2147483648/10

func reverse(x int) int {
	var ret int
	t := x
	for t != 0{
		c := t%10
		if ((ret>max || (ret>max && c > 7)) || (ret<min || (ret<min && c < -8))){
			return 0
		}
		ret = ret*10+c
		t = t/10
	}
	return ret
}
