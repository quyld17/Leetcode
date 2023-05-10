// 977. Squares of a Sorted Array

// Given an integer array nums sorted in non-decreasing order, 
// return an array of the squares of each number sorted in non-decreasing order.

// Input: nums = [-4,-1,0,3,10]
// Output: [0,1,9,16,100]
// Explanation: After squaring, the array becomes [16,1,0,9,100].
// After sorting, it becomes [0,1,9,16,100].

func sortedSquares(nums []int) []int {
    for i := 0; i < len(nums); i++ {
        nums[i] = nums[i] * nums[i]
    }
    for i := 1; i < len(nums); i++ {
        swapKey := nums[i]
        j := i - 1
        for j >= 0 && nums[j] > swapKey {
            nums[j+1] = nums[j]
            j--
        }
        nums[j+1] = swapKey
    }
    return nums
}