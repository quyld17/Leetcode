// 50. Pow(x, n)

// Implement pow(x, n), which calculates x raised to the power n (i.e., xn).

// Input: x = 2.00000, n = 10
// Output: 1024.00000

func myPow(x float64, n int) float64 {
    return math.Pow(x, float64(n))
}