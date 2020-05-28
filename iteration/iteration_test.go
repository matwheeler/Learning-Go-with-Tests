package iteration

import "testing"

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("expected: %q but got: %q", want, got)
		}
	}

	t.Run("repeat 'a' 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertCorrectMessage(t, got, want)
	})

	t.Run("repeat 'a' 2 times", func(t *testing.T) {
		got := Repeat("a", 2)
		want := "aa"
		assertCorrectMessage(t, got, want)
	})
}

// The testing.B gives you access to the cryptically named b.N.
// When the benchmark code is executed, it runs b.N times and measures how long it takes.
// The amount of times the code is run shouldn't matter to you,
// the framework will determine what is a "good" value for that to let you have some decent results.
// To run the benchmarks do go test -bench=. (or if you're in Windows Powershell go test -bench=".")
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
