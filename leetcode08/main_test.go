package main

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1 *ListNode
		l2 *ListNode
		want *ListNode
	}{
		{"1", &ListNode{2, &ListNode{4, &ListNode{3,nil}}}, &ListNode{5, &ListNode{6, &ListNode{4,nil}}}, &ListNode{7, &ListNode{0, &ListNode{8,nil}}} },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTwoNumbers(tt.l1, tt.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ListNode : ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers : AddTwoNumbers
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cur1 := l1
	next1 := l1.Next
	cur2 := l2
	next2 := l2.Next
	ans := &ListNode{}
	first := ans
	fi:=true
	for cur1 != nil || cur2 != nil {
		v:=0
		if cur2 != nil {
			v += cur2.Val
		}
		if cur1 != nil {
			v += cur1.Val
		}
		f := 0
		if v >= 10 {
			f = 1
		}
		if fi{
			fi = false
			ans.Val += v % 10
		}else{
			ans.Next = &ListNode{f, nil}
			ans.Val += v % 10
		}
		nilis:=false
		if cur1.Next != nil {
			cur1 = next1
			next1 = cur1.Next
		}else {
			nilis = true
		}
		if cur2.Next != nil{
			cur2 = next2
			next2 = cur2.Next
		}else if nilis{
			break
		}
		ans = ans.Next
	}
	return first
}
