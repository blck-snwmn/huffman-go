package huffmango

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

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
			name: "empty",
			args: args{s: ""},
			want: "",
		},
		{
			name: "empty",
			args: args{s: "A"},
			want: "0",
		},
		{
			name: "empty",
			args: args{s: "AA"},
			want: "00",
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
