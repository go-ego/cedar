package cedar

import (
	"testing"

	"github.com/vcaesar/tt"
)

func init() {
	loadTestData()
}

func BenchmarkPrefixMatch(t *testing.B) {
	fn := func() {
		cd.PrefixMatch([]byte("abcdefg"), 0)
	}

	tt.BM(t, fn)
}

func BenchmarkInsert(t *testing.B) {
	fn := func() {
		cd.Insert([]byte("a"), 1)
		cd.Insert([]byte("b"), 3)
		cd.Insert([]byte("d"), 6)
	}

	tt.BM(t, fn)
}

func BenchmarkJump(t *testing.B) {
	fn := func() {
		cd.Jump([]byte("a"), 1)
	}

	tt.BM(t, fn)
}

func BenchmarkKey(t *testing.B) {
	fn := func() {
		cd.Key(1)
	}

	tt.BM(t, fn)
}

func BenchmarkValue(t *testing.B) {
	fn := func() {
		cd.Value(1)
	}

	tt.BM(t, fn)
}
