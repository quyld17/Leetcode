// 234. Palindrome Linked List

// Given the head of a singly linked list, return true if it is a palindrome or false otherwise.

// Input: head = [1,2,2,1]
// Output: true


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func isPalindrome(head *ListNode) bool {
    arrPalindrome := []int{}
    for head != nil {
        arrPalindrome = append(arrPalindrome, head.Val)
        head = head.Next
    }
    p1, p2 := 0, len(arrPalindrome)-1
    for p1 < p2 {
        if arrPalindrome[p1] != arrPalindrome[p2] {
            return false
        }
        p1 += 1
        p2 -= 1
    }
    return true
}