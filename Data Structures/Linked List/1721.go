1721. Swapping Nodes in a Linked List

You are given the head of a linked list, and an integer k.

Return the head of the linked list after swapping the values of the kth node from the beginning and the kth node from the end (the list is 1-indexed).

Input: head = [7,9,6,6,7,8,3,0,9,5], k = 5
Output: [7,9,6,6,8,7,3,0,9,5]


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//My solution
func swapNodes(head *ListNode, k int) *ListNode {
    temp1, temp2, idx := head, head, 1
    linkedListLength := 1
    for temp1.Next != nil {
        linkedListLength += 1
        temp1 = temp1.Next
    }
    temp1 = head
    foundFirst := 0
    for temp1 != nil {
        fmt.Println(temp1.Val)
        if foundFirst == 0 {
            if idx == k || (linkedListLength - idx + 1) == k {
                foundFirst = 1
                for i := idx; i > 1; i-- {
                    temp2 = temp2.Next
                }
            }
        } else if idx == k || (linkedListLength - idx + 1) == k {
            temp3 := temp2.Val
            temp2.Val = temp1.Val
            temp1.Val = temp3
        }
        idx += 1
        temp1 = temp1.Next
    }
    return head
}


//Better solution
func swapNodes(head *ListNode, k int) *ListNode {
    temp1, temp2, p := head, head, head
    for p != nil {
        k -= 1
        if k == 0 {
            temp1 = p
        }
        if k < 0 {
            temp2 = temp2.Next
        }
        p = p.Next
    }
    temp1.Val, temp2.Val = temp2.Val, temp1.Val
    return head
}