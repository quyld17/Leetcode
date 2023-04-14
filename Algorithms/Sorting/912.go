// 912. Sort an Array

// Given an array of integers nums, sort the array in ascending order and return it.

// You must solve the problem without using any built-in functions in O(nlog(n)) time complexity 
// and with the smallest space complexity possible.

// Input: nums = [5,2,3,1]
// Output: [1,2,3,5]
// Explanation: After sorting the array, the positions of some numbers are not changed (for example, 2 and 3), 
// while the positions of other numbers are changed (for example, 1 and 5).


//Merge Sort
func sortArray(nums []int) []int {
    return split(nums)
}

func split(nums []int) []int {
    if len(nums) == 1 {
        return nums
    }
    mid := len(nums)/2
    left := nums[:mid]
    right := nums[mid:]

    left = split(left)
    right = split(right)
    merge(left, right)

    return merge(left, right)
}

func merge(left, right []int) []int {
    result := []int{}
    p1, p2 := 0, 0

    for p1 < len(left) && p2 < len(right) {
        if left[p1] < right[p2] {
            result = append(result, left[p1])
            p1 += 1
            continue
        }
        result = append(result, right[p2])
        p2 += 1
    }

    for p1 < len(left) {
        result = append(result, left[p1])
        p1 += 1
    }
    for p2 < len(right) {
        result = append(result, right[p2])
        p2 += 1
    }

    return result
}