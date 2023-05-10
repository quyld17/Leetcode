// 566. Reshape the Matrix

// In MATLAB, there is a handy function called reshape which can reshape an m x n matrix into a new one with a different size r x c keeping its original data.

// You are given an m x n matrix mat and two integers r and c representing the number of rows and the number of columns of the wanted reshaped matrix.

// The reshaped matrix should be filled with all the elements of the original matrix in the same row-traversing order as they were.

// If the reshape operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.

// Input: mat = [[1,2],[3,4]], r = 1, c = 4
// Output: [[1,2,3,4]]


func matrixReshape(mat [][]int, r int, c int) [][]int {
    m, n := len(mat), len(mat[0])   
    if m*n != r*c {
        return mat
    }
    curRow, curCol := 0, 0 
    count := 0
    subReshape := []int{}
    reshape := [][]int{}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if curRow == r {
                break
            }
            subReshape = append(subReshape, mat[i][j])
            curCol += 1
            count += 1
            if curCol == c {
                reshape = append(reshape, subReshape)
                subReshape = []int{}
                curRow += 1
                curCol = 0
            }
        }
    }
    if count == r*c {
        return reshape
    }
    return mat
}