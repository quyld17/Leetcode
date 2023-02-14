// 2139. Minimum Moves to Reach Target Score

// You are playing a game with integers. You start with the integer 1 and you want to reach the integer target.
// In one move, you can either:

// Increment the current integer by one (i.e., x = x + 1).
// Double the current integer (i.e., x = 2 * x).
// You can use the increment operation any number of times, however, you can only use the double operation at most maxDoubles times.

// Given the two integers target and maxDoubles, return the minimum number of moves needed to reach target starting with 1.

// Input: target = 5, maxDoubles = 0
// Output: 4
// Explanation: Keep incrementing by 1 until you reach target.

func minMoves(target int, maxDoubles int) int {
    count := 0
    for target != 1 {
        if maxDoubles != 0 {
            if target % 2 == 0 {
                target = target / 2
                maxDoubles--
            } else {
                target--
            }
        } else {
            target--
        }
        count++
    }
    return count
}