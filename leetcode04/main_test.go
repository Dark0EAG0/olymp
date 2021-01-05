package main

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{"1", &ListNode{1, &ListNode{1,nil}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteDuplicates(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}
	var a []int
	ls:=rm(head, []int{}, &a)
	return derem(ls, a)
}

func rm(list *ListNode, has []int, a *[]int) *ListNode{
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

func derem(list *ListNode, a []int) *ListNode{
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