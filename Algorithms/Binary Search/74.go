// 74. Search a 2D Matrix

// You are given an m x n integer matrix matrix with the following two properties:

// Each row is sorted in non-decreasing order.
// The first integer of each row is greater than the last integer of the previous row.
// Given an integer target, return true if target is in matrix or false otherwise.

// You must write a solution in O(log(m * n)) time complexity.

// Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// Output: true


func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)                    //length of 2D array
    n := len(matrix[0])                 //length of subarray
    for i := 0; i < m; i++ {            
        if target > matrix[i][n-1] {    //if target is bigger than 
            continue                    //last index of the subarray 
        }                               //then move to the next subarray
        for j := 0; j < n; j++ {        //check if target is in subarray
            if matrix[i][j] == target { 
                return true             //return true if found
            }
        }
        break                           //if not, break and return false
    }
    return false
}