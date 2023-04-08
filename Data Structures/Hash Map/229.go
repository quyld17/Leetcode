// 229. Majority Element II

// Given an integer array of size n, 
// find all elements that appear more than ⌊ n/3 ⌋ times.

// Input: nums = [3,2,3]
// Output: [3]


func majorityElement(nums []int) []int {
    element := map[int]int{}
    majorElement := []int{}
    for _, val := range nums {
        element[val] += 1
    }
    for key, value := range element {
        if value > len(nums)/3 {
            majorElement = append(majorElement, key)
        }
    }
    return majorElement
}