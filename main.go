package main

import (
	"fmt"
)

type Item struct {
	data string
	next *Item
}

func New(items ...string) (head *Item) {
	for i := range items {
		s := items[len(items)-i-1]
		head = NewItem(s, head)
		fmt.Printf("data: %v, addr: %p\n", head, head)
	}
	return head
}

func NewItem(d string, nxt *Item) *Item {
	return &Item{data: d, next: nxt}
}

func (i *Item) Print() string {
	var s string
	for i != nil {
		s += i.data
		if i.next != nil {
			s += ","
		}
		i = i.next
	}
	return s
}

func (i *Item) Last() *Item {
	if i == nil {
		return nil
	}
	for {
		if i.next == nil {
			return i
		}
		i = i.next
	}
}

func (i *Item) Next() *Item {
	if i == nil {
		return nil
	}
	return i.next
}

func (i *Item) AddFirst(n *Item) *Item {
	if i == nil {
		return nil
	}
	n.next = i
	return n
}

func (i *Item) AddLast(n *Item) *Item {
	if i == nil {
		return nil
	}
	n.next = nil
	last := i.Last()
	last.next = n
	return i
}

func (i *Item) Reverse() *Item {
	if i == nil {
		return nil
	}
	var current *Item
	var next *Item
	for i != nil {
		fmt.Printf(">>> head: %p, next: %p, current: %p\n", i, next, current)
		next = i.next
		i.next = current
		current = i
		i = next
	}
	return current
}

func main() {
	h := New("a", "b", "c")
	fmt.Println("New list")
	fmt.Printf("List: %s, head: %p, next: %p, last: %p\n", h.Print(), h, h.Next(), h.Last())

	h = h.AddFirst(NewItem("_", nil))
	fmt.Println("Added first")
	fmt.Printf("List: %s, head: %p, next: %p, last: %p\n", h.Print(), h, h.Next(), h.Last())

	h = h.AddLast(NewItem("zzz", nil))
	fmt.Println("Added last")
	fmt.Printf("List: %s, head: %p, next: %p, last: %p\n", h.Print(), h, h.Next(), h.Last())

	h = h.Reverse()
	fmt.Println("Reversed")
	fmt.Printf("List: %s, head: %p, next: %p, last: %p\n", h.Print(), h, h.Next(), h.Last())
}
