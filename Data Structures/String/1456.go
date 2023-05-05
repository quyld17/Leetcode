// 1456. Maximum Number of Vowels in a Substring of Given Length

// Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.

// Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.

// Input: s = "abciiidef", k = 3
// Output: 3
// Explanation: The substring "iii" contains 3 vowel letters.


func maxVowels(s string, k int) int {
    vowelMap := map[string]int{
        "a": 1,
        "e": 1,
        "i": 1,
        "o": 1,
        "u": 1,
    }
    p1, p2, count, max := 0, k-1, 0, 0
    for i := 0; i <= p2; i++ {
        _, ok := vowelMap[string(s[i])]
        if ok {
            count += 1
        }
        max = count
    }
    p1 += 1
    p2 += 1
    for p2 < len(s) {
        _, ok1 := vowelMap[string(s[p1-1])]
        _, ok2 := vowelMap[string(s[p2])]
        if ok1 {
            count -= 1
        }
        if ok2 {
            count += 1
        }
        if count > max {
            max = count
        }
        p1 += 1
        p2 += 1
    }
    return max
}