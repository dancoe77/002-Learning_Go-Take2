package mymath

import (
	"testing"
	"fmt"
)

func TestMyMath(t *testing.T) {
	type test struct {
		data 	[]int
		answer	float64
	}

	tests := []test {
		test{[]int{24, 25, 26, 27, 28}, 26},
		test{[]int{31, 32, 33, 34, 35}, 33},
		test{[]int{77, 78, 79, 80, 81}, 79},
		test{[]int{100, 101, 102, 103, 104}, 102},
	}

	for _, v := range tests {
		x := CenteredAvg(v.data)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}

func ExampleCenteredAvg() {
	xi := []int{24, 25, 26, 27, 28}
	fmt.Println(CenteredAvg(xi))
	// Output:
	// 26
}

func BenchmarkCenteredAvg(b *testing.B) {
	xi := []int{24, 25, 26, 27, 28}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CenteredAvg(xi)
	}
}