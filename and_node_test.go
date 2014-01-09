package rel

import (
	"testing"
)

func TestAndNodeEq(t *testing.T) {
	and1 := AndNode{Children: &[]Visitable{Sql("foo"), Sql("bar")}}
	and2 := AndNode{Children: &[]Visitable{Sql("foo"), Sql("bar")}}
	if !and1.Eq(and2) {
		t.Fail()
	}
}
