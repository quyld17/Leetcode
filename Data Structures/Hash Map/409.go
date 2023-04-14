// 409. Longest Palindrome

// Given a string s which consists of lowercase or uppercase letters, 
// return the length of the longest palindrome that can be built with those letters.

// Letters are case sensitive, for example, "Aa" is not considered a palindrome here.

// Input: s = "abccccdd"
// Output: 7
// Explanation: One longest palindrome that can be built is "dccaccd", whose length is 7.


func longestPalindrome(s string) int {
    letter := map[string]int{}
    for _, val := range s {
        letter[string(val)] += 1
    }
    longestPalindrome, oddExist := 0, 0
    for _, val := range letter {
        if val % 2 == 0 {
            longestPalindrome += val
            continue
        } 
        longestPalindrome += val - 1
        oddExist = 1
    }
    if oddExist == 1 {
        return longestPalindrome + 1
    }
    return longestPalindrome 
}