// 189. Rotate Array

// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]


func rotate(nums []int, k int)  {
    if k >= len(nums) {
        k = k % len(nums)
    }
    l, r := len(nums)-k, len(nums)-1
    for l < r {
        swap := nums[l]
        nums[l] = nums[r]
        nums[r] = swap
        l += 1
        r -= 1
    }
    l, r = 0, len(nums)-1-k
    for l < r {
        swap := nums[l]
        nums[l] = nums[r]
        nums[r] = swap
        l += 1
        r -= 1
    }
    l, r = 0, len(nums)-1
    for l < r {
        swap := nums[l]
        nums[l] = nums[r]
        nums[r] = swap
        l += 1
        r -= 1
    }
}