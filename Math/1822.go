// 1822. Sign of the Product of an Array

// There is a function signFunc(x) that returns:
// 	1 if x is positive.
// 	-1 if x is negative.
// 	0 if x is equal to 0.

// You are given an integer array nums. Let product be the product of all values in the array nums.
// Return signFunc(product).

// Input: nums = [-1,-2,-3,-4,3,2,1]
// Output: 1
// Explanation: The product of all values in the array is 144, and signFunc(144) = 1


func arraySign(nums []int) int {
    signFunc := 1
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            return 0
        }
        if nums[i] > 0 {
            continue
        }
        signFunc *= -1
    }
    return signFunc
}