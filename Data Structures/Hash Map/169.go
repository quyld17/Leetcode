// 169. Majority Element

// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times. 
// You may assume that the majority element always exists in the array.

// Input: nums = [3,2,3]
// Output: 3

func majorityElement(nums []int) int {
    majorityMap := map[int]int{}        
    for _, value := range nums {
        majorityMap[value] += 1
    }
    for index, _ := range majorityMap {
        check := majorityMap[index]
        if check > len(nums)/2 {
            return index
        }
    }
    return 0
}