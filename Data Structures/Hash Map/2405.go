// 2405. Optimal Partition of String

// Given a string s, partition the string into one or more substrings such that the characters in each substring are unique. 
// That is, no letter appears in a single substring more than once.

// Return the minimum number of substrings in such a partition.

// Note that each character should belong to exactly one substring in a partition.

// Input: s = "abacaba"
// Output: 4
// Explanation:
// Two possible partitions are ("a","ba","cab","a") and ("ab","a","ca","ba").
// It can be shown that 4 is the minimum number of substrings needed.

func partitionString(s string) int {
    uniqueCheck := map[string]int{}
    p, count := 0, 1
    for p < len(s) {
        _, ok := uniqueCheck[string(s[p])]
        if ok {
            count += 1
            uniqueCheck = make(map[string]int)
            uniqueCheck[string(s[p])] = p
        }
        uniqueCheck[string(s[p])] = p
        p += 1
    }
    return count
}