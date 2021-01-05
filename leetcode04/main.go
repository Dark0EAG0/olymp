package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var a []int
	ls := rm(head, []int{}, &a)
	return derem(ls, a)
}

func rm(list *ListNode, has []int, a *[]int) *ListNode {
	for _, el := range has {
		if list.Val == el {
			*a = append(*a, el)
			if list.Next == nil {
				return nil
			}
			return rm(list.Next, has, a)
		}
	}
	has = append(has, list.Val)
	if list.Next == nil {
		return list
	}
	return &ListNode{list.Val, rm(list.Next, has, a)}
}

func derem(list *ListNode, a []int) *ListNode {
	for _, el := range a {
		if list.Val == el {
			if list.Next == nil {
				return nil
			}
			return derem(list.Next, a)
		}
	}
	if list.Next == nil {
		return list
	}
	return &ListNode{list.Val, derem(list.Next, a)}
}
