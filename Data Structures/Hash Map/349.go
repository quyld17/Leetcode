349. Intersection of Two Arrays

Given two integer arrays nums1 and nums2, return an array of their intersection. 
Each element in the result must be unique and you may return the result in any order.

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]


func intersection(nums1 []int, nums2 []int) []int {
    arr1, arr2, resArr := map[int]int{}, map[int]int{}, map[int]int{}
    result := []int{}
    for _, val := range nums1 {
        arr1[val] += 1
    }
    for _, val := range nums2 {
        arr2[val] += 1
    }
    if len(nums1) >= len(nums2) {
        for _, val := range nums2 {
            _, ok1 := arr1[val]
            _, ok2 := resArr[val]
            if ok1 {
                if ok2 {
                    continue
                }
                resArr[val] = 1
                result = append(result, val)
            }
        }
    } else {
        for _, val := range nums1 {
            _, ok1 := arr2[val]
            _, ok2 := resArr[val]
            if ok1 {
                if ok2 {
                    continue
                }
                resArr[val] = 1
                result = append(result, val)
            }
        }
    }
    return result
}