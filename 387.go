// 387. First Unique Character in a String

// Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.

// Input: s = "leetcode"
// Output: 0

func firstUniqChar(s string) int {
    charProps := map[string]property{}
    for index, char := range s {
        prop, ok := charProps[string(char)]
        if ok == true{
            prop.count += 1
            charProps[string(char)] = prop
        } else {
            charProps[string(char)] = property{1, index}
        }
    }

    firstNonRepeatingCharIndex := len(s)
    for _, prop := range charProps {
        if prop.count != 1 {
            continue
        }

        if prop.index < firstNonRepeatingCharIndex {
            firstNonRepeatingCharIndex = prop.index
        }
    }

    if firstNonRepeatingCharIndex != len(s) {
        return firstNonRepeatingCharIndex
    }
    return -1

}

type property struct {
    count int
    index int
}
