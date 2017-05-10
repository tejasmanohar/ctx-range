package ctx_range

import (
	"context"
	"testing"
)

func TestRangeWorks(t *testing.T) {
	in := make(chan int, 5)
	out := make(chan int, 5)

	var expected []int
	for i := 0; i < 5; i++ {
		expected = append(expected, i)
		in <- i
	}

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	Range(ctx, in, out)
	close(in)

	i := 0
	for actual := range out {
		if expected[i] != actual {
			t.Errorf("expected elem %d to be %d, got %d", i, expected[i], actual)
		}
		i++
	}

	if i != 5 {
		t.Errorf("expected 5 items, got %d", i)
	}
}
