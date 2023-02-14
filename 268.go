// 268. Missing Number

// Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

// Input: nums = [3,0,1]
// Output: 2
// Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.

func missingNumber(nums []int) int {
    max := 0
    for _, v := range nums {
        if v > max {
            max = v 
        }
    }
    var compare []int
    for i := 0; i <= max; i++ {
        compare = append(compare, i)
    }
    count := 0
    for k := 0; k < len(compare); k++ {
        for u := 0; u < len(nums); u++ {
            if compare[k] == nums[u] {
                count++
            }
        }
        if count == 0 {
            return compare[k]
        }
		count = 0
    }
    return max + 1
}