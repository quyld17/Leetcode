// 1137. N-th Tribonacci Number

// The Tribonacci sequence Tn is defined as follows: 
// T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.
// Given n, return the value of Tn.

// Input: n = 4
// Output: 4
// Explanation:
// T_3 = 0 + 1 + 1 = 2
// T_4 = 1 + 1 + 2 = 4

func tribonacci(n int) int {
    if n == 0 || n == 1 {
        return n
    }
    if n == 2 {
        return n - 1
    }
    return tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
}