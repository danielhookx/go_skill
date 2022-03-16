package test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestAbs(t *testing.T) {
	got := math.Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %f; want 1", got)
	}
}

//func TestMin(t *testing.T) {
//	got := math.Min(1,-1)
//	if got != 1 {
//		t.Errorf("Min(1, 2) = %f; want 1", got)
//	}
//}

func TestPrintUser(t *testing.T) {
	printUser := func(name string, age int) string {
		return fmt.Sprintf("%s is %d years old", name, age)
	}
	tests := []struct {
		name string
		User struct {
			name string
			age  int
		}
		id   string
		want string
	}{
		{
			name: "test1",
			User: struct {
				name string
				age  int
			}{
				name: "张三",
				age:  12,
			},
			id:   "u1",
			want: "张三 is 18 years old",
		},
		{
			name: "test2",
			User: struct {
				name string
				age  int
			}{
				name: "李四",
				age:  18,
			},
			id:   "u2",
			want: "李四 is 18 years old",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log(test)
			if got := printUser(test.User.name, test.User.age); got != test.want {
				t.Errorf("got %s; want %s", got, test.want)
			}
		})
	}
}

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}
