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

func BenchmarkSelectJoin(b *testing.B) {
	users := NewTable("users")
	preferences := NewTable("preferences")
	for i := 0; i < b.N; i++ {
		Select(Star()).Join(preferences).On(preferences.Attr("user_id").Eq(users.Attr("user_id"))).ToSql()
	}
}
