// 167. Two Sum II - Input Array Is Sorted

// Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, 
// find two numbers such that they add up to a specific target number. 
// Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.

// Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.

// The tests are generated such that there is exactly one solution. You may not use the same element twice.

// Your solution must use only constant extra space.

// Input: numbers = [2,7,11,15], target = 9
// Output: [1,2]
// Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].


func twoSum(numbers []int, target int) []int {
    p1, p2 := 0, len(numbers)-1
    for numbers[p1] + numbers[p2] != target {
        if numbers[p1] + numbers[p2] < target {
            p1 += 1
        } else {
            p2 -= 1
        }
    }
    return []int{p1+1, p2+1}
}