// 20. Valid Parentheses

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

// Input: s = "()[]{}"
// Output: true


func isValid(s string) bool {
    if len(s) % 2 == 1 {
        return false
    }
    pairs := map[rune]rune{
        '(': ')',
        '{': '}',
        '[': ']',
    }
    stack := []rune{}
    for _, val := range s {
        _, ok := pairs[val]
        if ok {
            stack = append(stack, val)
        } else if len(stack) == 0 || pairs[stack[len(stack)-1]] != val {
            return false
        } else {
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}