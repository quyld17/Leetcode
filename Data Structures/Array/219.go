// 219. Contains Duplicate II

// Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array 
// such that nums[i] == nums[j] and abs(i - j) <= k.

// Input: nums = [1,2,3,1], k = 3
// Output: true

func containsNearbyDuplicate(nums []int, k int) bool {
    nearbyDups := make(map[int]int)
    for index, value := range nums {
        lastPos, exist := nearbyDups[value]
        if exist && (index - lastPos) <= k {
            return true
        }
        nearbyDups[value] = index
    }   
    return false
}