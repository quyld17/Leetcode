// 567. Permutation in String

// Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

// In other words, return true if one of s1's permutations is the substring of s2.

// Input: s1 = "ab", s2 = "eidbaooo"
// Output: true
// Explanation: s2 contains one permutation of s1 ("ba").


func checkInclusion(s1 string, s2 string) bool {
    need := [26]int{}
    for idx := range s1 { 
        need[s1[idx] - 97] += 1 
    }
    have := [26]int{}
    for idx := range s2 {
        if idx >= len(s1) {
            have[s2[idx-len(s1)] - 97] -= 1
        }
        have[s2[idx] - 97] += 1
        if need == have {
            return true
        }
    }
    return false
}