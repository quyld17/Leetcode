// 344. Reverse String

// Write a function that reverses a string. The input string is given as an array of characters s.

// You must do this by modifying the input array in-place with O(1) extra memory.

// Input: s = ["h","e","l","l","o"]
// Output: ["o","l","l","e","h"]


func reverseString(s []byte)  {
    l, r := 0, len(s)-1
    for l <= r {
        swap := s[l]
        s[l] = s[r]
        s[r] = swap
        l += 1
        r -= 1
    }
}