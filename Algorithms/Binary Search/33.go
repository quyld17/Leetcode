33. Search in Rotated Sorted Array

There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) 
such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). 
For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, 
return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4


//Solution 1: 
func search(nums []int, target int) int {
    p1, p2 := 0, len(nums)-1
    for p1 <= p2 {
        if nums[p1] == target {
            return p1 
        }
        if nums[p2] == target {
            return p2
        }
        p1 += 1
        p2 -= 1
    }
    return -1
}


//Solution 2:
func search(nums []int, target int) int {
    if nums[0] == target {
        return 0
    }
    if nums[0] > target {
        for i := len(nums)-1; i > 0; i-- {
            if nums[i] == target {
                return i
            }
            if nums[i-1] > nums[i] {
                break
            }
        }
    } else {
        for i := 0; i < len(nums); i++ {
            if nums[i] == target {
                return i
            }

            if i < len(nums)-1 && nums[i+1] < nums[i] {
                break
            }
        }
    }
    return -1
}