// 1446. Consecutive Characters

// The power of the string is the maximum length of a non-empty substring that contains only one unique character.
// Given a string s, return the power of s.

// Input: s = "leetcode"
// Output: 2
// Explanation: The substring "ee" is of length 2 with the character 'e' only.

func maxPower(s string) int {
    var count, max int
    var char string
    for _, value := range s {
        if string(value) == char {
            count += 1
        } else {
            char = string(value)
            count = 1
        }
        if count > max {
            max = count
        }
    }
    return max
}