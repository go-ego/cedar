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
