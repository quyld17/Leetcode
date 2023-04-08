// 904. Fruit Into Baskets

// You are visiting a farm that has a single row of fruit trees arranged from left to right. 
// The trees are represented by an integer array fruits where fruits[i] is the type of fruit the ith tree produces.

// You want to collect as much fruit as possible. However, the owner has some strict rules that you must follow:

// You only have two baskets, and each basket can only hold a single type of fruit. There is no limit on the amount of fruit each basket can hold.
// Starting from any tree of your choice, you must pick exactly one fruit from every tree (including the start tree) while moving to the right. 
// The picked fruits must fit in one of your baskets.
// Once you reach a tree with fruit that cannot fit in your baskets, you must stop.
// Given the integer array fruits, return the maximum number of fruits you can pick.

// Input: fruits = [1,2,1]
// Output: 3
// Explanation: We can pick from all 3 trees.


func totalFruit(fruits []int) int {
    treeType := map[int]int{}
    currTreeNum, max := 0, 0
    p1, p2, skip, check := 0, 0, 0 ,0
    for p1 < len(fruits) {
        if currTreeNum > max {              //update max
            max = currTreeNum
        }
        if p2 == len(fruits) {              //thoát loop nếu p2 bằng len
            break
        }
        if fruits[p2] == fruits[p1] {       //lưu lại số phần tử bằng fruits[p1] liên tiếp
            if check == 0 {                 //để khi pair mới thì skip, tiết kiệm thời gian
                skip += 1
            }
        } else {
            check = 1
        }
        _, ok := treeType[fruits[p2]]       //tạo map để lưu các loại quả đang trong giỏ
        if ok {                             //check nếu trùng mới loại có trong map rồi thì 
            currTreeNum += 1                //tăng currNum lên và continue
            p2 += 1
            continue
        }
        if len(treeType) < 2 {              //check xem đủ 2 loại chưa, nếu chưa thì lưu thêm
            treeType[fruits[p2]] = 1
            currTreeNum += 1 
            p2 += 1
            continue
        }
        treeType = map[int]int{}            //phát hiện loại thứ 3 nên tạo lại map mới 
        p1 += skip                          //tăng value cho các con trỏ
        p2 = p1                             
        skip, check = 0, 0                  //currNum, skip, check quay về 0
        currTreeNum = 0
    }
    return max
}