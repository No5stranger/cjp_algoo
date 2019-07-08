package linkedlist

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	testLinkedList := NewLinkedList()
	testLinkedList.InsertToHead(0)
	for i := 1; i <= 7; i++ {
		testLinkedList.InsertToTail(uint(i))
	}
	testLinkedList.PrintLinkedList()
	mid, midNode := testLinkedList.GetMidNode()
	fmt.Println(fmt.Sprintf("mid:%d", mid))
	PrintNode(midNode)
	testLinkedList.Revert()
	testLinkedList.PrintLinkedList()
}

func TestIsPalindrome(t *testing.T) {
	testLinkedList := NewLinkedList()
	testLinkedList.InsertToHead(1)
	for i := 2; i <= 5; i++ {
		testLinkedList.InsertToTail(i)
	}
	for i := 5; i >= 1; i-- {
		testLinkedList.InsertToTail(i)
	}
	testLinkedList.PrintLinkedList()
	isPalindrome := testLinkedList.IsPalindrome()
	testLinkedList.PrintLinkedList()
	if !isPalindrome {
		t.Errorf("test isPalindrome failed")
	}
}
