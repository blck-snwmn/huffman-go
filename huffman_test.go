package huffmango

import (
	"container/heap"
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestPriorityQueue(t *testing.T) {
	{
		tmp := make(queue, 0, 10)
		table := &tmp

		heap.Init(table)

		heap.Push(table, &leaf{'a', 2})
		heap.Push(table, &leaf{'b', 2})
		heap.Push(table, &leaf{'c', 4})
		heap.Push(table, &leaf{'d', 4})
		heap.Push(table, &leaf{'f', 4})
		heap.Push(table, &leaf{'e', 8})

		if p := heap.Pop(table); !reflect.DeepEqual(p, &leaf{'b', 2}) {
			t.Errorf("got=%+v,want=%+v", p, &leaf{'b', 2})
		}
		if p := heap.Pop(table); !reflect.DeepEqual(p, &leaf{'a', 2}) {
			t.Errorf("got=%+v,want=%+v", p, &leaf{'a', 2})
		}
		if p := heap.Pop(table); !reflect.DeepEqual(p, &leaf{'f', 4}) {
			t.Errorf("got=%+v,want=%+v", p, &leaf{'f', 4})
		}
		if p := heap.Pop(table); !reflect.DeepEqual(p, &leaf{'d', 4}) {
			t.Errorf("got=%+v,want=%+v", p, &leaf{'d', 4})
		}
		if p := heap.Pop(table); !reflect.DeepEqual(p, &leaf{'c', 4}) {
			t.Errorf("got=%+v,want=%+v", p, &leaf{'c', 4})
		}
		if p := heap.Pop(table); !reflect.DeepEqual(p, &leaf{'e', 8}) {
			t.Errorf("got=%+v,want=%+v", p, &leaf{'e', 8})
		}
	}
}

func Test_buildTree(t *testing.T) {
	c := count("SARASARAHAIR")
	table := make(map[rune]string, len(c))
	tt := buildTree(c)
	fmt.Println(tt)
	createTable(tt, -1, 0, table)

	expect := map[rune]string{
		'A': "0",
		'R': "10",
		'S': "110",
		'H': "1110",
		'I': "1111",
	}
	fmt.Println("table")
	for k, v := range table {
		fmt.Printf("\tkey=%s, value=%v\n", string([]rune{k}), v)
	}
	for k, v := range table {
		expectValue, ok := expect[k]
		if !ok {
			t.Fatalf("invalid value: %v", k)
		}
		if expectValue != v {
			t.Errorf("got=%s, want=%s", v, expectValue)
		}
	}
}

func Test_Encode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{s: "ADBCBABCBBCE"},
			want: "1101110010011001000101111",
		},
		{
			name: "success",
			args: args{s: "SARASARAHAIR"},
			want: "1100100110010011100111110",
		},
		{
			name: "empty",
			args: args{s: ""},
			want: "",
		},
		{
			name: "1 kind word",
			args: args{s: "A"},
			want: "0",
		},
		{
			name: "1 kind 2 word",
			args: args{s: "AA"},
			want: "00",
		},
		{
			name: "2 kind 2 word",
			args: args{s: "AB"},
			want: "01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.s); got != tt.want {
				t.Errorf("encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
func BenchmarkEncode(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	s := randString(1000)

	m := map[rune]int{}
	for _, ss := range s {
		m[ss]++
	}
	// for k, v := range m {
	// 	fmt.Printf("%s:%d\n", string([]rune{k}), v)
	// }
	fmt.Println(len(m))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Encode(s)
	}
}
