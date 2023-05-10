// 414. Third Maximum Number

// Given an integer array nums, return the third distinct maximum number in this array. 
// If the third maximum does not exist, return the maximum number.

// Input: nums = [3,2,1]
// Output: 1
// Explanation:
// The first distinct maximum is 3.
// The second distinct maximum is 2.
// The third distinct maximum is 1.


func thirdMax(nums []int) int {
    sort.Ints(nums)
    if len(nums) < 3 {
        return nums[len(nums)-1]
    }
    count := 1
    for i := len(nums)-2; i >= 0; i-- {
        if nums[i] != nums[i+1] {
            count += 1
        } 
        if count == 3 {
            return nums[i]
        }
    }
    return nums[len(nums)-1]
}

 