// 54. Spiral Matrix

// Given an m x n matrix, return all elements of the matrix in spiral order.

// Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// Output: [1,2,3,4,8,12,11,10,9,5,6,7]


func spiralOrder(matrix [][]int) []int {
    result := []int{}
    originalMatrixLength := len(matrix) * len(matrix[0])
    return boundariesTravel(matrix, result, originalMatrixLength)
}

func boundariesTravel(matrix [][]int, result []int, originalMatrixLength int) []int{
    row := len(matrix)
    col := len(matrix[0])
    boundaryLength := (row + col - 2) * 2   
    if row == 1 || col == 1 {
        boundaryLength = row * col
    } 
    p1, p2, p3 := 0, 0, 0
    direction := "right" 

    for p3 < boundaryLength {
        switch direction {
            case "right":
                result = append(result, matrix[p1][p2])
                if (p2+1) == col && (p1+1) != row {
                    p1 += 1
                    direction = "down"
                    break
                } 
                p2 += 1
                break

            case "down":
                result = append(result, matrix[p1][p2])
                if (p1+1) == row {
                    p2 -= 1
                    direction = "left"
                    break
                } 
                p1 += 1
                break     
            
            case "left":
                result = append(result, matrix[p1][p2])
                if (p2-1) == -1  {
                    p1 -= 1
                    direction = "up"
                    break
                } 
                p2 -= 1
                break

            case "up":
                result = append(result, matrix[p1][p2])
                if (p1-1) == -1 {
                    p2 += 1
                    direction = "right"
                    break
                } 
                p1 -= 1
                break    
        }
        p3 += 1
    }
    if len(result) == originalMatrixLength {
        return result
    }
    slice1 := matrix[1:(row-1)]
    slice2 := [][]int{}
    for i := 1; i < row-1; i++ {
		slice2 = append(slice2, appendSubMatrix(slice1[i-1], col-1))
	}
    return boundariesTravel(slice2, result, originalMatrixLength)
}

func appendSubMatrix(slice1 []int, i int) []int {
	return slice1[1:i]
} 