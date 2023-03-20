// 605. Can Place Flowers

// You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.

// Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, 
// return if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule.

// Input: flowerbed = [1,0,0,0,1], n = 1
// Output: true

func canPlaceFlowers(flowerbed []int, n int) bool {
    for i := 0; i < len(flowerbed); i++ {
        if flowerbed[i] == 1 {
            i += 1
            continue
        }
        if i == len(flowerbed)-1 {
            n -= 1 
            continue
        }
        if flowerbed[i+1] == 1 {
            i += 2
            continue
        }
        n -= 1 
        i += 1
    }
    if n <= 0 {
        return true
    }
    return false
}