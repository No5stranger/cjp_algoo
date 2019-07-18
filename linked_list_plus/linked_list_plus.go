package linked_list_plus

import "fmt"

type Node struct {
	value int
	next  *Node
}

func NewNode(v int) *Node {
	return &Node{value: v}
}

func (node *Node) Print() {
	var n *Node = node
	for n != nil {
		fmt.Println(n.value)
		n = n.next
	}
}

func LinkedListPlus(a, b int) *Node {
	al := Int2LinkedNode(a)
	bl := Int2LinkedNode(b)
	s := revertLinkedList(linkedListPlus(revertLinkedList(al), revertLinkedList(bl)))
	s.Print()
	return s
}

func Int2LinkedNode(n int) *Node {
	// covert int to linkedNode
	// e.g: 1123   1 -> 1 -> 2 -> 3
	var l []int
	var ll int
	for n > 0 {
		l = append(l, n%10)
		n = n / 10
		ll++
	}
	var head *Node = nil
	var pre *Node = nil
	for i := 0; i < ll; i++ {
		n := NewNode(l[i])
		if head == nil {
			head = n
		}
		if pre != nil {
			pre.next = n
		}
		pre = n
	}
	return revertLinkedList(head)
}

func NodePlus() {
	one := NewNode(1)
	two := NewNode(2)
	four := NewNode(4)
	three := NewNode(3)
	five := NewNode(5)
	six := NewNode(6)
	seven := NewNode(7)
	one.next = two
	two.next = four
	four.next = three
	five.next = six
	six.next = seven
	rOne := revertLinkedList(one)
	rFive := revertLinkedList(five)
	result := linkedListPlus(rOne, rFive)
	rResult := revertLinkedList(result)
	rResult.Print()
}

func linkedListPlus(a *Node, b *Node) *Node {
	var c *Node
	var head *Node = nil
	var av, bv, sv int
	var plus bool = false
	var pre *Node
	for a != nil || b != nil {
		if a != nil {
			av = a.value
			a = a.next
		} else {
			av = 0
		}
		if b != nil {
			bv = b.value
			b = b.next
		} else {
			bv = 0
		}
		sv = av + bv
		if plus == true {
			sv++
		}
		if sv >= 10 {
			sv = sv - 10
			plus = true
		} else {
			plus = false
		}
		c = NewNode(sv)
		if pre != nil {
			pre.next = c
		}
		if head == nil {
			head = c
		}
		pre = c
	}
	return head
}

func revertLinkedList(head *Node) *Node {
	var pre *Node
	var next *Node
	for head != nil {
		next = head.next
		head.next = pre
		pre = head
		head = next
	}
	return pre
}
