// 35. Search Insert Position

// Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

// You must write an algorithm with O(log n) runtime complexity.

// Input: nums = [1,3,5,6], target = 5
// Output: 2

func searchInsert(nums []int, target int) int {
    index := len(nums)
    for i := 0; i < len(nums); i++ {
        if nums[i] == target {
            index = i
            break
        } else if nums[i] > target {
            index = i 
            break
        }
    }
    return index
}