// 367. Valid Perfect Square

// Given a positive integer num, return true if num is a perfect square or false otherwise.
// A perfect square is an integer that is the square of an integer. In other words, it is the product of some integer with itself.

// You must not use any built-in library function, such as sqrt.

// Input: num = 16
// Output: true
// Explanation: We return true because 4 * 4 = 16 and 4 is an integer.

func isPerfectSquare(num int) bool {
    a, b := true, false
    x := 1
    for x <= num {
        z := x * x
        if z == num {
            return a
        } else if z > num {
            break
        }
        x++
    }
    return b
}