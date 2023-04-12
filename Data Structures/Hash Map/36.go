36. Valid Sudoku

Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

Note:
A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.

Input: board = 
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true


func isValidSudoku(board [][]byte) bool {
    check1 := map[byte]int{}
    check2 := map[byte]int{}
    for i := 0; i < len(board); i++ {           //check rows and columns
        for j := 0; j < len(board[i]); j++ {  
            if board[i][j] != '.' {             //check rows
                _, ok := check1[board[i][j]]
                if ok {
                    return false
                }
                check1[board[i][j]] = 1
            }
            if board[j][i] != '.' {
                _, ok := check2[board[j][i]]    //check columns
                if ok {
                    return false
                }
                check2[board[j][i]] = 1
            }
        }
        check1 = map[byte]int{}
        check2 = map[byte]int{}
    }
    p1, box := 0, 3
    for p1 < len(board) {                       //check sub-boxes
        if p1 == 0 || p1 == 3 || p1 == 6 {      //reset hashmap when move to a new box
            check1 = map[byte]int{}
        }
        p2 := box - 3
        for p2 < box {                          //check box
            if board[p1][p2] == '.' {
                p2 += 1
                continue
            }
            _, ok := check1[board[p1][p2]]
            if ok {
                return false
            }
            check1[board[p1][p2]] = 1
            p2 += 1
        }
        p1 += 1
        if p1 == 9 && box != 9 {                //first pointer return to start if
            p1 = 0                              //it is not final box
            box += 3
        }
    }
    return true
}