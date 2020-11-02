package _2_linkedlist

import "testing"

func TestInsertTail(t *testing.T) {
	count := 10
	ll := NewLinkedList()
	for i := 0; i < count; i++ {
		ll.InsertToTail(i)
	}
	ll.Print()
}

func TestInsertHead(t *testing.T) {
	count := 10
	ll := NewLinkedList()
	for i := 0; i < count; i++ {
		ll.InsertToHead(i)
	}
	ll.Print()
}

func TestFindByIndex(t *testing.T) {
	ll := NewLinkedList()
	for i := 0; i < 10; i++ {
		ll.InsertToTail(i)
	}
	ln := ll.FindByIndex(uint(3))
	ll.Print()
	t.Log(ln.GetValue())
}

func TestDelete(t *testing.T) {
	ll := NewLinkedList()
	for i := 0; i < 10; i++ {
		ll.InsertToTail(i)
	}
	ln := ll.FindByIndex(uint(3))
	ll.Print()
	ll.Delete(ln)
	ll.Print()
}

func TestReverse(t *testing.T) {
	ll := NewLinkedList()
	for i := 0; i < 10; i++ {
		ll.InsertToTail(i)
	}
	ll.Print()
	ll.Reverse()
	ll.Print()
}

func TestHasCycle(t *testing.T) {
	ll := NewLinkedList()
	for i := 0; i < 15; i++ {
		ll.InsertToTail(i)
	}
	ll.Print()
	ln := ll.FindByIndex(uint(13))
	t.Log(ll.HasCycle())
	ln.next = ll.head.next
	t.Log(ll.HasCycle())
}

func TestFindMiddleNode(t *testing.T) {
	ll := NewLinkedList()
	for i := 0; i < 10; i++ {
		ll.InsertToTail(i)
	}
	ll.Print()
	mid := ll.FindMiddleNode()
	t.Log(mid.GetValue())
}

func TestMerge(t *testing.T) {
	ll := NewLinkedList()
	for i := 0; i < 10; i++ {
		ll.InsertToTail(i)
	}
	ll.Print()
	ll1 := NewLinkedList()
	for i := 5; i < 15; i++ {
		ll1.InsertToTail(i)
	}
	ll1.Print()
	res := MergeSortedList(ll, ll1)
	res.Print()
}

func TestDeleteBottomN(t *testing.T) {
	ll := NewLinkedList()
	for i := 0; i < 10; i++ {
		ll.InsertToTail(i)
	}
	ll.Print()
	ll.DeleteBottomN(4)
	ll.Print()
}
