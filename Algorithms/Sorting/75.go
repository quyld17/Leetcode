// 75. Sort Colors

// Given an array nums with n objects colored red, white, or blue, 
// sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

// We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

// You must solve this problem without using the library's sort function.

// Input: nums = [2,0,2,1,1,0]
// Output: [0,0,1,1,2,2]


func sortColors(nums []int)  {
    //first method: built-in function
    sort.Ints(nums)                                

    //second method: bubble sort
    for i := 0; i < len(nums); i++ {                
        for j := 0; j < len(nums)-1; j++ {
            swap := nums[j]
            if nums[j] > nums[j+1] {
                nums[j] = nums[j+1]
                nums[j+1] = swap
            }
        }
    }
}