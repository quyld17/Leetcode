// 557. Reverse Words in a String III

// Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

// Input: s = "Let's take LeetCode contest"
// Output: "s'teL ekat edoCteeL tsetnoc"


func reverseWords(s string) string {
    var result string
    space := []int{-1}
    for i := 0; i < len(s); i++ {
        if string(s[i]) == " " {
            space = append(space, i)
        }
    }
    space = append(space, len(s))
    l, r := 0, 1
    for r < len(space) {
        for i := space[r]-1; i >= space[l]+1; i-- {
            result += string(s[i])
        }
        if r != len(space)-1 {
            result += " "
        }
        l += 1
        r += 1
    }
    return result
}