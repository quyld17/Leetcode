// 263. Ugly Number

// An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.

// Given an integer n, return true if n is an ugly number.

// Input: n = 6
// Output: true
// Explanation: 6 = 2 × 3


func isUgly(n int) bool {
    if n <= 0 {
        return false
    }
    for n % 2 == 0 || n % 3 == 0 || n % 5 == 0 {
        if n % 2 == 0 {n = n / 2}
        if n % 3 == 0 {n = n / 3}
        if n % 5 == 0 {n = n / 5}
    }
    if n != 1 && n != 2 && n != 3 && n != 5 {
        return false
    }
    return true
}