// 41. First Missing Positive

// Given an unsorted integer array nums, return the smallest missing positive integer.

// You must implement an algorithm that runs in O(n) time and uses constant extra space.

// Input: nums = [1,2,0]
// Output: 3
// Explanation: The numbers in the range [1,2] are all in the array.

func firstMissingPositive(nums []int) int {
    existingIntMap := map[int]int{}
    for i := 0; i < len(nums); i++ {
        existingIntMap[nums[i]] = i
    }
    currExistingPositiveInt := 1
    for i := 0; i < len(existingIntMap); i++ {
        _, ok := existingIntMap[currExistingPositiveInt]
        if ok {
            currExistingPositiveInt += 1
            continue
        }
        return currExistingPositiveInt
    }
    return currExistingPositiveInt
}