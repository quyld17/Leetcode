// 220. Contains Duplicate III

// You are given an integer array nums and two integers indexDiff and valueDiff.

// Find a pair of indices (i, j) such that:

// i != j,
// abs(i - j) <= indexDiff.
// abs(nums[i] - nums[j]) <= valueDiff, and
// Return true if such pair exists or false otherwise.

// Input: nums = [1,2,3,1], indexDiff = 3, valueDiff = 0
// Output: true
// Explanation: We can choose (i, j) = (0, 3).
// We satisfy the three conditions:
// i != j --> 0 != 3
// abs(i - j) <= indexDiff --> abs(0 - 3) <= 3
// abs(nums[i] - nums[j]) <= valueDiff --> abs(1 - 1) <= 0


func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
    p1, p2 := 0, 1
    for p1 < len(nums)-1 {
        if p2 == len(nums) {
            if p1 < len(nums)-2 {
                p1 += 1
                p2 = p1 + 1
            } else {
                break
            }
        }
        if (p2 - p1) > indexDiff {
            p1 += 1
            p2 = p1 + 1
            continue
        }
        if nums[p2] >= nums[p1] && (nums[p2] - nums[p1]) <= valueDiff {
            return true
        }
        if nums[p2] < nums[p1] && (nums[p1] - nums[p2]) <= valueDiff {
            return true
        }
        p2 += 1
    }
    return false
}