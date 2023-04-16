// 438. Find All Anagrams in a String

// Given two strings s and p, return an array of all the start indices of p's anagrams in s. 
// You may return the answer in any order.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, 
// typically using all the original letters exactly once.

// Input: s = "cbaebabacd", p = "abc"
// Output: [0,6]
// Explanation:
// The substring with start index = 0 is "cba", which is an anagram of "abc".
// The substring with start index = 6 is "bac", which is an anagram of "abc".


func findAnagrams(s string, p string) []int {
    need := [26]int{}
    result := []int{}
    for idx := range p { 
        need[p[idx] - 97] += 1 
    }
    have := [26]int{}
    for idx := range s {
        if idx >= len(p) {
            have[s[idx-len(p)] - 97] -= 1
        }
        have[s[idx] - 97] += 1
        if need == have {
            result = append(result, idx - len(p) + 1)
        }
    }
    return result
}