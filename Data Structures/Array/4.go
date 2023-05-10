// 4. Median of Two Sorted Arrays

// Given two sorted arrays nums1 and nums2 of size m and n respectively, 
// return the median of the two sorted arrays.

// The overall run time complexity should be O(log (m+n)).

// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var newNums []int
    idx, fp, sp := 0, 0, 0
    for idx < len(nums1) + len(nums2) {
        if fp == len(nums1) || sp == len(nums2) {
            for sp < len(nums2) {
                newNums = append(newNums, nums2[sp])
                sp += 1
            }
            for fp < len(nums1) {
                newNums = append(newNums, nums1[fp])
                fp += 1
            }
            break
        }
        if nums1[fp] <= nums2[sp] {
            newNums = append(newNums, nums1[fp])
            fp += 1
        } else {
            newNums = append(newNums, nums2[sp])
            sp += 1
        }
        idx += 1
    }
    if (len(nums1) + len(nums2)) % 2 == 0 {
        sum := float64(newNums[len(newNums)/2 - 1]) + float64(newNums[len(newNums)/2])
        return sum / 2
    } 
    return float64(newNums[len(newNums)/2])
}