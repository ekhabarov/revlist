package main

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	want := "a,b,c"
	got := New("a", "b", "c").Print()

	if got != want {
		t.Fail()
	}
}

func TestNewItem(t *testing.T) {
	want := "aaa"
	got := NewItem("aaa", nil).Print()

	if got != want {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	i := NewItem("a", nil)
	want := fmt.Sprintf("%p", i)
	got := i.String()

	if got != want {
		t.Fail()
	}
}

func TestHead(t *testing.T) {
	want := NewItem("a", nil)
	got := want.Head()
	if got != want {
		t.Fail()
	}
}

func TestNext(t *testing.T) {
	want := NewItem("a", nil)
	got := want.Next()

	if got != want.next {
		t.Fail()
	}
}

func TestLast(t *testing.T) {
	i := NewItem("a", nil)
	got := i.Last()

	var want *Item
	for {
		if i.Next() == nil {
			want = i
			break
		}
		i = i.Next()
	}

	if got != want {
		t.Fail()
	}
}

func TestData(t *testing.T) {
	i := NewItem("a", nil)
	got := i.Data()
	if got != "a" {
		t.Fail()
	}
}

func TestSetNext(t *testing.T) {
	a := NewItem("a", nil)
	b := NewItem("b", nil)
	b.SetNext(a)

	if b.Next() != a {
		t.Fail()
	}
}

func TestAddFirst(t *testing.T) {
	want := "a,b,c"
	got := New("b", "c").AddFirst(NewItem("a", nil)).Print()

	if got != want {
		t.Fail()
	}
}

func TestRemoveFirst(t *testing.T) {
	want := "b,c"
	got := New("a", "b", "c").RemoveFirst().Print()

	if got != want {
		t.Fail()
	}
}

func TestAddLast(t *testing.T) {
	want := "a,b,c"
	got := New("a", "b").AddLast(NewItem("c", nil)).Print()

	if got != want {
		t.Fail()
	}
}

func TestRemoveLast(t *testing.T) {
	want := "a,b"
	got := New("a", "b", "c").RemoveLast().Print()

	if got != want {
		t.Fail()
	}
}

func TestReverse(t *testing.T) {
	want := "d,c,b,a"
	got := New("a", "b", "c", "d").Reverse(false).Print()

	if got != want {
		t.Fail()
	}
}

func TestContainsTrue(t *testing.T) {
	if !New("a", "b", "c", "d").Contains("b") {
		t.Fail()
	}
}

func TestContainsFalse(t *testing.T) {
	if New("a", "b", "c", "d").Contains("z") {
		t.Fail()
	}
}

func TestCount(t *testing.T) {
	want := 4
	got := New("a", "b", "c", "d").Count()

	if got != want {
		t.Fail()
	}
}

func TestMin(t *testing.T) {
	want := "a"
	got := New("a", "b", "c").Min()

	if got != want {
		t.Fail()
	}
}

func TestMax(t *testing.T) {
	want := "c"
	got := New("a", "b", "c").Max()

	if got != want {
		t.Fail()
	}
}

//func TestHead(t *testing.T) {
//if got != want {
//t.Fail()
//}
//}
