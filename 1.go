// 1. Two Sum

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

func twoSum(nums []int, target int) []int {
    var answer []int
    for i := 0; i < len(nums) - 1; i++ {
        for k := i+1; k < len(nums); k++ {
            if nums[i] + nums[k] == target {
                answer = append(answer, i, k)
                break
            }
        }
    }
    return answer
}