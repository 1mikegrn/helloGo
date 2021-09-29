package itertools

import (
	"testing"
	"container/list"
)

func TestItertools(t *testing.T) {
	items := list.New()

	want := new(ListIter)
	want.Init(items)

	if got := New(items); got.items != want.items {
		t.Errorf("fail")
	}
}
