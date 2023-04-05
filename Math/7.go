// 7. Reverse Integer

// Given a signed 32-bit integer x, return x with its digits reversed. 
// If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

// Input: x = 123
// Output: 321

func reverse(x int) int {
    invertedX := 0
    for x != 0 {
        remainder := x % 10
        invertedX = invertedX * 10 + remainder
        x = x / 10
    }
    if invertedX < -2147483648 || invertedX > 2147483647 {
        return 0
    }
    return invertedX
}