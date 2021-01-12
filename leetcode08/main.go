package main

//// ListNode : ListNode
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//// AddTwoNumbers : AddTwoNumbers
//func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	cur1 := l1
//	next1 := l1.Next
//	cur2 := l2
//	next2 := l2.Next
//	ans := &ListNode{}
//	for next2 != nil && next1 != nil {
//		v:=cur1.Val+cur2.Val
//		if v > 10 {
//			next1.Val++
//		}
//		ans = &ListNode{v%10, ans}
//		cur1 = next1
//		next1 = cur1.Next
//		cur2 = next2
//		next2 = cur2.Next
//	}
//	return ans
//}
