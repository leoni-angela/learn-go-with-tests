package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := ExampleRepeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}

func Benchmark(b *testing.B) {
	// setup code
	for b.Loop() {
		ExampleRepeat("a", 5)
		// benchmarked code
	}
	// cleanup code
}