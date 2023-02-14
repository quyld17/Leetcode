// 643. Maximum Average Subarray I

// You are given an integer array nums consisting of n elements, and an integer k.
// Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. 
// Any answer with a calculation error less than 10-5 will be accepted.

// Input: nums = [1,12,-5,-6,50,3], k = 4
// Output: 12.75000
// Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75

func findMaxAverage(nums []int, k int) float64 {
    sum := 0
    for i := 0; i < k; i++ {
        sum += nums[i]
    }
    max := sum
    pointer := k
    for pointer < len(nums) {
        sum = sum + nums[pointer] - nums[pointer-k]
        pointer += 1
        if sum > max {
            max = sum
        }
    }

    fmt.Println(max)
    return float64(max) / float64(k)
}