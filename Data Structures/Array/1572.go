// 1572. Matrix Diagonal Sum

// Given a square matrix mat, return the sum of the matrix diagonals.

// Only include the sum of all the elements on the primary diagonal and all the elements on the secondary diagonal 
// that are not part of the primary diagonal.

// Input: mat = [[1,2,3],
//               [4,5,6],
//               [7,8,9]]
// Output: 25
// Explanation: Diagonals sum: 1 + 5 + 9 + 3 + 7 = 25
// Notice that element mat[1][1] = 5 is counted only once.


func diagonalSum(mat [][]int) int {
    row := len(mat)
    col := len(mat[0])
    sum, p1, p2, p3, mid := 0, 0, col-1, 0, false
    for p3 < row {
        if p1 >= p2 {
            mid = true
            sum += mat[p3][p1]
            p1 -= 1
            p2 += 1
            p3 += 1
            continue
        }
        if p1 > p2 {
            p1 -= 1
            p2 += 1
            sum += mat[p3][p1] + mat[p3][p2]
            mid = true
            p1 -= 1 
            p2 += 1
            p3 += 1 
            continue
        }
        sum += mat[p3][p1] + mat[p3][p2]
        if mid == false {
            p1 += 1
            p2 -= 1
        }
        if mid == true {
            p1 -= 1
            p2 += 1
        }
        p3 += 1
    }
    return sum
}