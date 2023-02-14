// 136. Single Number

// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

// You must implement a solution with a linear runtime complexity and use only constant extra space.

// Input: nums = [2,2,1]
// Output: 1

func singleNumber(nums []int) int {
    xor_result := 0
    for _, number := range nums{
        xor_result ^= number
    }
    return xor_result
}