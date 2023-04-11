// 118. Pascal's Triangle

// Given an integer numRows, return the first numRows of Pascal's triangle.

// In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

// Input: numRows = 5
// Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]


func generate(numRows int) [][]int {
    result := [][]int{}
    result = append(result, []int{1})
    for i := 2; i <= numRows; i++ {
        arr := []int{1}
        for j := 1; j < i-1; j++ {
            prevRow := result[i-2]
            arr = append(arr, (prevRow[j-1]) + prevRow[j])
        }
        arr = append(arr, 1)
        result = append(result, arr)
    }
    return result
}