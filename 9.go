// 9. Palindrome Number

// Given an integer x, return true if x is a palindrome, and false otherwise.

// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.

func isPalindrome(x int) bool {
    a := true
    b := false
    original := x
    var remainder int
    reverse := 0
    if x < 0 {
        return b
    }
    for x != 0 {
        remainder = x % 10
        reverse = reverse * 10 + remainder
        x = x / 10
    }
    if original == reverse {
        return a
    } else {
        return b
    }
}