// 69. Sqrt(x)

// Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.
// You must not use any built-in exponent function or operator.

// For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.

// Input: x = 4
// Output: 2
// Explanation: The square root of 4 is 2, so we return 2.

func mySqrt(x int) int {
    a := 1
    for a <= x {
        b := a * a
        if b < x {
            a++
        } else if b == x {
            return a
        } else if b > x {
            return a - 1
        }
    }
    return 0
}