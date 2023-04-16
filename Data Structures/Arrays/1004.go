// 1004. Max Consecutive Ones III

// Given a binary array nums and an integer k, 
// return the maximum number of consecutive 1's in the array if you can flip at most k 0's.

// Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
// Output: 6
// Explanation: [1,1,1,0,0,1,1,1,1,1,1]
// Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.


func longestOnes(nums []int, k int) int {
    p1, p2, numOfZeros, max := 0, 0, 0, 0
    for p2 < len(nums) {
        if nums[p2] == 1 {
            p2 += 1
            if p2 - p1 > max {
                max = p2 - p1
            }
            continue
        }
        if numOfZeros == k {
            if nums[p1] == 0 {
                p1 += 1
            } else {
                for nums[p1] != 0 {
                    p1 += 1
                }
                p1 += 1
            }
            p2 += 1
            continue
        }
        p2 += 1
        numOfZeros += 1
        if p2 - p1 > max {
            max = p2 - p1
        }
    }
    return max
}