// 66. Plus One

// You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. 
// The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

// Increment the large integer by one and return the resulting array of digits.

// Input: digits = [1,2,3]
// Output: [1,2,4]
// Explanation: The array represents the integer 123.
// Incrementing by one gives 123 + 1 = 124.
// Thus, the result should be [1,2,4].

func plusOne(digits []int) []int {
    var new []int
    a := len(digits)
    count := 0
    if digits[a-1] != 9 {
        for i := 0; i < a-1; i++ {
            new = append(new, digits[i])
        }
        new = append(new, digits[a-1]+1)
    } else {
        for j := a-1; j > -1; j-- {
            if digits[j] == 9 {
                count++
            } else {
                break
            }
        }
        if count == a {
            for k := 0; k < (a - count); k++ {
                new = append(new, digits[k])
            }
            new = append(new, 1)
            for count != 0 {
                new = append(new, 0)
                count--
            }
        } else {
            for k := 0; k < (a - count - 1); k++ {
                new = append(new, digits[k])
            }
            new = append(new, digits[a - count - 1] + 1)
            for count != 0 {
                new = append(new, 0)
                count--
            }
        }
    }
    return new
}