1768. Merge Strings Alternately

You are given two strings word1 and word2. 
Merge the strings by adding letters in alternating order, starting with word1. 
If a string is longer than the other, append the additional letters onto the end of the merged string.

Return the merged string.

Input: word1 = "abc", word2 = "pqr"
Output: "apbqcr"
Explanation: The merged string will be merged as so:
word1:  a   b   c
word2:    p   q   r
merged: a p b q c r


func mergeAlternately(word1 string, word2 string) string {
    var result string
    p1, p2 := 0, 0
    for p1 < len(word1) && p2 < len(word2) {
        if p1 <= p2 {
            result += string(word1[p1])
            p1 += 1
        } else {
            result += string(word2[p2])
            p2 += 1
        }
    }
    for p1 < len(word1) {
        result += string(word1[p1])
        p1 += 1
    }
    for p2 < len(word2) {
        result += string(word2[p2])
        p2 += 1
    }
    return result
}