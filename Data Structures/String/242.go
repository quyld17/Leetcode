// 242. Valid Anagram

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

// Input: s = "anagram", t = "nagaram"
// Output: true

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    var mapLetters = make(map[string]int)
    for _, letter := range s {
        mapLetters[string(letter)] += 1
    }

    for _, letter := range t {
        if mapLetters[string(letter)] == 0 {
            return false
        }
        mapLetters[string(letter)] -= 1
    }
    return true
}
