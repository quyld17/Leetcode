540. Single Element in a Sorted Array

You are given a sorted array consisting of only integers where every element appears exactly twice, 
except for one element which appears exactly once.
Return the single element that appears only once.

Your solution must run in O(log n) time and O(1) space.

Input: nums = [1,1,2,3,3,4,4,8,8]
Output: 2

func singleNonDuplicate(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    pointer := 0
    for pointer < len(nums) {
        if nums[pointer] != nums[pointer + 1] {
            return nums[pointer]
        }
        pointer += 2
        if pointer == len(nums) - 1 {
            return nums[pointer]
        }
    }
    return 0
}