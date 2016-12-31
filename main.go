package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type (
	List struct {
		head  *Item
		count int
	}

	Item struct {
		data string
		next *Item
	}

	Printer interface {
		Print() string
		Head() *Item
		Next() *Item
		Last() *Item
		Count() int
	}

	Content [][]string
)

func (c *Content) Add(method string, p Printer) {
	*c = append(*c, []string{
		method,
		p.Print(),
		p.Head().String(),
		p.Next().String(),
		p.Last().String(),
		fmt.Sprintf("%d", p.Count()),
	})
}

//Creates new list
func New(items ...string) *List {
	l := new(List)
	var n *Item
	for i := range items {
		n = NewItem(items[len(items)-i-1])
		n.next = l.head
		l.head = n
	}
	l.count = len(items)
	return l
}

//Creates new item
func NewItem(d string) *Item {
	return &Item{data: d, next: nil}
}

func (i *Item) String() string {
	return fmt.Sprintf("%p", i)
}

//Prints list as string
func (l *List) Print() string {
	var s string
	p := l.head
	for p != nil {
		s += p.data
		if p.next != nil {
			s += ","
		}
		p = p.next
	}
	return s
}

//Prints list as string with addresses
func (l *List) Print2() (r string) {
	if l == nil {
		return
	}

	var s []*Item

	h := l.head
	for h != nil {
		s = append(s, h)
		h = h.next
	}

	for _, a := range s {
		r += fmt.Sprintf("\n%p (v: %q)", a, a.Data())
	}
	return
}

//Return itself as head item
func (l *List) Head() *Item {
	if l == nil {
		return nil
	}
	return l.head
}

//Returns Next item
func (l *List) Next() *Item {
	if l == nil || l.head == nil {
		return nil
	}
	return l.head.next
}

//Returns Last item
func (l *List) Last() *Item {
	if l == nil || l.head == nil {
		return nil
	}
	h := l.head
	for {
		if h.next == nil {
			return h
		}
		h = h.next
	}
}

//Returns data
func (i *Item) Data() string {
	if i == nil {
		return ""
	}
	return i.data
}

//Returns number of items
func (l *List) Count() int {
	return l.count
}

//Adds Next item
func (i *Item) SetNext(n *Item) *Item {
	if i == nil {
		return nil
	}
	i.next = n
	return i
}

//Adds new item to start
func (l *List) AddFirst(n *Item) {
	if l == nil {
		return
	}
	n.next = l.head
	l.head = n
	l.count++
}

//Removes first item
func (l *List) RemoveFirst() {
	if l == nil || l.head == nil {
		return
	}
	l.head = l.Next()
	l.count--
}

//Addes new item to the end
func (l *List) AddLast(n *Item) {
	if l == nil {
		return
	}
	if l.Count() < 1 {
		l.head = n
	} else {
		l.Last().next = n
	}
	l.count++
}

//Removes item from the end
func (l *List) RemoveLast() {
	if l == nil || l.head == nil {
		return
	}
	last := l.Last()
	n := l.Head()
	for {
		if n.next == last {
			n.next = nil
			l.count--
			return
		}
		n = n.next
	}
}

//Returns reversed head item of reversed list
func (l *List) Reverse() {
	if l == nil {
		return
	}

	var current *Item
	var next *Item

	for l.head != nil {
		next = l.head.next
		l.head.next = current
		current = l.head
		l.head = next
	}
	l.head = current
}

//Returns true if list contains item with provided string
func (l *List) Contains(s string) bool {
	h := l.head
	for {
		if h == nil {
			return false
		}
		if h.data == s {
			return true
		}
		h = h.next
	}
}

//Returns min item
func (l *List) Min() (m string) {
	h := l.head
	m = h.Data()
	for {
		if h == nil {
			return ""
		}
		if strings.Compare(m, h.Data()) == 1 {
			m = h.Data()
		}
		if h.next == nil {
			return m
		}
		h = h.next
	}
}

//Returns max item
func (l *List) Max() (m string) {
	h := l.head
	m = h.Data()
	for {
		if h == nil {
			return ""
		}
		if strings.Compare(m, h.Data()) == -1 {
			m = h.Data()
		}
		if h.next == nil {
			return m
		}
		h = h.next
	}
}

func (l *List) Find(s string) *List {
	h := l.head
	fl := new(List)
	for {
		if l == nil || l.head == nil {
			return nil
		}
		if strings.Compare(h.Data(), s) == 0 {
			fl.AddLast(NewItem(h.Data()))
		}
		if h.next == nil {
			return fl
		}
		h = h.next
	}
}

func main() {
	var ctn Content

	list := New("d", "a", "f", "g", "h", "b", "c", "d", "e", "d")
	ctn.Add("New list", list)

	fmt.Println(list.Print2())

	list.AddFirst(NewItem("!"))
	ctn.Add("Added first", list)

	list.AddLast(NewItem("z"))
	list.AddLast(NewItem("z"))
	list.AddLast(NewItem("z"))
	ctn.Add("Added last", list)

	list.Reverse()
	ctn.Add("Reverse", list)

	list.RemoveFirst()
	ctn.Add("Removed first", list)

	list.RemoveLast()
	ctn.Add("Removed last", list)

	list.RemoveLast()
	ctn.Add("Removed last", list)

	list.RemoveFirst()
	ctn.Add("Removed first", list)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Method", "List", "Head", "Next", "Last", "Count"})
	table.SetRowLine(false)
	for _, v := range ctn {
		table.Append(v)
	}
	table.Render()

	fmt.Printf("Contains %s: %v\n", "c", list.Contains("c"))
	fmt.Printf("Contains %s: %v\n", "v", list.Contains("v"))
	fmt.Printf("Min: %q\n", list.Min())
	fmt.Printf("Max: %q\n", list.Max())
	fmt.Printf("Find: %s\n", list.Find("d").Print2())

}
