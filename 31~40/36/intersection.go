/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	A := headA
	B := headB

	for A != B {
		if A == nil {
			A = headB
		} else {
			A = A.Next
		}
		if B == nil {
			B = headA
		} else {
			B = B.Next
		}
	}
	return A
}