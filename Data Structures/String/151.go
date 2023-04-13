// 151. Reverse Words in a String

// Given an input string s, reverse the order of the words.

// A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

// Return a string of the words in reverse order concatenated by a single space.

// Note that s may contain leading or trailing spaces or multiple spaces between two words. 
// The returned string should only have a single space separating the words. Do not include any extra spaces.

// Input: s = "  hello world  "
// Output: "world hello"
// Explanation: Your reversed string should not contain leading or trailing spaces.


func reverseWords(s string) string {
    var result, word string
    wordsArr := []string{}
    for _, val := range s {
        if string(val) == " " {
            if word != "" {
                wordsArr = append(wordsArr, word)
            }
            word = ""
            continue
        }
        word += string(val)
    }
    if word != "" {
        wordsArr = append(wordsArr, word)
    }
    for i := len(wordsArr)-1; i >= 0; i-- {
        result += wordsArr[i]
        if i != 0 {
            result += " "
        }
    }
    return result
}