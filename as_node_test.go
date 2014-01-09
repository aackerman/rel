package rel

import (
	"testing"
)

func TestAsNodeEq(t *testing.T) {
	as1 := AsNode{Left: Sql("foo"), Right: Sql("bar")}
	as2 := AsNode{Left: Sql("foo"), Right: Sql("bar")}
	if !as1.Eq(as2) {
		t.Fail()
	}
}
