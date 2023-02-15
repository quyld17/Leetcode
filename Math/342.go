// 342. Power of Four

// Given an integer n, return true if it is a power of four. Otherwise, return false.
// An integer n is a power of four, if there exists an integer x such that n == 4x.

// Input: n = 16
// Output: true

func isPowerOfFour(n int) bool {
    if n == 1 {
        return true
    }
    if n == 0 {
        return false
    }
    if n % 4 == 0 {
        n = n/4
        return isPowerOfFour(n)
    }
    return false
}