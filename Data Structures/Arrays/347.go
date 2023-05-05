// 347. Top K Frequent Elements

// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]


func topKFrequent(nums []int, k int) []int {
    uniqueMap := map[int]int{}
    for _, val := range nums {
        uniqueMap[val] += 1
    }
    result := []int{}
    numOfUniqueElements := []int{}
    for _, val := range uniqueMap {
        numOfUniqueElements = append(numOfUniqueElements, val)
    }
    p := len(numOfUniqueElements)-1
    sort.Ints(numOfUniqueElements)
    for k > 0 {
        for key, val := range uniqueMap {
            if val == numOfUniqueElements[p] {
                k -= 1
                uniqueMap[key] = 0
                result = append(result, key)
                p -= 1
                fmt.Println(uniqueMap)
                break
            }
        }
    }
    return result
}