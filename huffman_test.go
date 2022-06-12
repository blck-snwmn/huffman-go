package huffmango

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	ih := &tree{}

	heap.Init(ih)
	heap.Push(ih, node{value: 4})
	heap.Push(ih, node{value: 1})
	heap.Push(ih, node{value: 5})
	heap.Push(ih, node{value: 10})

	if got := heap.Pop(ih); !reflect.DeepEqual(got, node{value: 1}) {
		t.Errorf("pop=%v, want=%v", got, node{value: 1})
	}

	if got := heap.Pop(ih); !reflect.DeepEqual(got, node{value: 4}) {
		t.Errorf("pop=%v, want=%v", got, node{value: 4})
	}
	if got := heap.Pop(ih); !reflect.DeepEqual(got, node{value: 5}) {
		t.Errorf("pop=%v, want=%v", got, node{value: 5})
	}
	if got := heap.Pop(ih); !reflect.DeepEqual(got, node{value: 10}) {
		t.Errorf("pop=%v, want=%v", got, node{value: 10})
	}
}
