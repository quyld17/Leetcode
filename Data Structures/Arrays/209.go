// 209. Minimum Size Subarray Sum

// Given an array of positive integers nums and a positive integer target, return the minimal length of a subarray 
// whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

// Input: target = 7, nums = [2,3,1,2,4,3]
// Output: 2
// Explanation: The subarray [4,3] has the minimal length under the problem constraint.


func minSubArrayLen(target int, nums []int) int {
    sum, minLen := 0, len(nums) + 1
    firstPointer, secondPointer := 0, 0
    for secondPointer < len(nums) {
        sum += nums[secondPointer]
        for sum >= target {
            if secondPointer - firstPointer < minLen {
                minLen = secondPointer - firstPointer + 1
            }
            sum -= nums[firstPointer]
            firstPointer += 1
        }
        secondPointer += 1
    }
    if minLen == len(nums) + 1 {
        return 0
    }
    return minLen
}