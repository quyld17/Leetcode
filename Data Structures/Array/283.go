// 283. Move Zeroes

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
// Note that you must do this in-place without making a copy of the array.

// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]

func moveZeroes(nums []int)  {
    firstPointer, secondPointer := 0, 1
    if len(nums) > 1 {
        for firstPointer < len(nums) - 1 {
            if secondPointer == len(nums) {
                break
            }
            if nums[firstPointer] != 0 {
                firstPointer += 1 
                secondPointer = firstPointer + 1
                continue
            }
            if nums[firstPointer] == 0 {
                if nums[secondPointer] != 0 {
                    swap := nums[firstPointer]
                    nums[firstPointer] = nums[secondPointer]
                    nums[secondPointer] = swap
                }
                secondPointer += 1
            }
        }
    }
}