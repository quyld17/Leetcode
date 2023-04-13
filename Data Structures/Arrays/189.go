// 189. Rotate Array

// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]


func rotate(nums []int, k int)  {
    k = k % len(nums)
    n := len(nums) - 1
    reverse(nums, 0,n-k)
    reverse(nums, n-k+1, n)
    reverse(nums, 0, n)
}

func reverse(nums []int, l int, r int) {
    for l < r {
        swap := nums[l]
        nums[l] = nums[r]
        nums[r] = swap
        l += 1
        r -= 1
    }   
}