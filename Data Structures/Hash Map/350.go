// 350. Intersection of Two Arrays II

// Given two integer arrays nums1 and nums2, return an array of their intersection. 
// Each element in the result must appear as many times as it shows in both arrays and 
// you may return the result in any order.

// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2,2]


func intersect(nums1 []int, nums2 []int) []int {
    arr1, arr2 := map[int]int{}, map[int]int{}
    result := []int{}
    for _, val := range nums1 {
        arr1[val] += 1
    }
    for _, val := range nums2 {
        arr2[val] += 1
    }
    if len(nums1) >= len(nums2) {
        for _, val := range nums2 {
            _, ok := arr1[val]
            if ok == true && arr1[val] > 0 {
                result = append(result, val)
                arr1[val] -= 1
            }
        }
    } else {
        for _, val := range nums1 {
            _, ok := arr2[val]
            if ok == true && arr2[val] > 0 {
                result = append(result, val)
                arr2[val] -= 1
            }
        }
    }
    return result
}