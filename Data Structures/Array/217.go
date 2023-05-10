// 217. Contains Duplicate

// Given an integer array nums, return true if any value appears at least twice in the array, 
// and return false if every element is distinct.

// Input: nums = [1,2,3,1]
// Output: true

func containsDuplicate(nums []int) bool {
    valueAppearances := make(map[int]int)
    for _, value := range nums {
        valueAppearances[value] += 1
        count := valueAppearances[value]
        if count > 1 {
            return true
        }
    }
    return false
}