// 1539. Kth Missing Positive Number

// Given an array arr of positive integers sorted in a strictly increasing order, and an integer k.

// Return the kth positive integer that is missing from this array.

// Input: arr = [2,3,4,7,11], k = 5
// Output: 9
// Explanation: The missing positive integers are [1,5,6,8,9,10,12,13,...]. The 5th missing positive integer is 9.

func findKthPositive(arr []int, k int) int {
    missingInt, count, pointer := 1, 1, 0
    for pointer < len(arr) {
        if missingInt != arr[pointer] && count == k {
            return missingInt
        }
        if missingInt == arr[pointer] {
            missingInt += 1
            pointer += 1
            continue
        }
        missingInt += 1
        count += 1
    }
    for count < k {
        missingInt += 1
        count += 1
    }
    return missingInt
}