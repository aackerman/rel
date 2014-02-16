package rel_test

import (
	. "."
	"testing"
)

func BenchmarkSelect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Select().ToSql()
	}
}
