package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("repeated %q expected %q", repeated, expected)
	}
}

func ExampleRepeat() {
	repeated := Repeat("c", 3)
	fmt.Println(repeated)
	// Output: ccc
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}

func BenchmarkStdRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Repeat("a", 6)
	}
}
