package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return &ListNode{l2.Val, mergeTwoLists(nil, l2.Next)}
	}
	if l2 == nil {
		return &ListNode{l1.Val, mergeTwoLists(l1.Next, nil)}
	}
	return &ListNode{min(l1.Val, l2.Val), &ListNode{max(l1.Val, l2.Val), mergeTwoLists(l1.Next, l2.Next)}}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
