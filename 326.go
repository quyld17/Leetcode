// 326. Power of Three

// Given an integer n, return true if it is a power of three. Otherwise, return false.
// An integer n is a power of three, if there exists an integer x such that n == 3x.

// Input: n = 27
// Output: true
// Explanation: 27 = 33

func isPowerOfThree(n int) bool {
    if n == 0 {
        return false
    }
    if n == 1 {
        return true
    }
    if n % 3 == 0 {
        return isPowerOfThree(n/3)
    }
    return false
}