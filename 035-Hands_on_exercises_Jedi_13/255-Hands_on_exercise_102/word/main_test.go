package word

import (
	"fmt"
	"testing"
)

func TestUseCount(t *testing.T) {
	type test struct {
		data 	string
		answer	int
	}
	tests := []test {
		test{"The Milky Way is big", 5},
		test{"The address of the Earth is Planet Earth, the Solar System, the Milky Way Galaxy", 15},
		test{"Mayall II is in the Andromeda Galaxy", 7},
	}
	for _, v := range tests {
		x := UseCount(v.data)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}

func TestCount(t *testing.T) {
	type test struct {
		data 	string
		answer	int
	}
	tests := []test {
		test{"The Milky Way is big", 5},
		test{"The address of the Earth is Planet Earth, the Solar System, the Milky Way Galaxy", 15},
		test{"Mayall II is in the Andromeda Galaxy", 7},
	}
	for _, v := range tests {
		x := Count(v.data)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}

func ExampleUseCount() {
	s := "The address of the Earth is Planet Earth, the Solar System, the Milky Way Galaxy"
	fmt.Println(UseCount(s))
	// Output:
	// 15
}

func ExampleCount() {
	s := "Mayall II is in the Andromeda Galaxy"
	fmt.Println(Count(s))
	// Output:
	// 7
}

func BenchmarkUseCount(b *testing.B) {
	var xs string = "The address of the Earth is Planet Earth, the Solar System, the Milky Way Galaxy"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UseCount(xs)
	}
}

func BenchmarkCount(b *testing.B) {
	var xs string = "The address of the Earth is Planet Earth, the Solar System, the Milky Way Galaxy"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Count(xs)
	}
}