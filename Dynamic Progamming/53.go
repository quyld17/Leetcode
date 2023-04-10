53. Maximum Subarray

Given an integer array nums, find the subarray with the largest sum, and return its sum.

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.


func maxSubArray(nums []int) int {
    maxSum := nums[0]
    curSum := 0
    for _, val := range nums {
        if curSum < 0 {
            curSum = 0
        }
        curSum += val
        if curSum > maxSum {
            maxSum = curSum
        }
    }
    return maxSum
}