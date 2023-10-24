package slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := Sum(numbers)
	want := 55

	if got != want {
		t.Errorf("got %d, want %d, given %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	numbers1 := []int{1, 2, 3}
	numbers2 := []int{4, 5, 6}
	numbers3 := []int{7, 8, 9}
	numbers4 := []int{10}

	got := SumAll(numbers1, numbers2, numbers3, numbers4)
	want := []int{6, 15, 24, 10}

	checkSums(t, got, want, [][]int{numbers1, numbers2, numbers3, numbers4})
}

func BenchmarkTestSumAll(b *testing.B) {
	numbers1 := []int{1, 2, 3}
	numbers2 := []int{4, 5, 6}
	numbers3 := []int{7, 8, 9}
	numbers4 := []int{10}

	for i := 0; i < b.N; i++ {
		SumAll(numbers1, numbers2, numbers3, numbers4)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sum of some slices", func(t *testing.T) {
		numbers1 := []int{1, 2, 3}
		numbers2 := []int{4, 5, 6}

		got := SumAllTails(numbers1, numbers2)
		want := []int{5, 11}

		checkSums(t, got, want, [][]int{numbers1, numbers2})
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		numbers1 := []int{1}
		numbers2 := []int{}

		got := SumAllTails(numbers1, numbers2)
		want := []int{0, 0}

		checkSums(t, got, want, [][]int{numbers1, numbers2})
	})
}

func checkSums(t *testing.T, got []int, want []int, given [][]int) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v, given %v", got, want, given)
	}
}
