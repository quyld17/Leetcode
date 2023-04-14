// 680. Valid Palindrome II

// Given a string s, return true if the s can be palindrome after deleting at most one character from it.

// Input: s = "aba"
// Output: true


func validPalindrome(s string) bool {
    left, right := 0, len(s) - 1

    for left < right {
        if s[left] != s[right] {
            return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
            // the line above means:
            // check1 := isPalindrome(s, left+1, right) 
            // check2 := isPalindrome(s, left, right-1)
            // if check1 == true || check2 == true {
            //     return true
            // } else {
            //     return false
            // }
        }
        left++
        right--
    }

    return true
}

func isPalindrome(s string, left, right int) bool {
    for left < right {
        if s[left] != s[right] {            
            return false
        }
        left++
        right--
    }
    return true
}   
