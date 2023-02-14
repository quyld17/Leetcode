// 58. Length of Last Word

// Given a string s consisting of words and spaces, return the length of the last word in the string.

// A word is a maximal substring consisting of non-space characters only.

// Input: s = "Hello World"
// Output: 5
// Explanation: The last word is "World" with length 5.

func lengthOfLastWord(s string) int {
    count := 0
    recount := 0
    for _, v := range s {
        if string(v) != " " {
            count++
        } else {
            count = 0
        } 
        if count != 0 {
            recount = count
        }
    }
    if count == 0 {
        return recount
    } else {
        return count
    }
}