package dog

import (
	"testing"
	"fmt"
)

func TestYears(t *testing.T) {
	type test struct {
		data	int
		answer	int
	}
	tests := []test{
		test{2, 14},
		test{3, 21},
		test{4, 28},
		test{5, 35},
	}
	for _, v := range tests {
		x := Years(v.data)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}

func TestYearsTwo(t *testing.T) {
	type test struct {
		data	int
		answer	int
	}
	tests := []test{
		test{2, 14},
		test{3, 21},
		test{4, 28},
		test{5, 35},
	}
	for _, v := range tests {
		x := YearsTwo(v.data)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}

func ExampleYears() {
	n := 7
	fmt.Println(Years(n))
	// Output:
	// 49
}

func ExampleYearsTwo() {
	n := 7
	fmt.Println(YearsTwo(n))
	// Output:
	// 49
}

func BenchmarkYears(b *testing.B) {
	var xi int = 2
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Years(xi)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	var xi int = 2
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		YearsTwo(xi)
	}
}