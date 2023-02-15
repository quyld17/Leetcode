// 125. Valid Palindrome

// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, 
// it reads the same forward and backward. Alphanumeric characters include letters and numbers.
// Given a string s, return true if it is a palindrome, or false otherwise.

// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.

func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    var newPhrase []string
    for _, character := range s {
        if ('a' <= character && character <= 'z') ||
            ('0' <= character && character <= '9') {
            newPhrase = append(newPhrase, string(character))
        }
    }
    charCheckReverse := len(newPhrase)-1
    for charCheck := 0; charCheck < len(newPhrase)/2; charCheck++ {
        if newPhrase[charCheck] != newPhrase[charCheckReverse] {
            return false
        }
        charCheckReverse--
    }
    fmt.Println(newPhrase)
    return true
}