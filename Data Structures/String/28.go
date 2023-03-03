// 28. Find the Index of the First Occurrence in a String

// Given two strings needle and haystack, 
// return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

// Input: haystack = "sadbutsad", needle = "sad"
// Output: 0
// Explanation: "sad" occurs at index 0 and 6.
// The first occurrence is at index 0, so we return 0.

func strStr(haystack string, needle string) int {
    var haystackArr, needleArr []string
    for _, value := range haystack {
        haystackArr = append(haystackArr, string(value))
    }
    for _, value := range needle {
        needleArr = append(needleArr, string(value))
    }
    fp, sp, count := 0, 0, 0
    for fp < len(haystackArr) {
        if haystackArr[fp] == needleArr[sp] {
            fp += 1
            sp += 1
            count += 1
            if sp == len(needleArr) {
                return fp - len(needleArr)
            }
            continue
        }
        fp = fp - count + 1
        count = 0
        sp = 0
    }
    return -1
}