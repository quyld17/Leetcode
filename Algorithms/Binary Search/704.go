// 704. Binary Search

// Given an array of integers nums which is sorted in ascending order, and an integer target, 
// write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

// You must write an algorithm with O(log n) runtime complexity.

// Input: nums = [-1,0,3,5,9,12], target = 9
// Output: 4
// Explanation: 9 exists in nums and its index is 4


//Recursive Search
func search(nums []int, target int) int {
	p1, p2 := 0, len(nums)-1
	return binary(nums, target, p1, p2)
}
  
func binary(nums []int, target int, p1 int, p2 int) int {
	if p1 > p2 {
	  return -1
	}
	mid := (p1 + p2)/2
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return binary(nums, target, p1, mid-1)
	} else {
		return binary(nums, target, mid+1, p2)
	}
}


//Binary Search 
func search(nums []int, target int) int {
	p1, p2 := 0, len(nums)-1
	for p1 <= p2 {
	  mid := (p1+p2)/2
	  if nums[mid] == target {
		return mid
	  } else if nums[mid] > target {
		p2 = mid - 1
	  } else {
		p1 = mid + 1
	  }
	}
	return - 1
}
