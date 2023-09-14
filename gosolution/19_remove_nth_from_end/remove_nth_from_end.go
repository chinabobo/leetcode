/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{0, head}
    left, right := dummy, dummy
    for i := 0; i < n; i++ {
        if right.Next == nil {
            return nil
        }
        right = right.Next
    }
    for right.Next != nil {
        left = left.Next
        right = right.Next
    }

    left.Next = left.Next.Next

    return dummy.Next
}
