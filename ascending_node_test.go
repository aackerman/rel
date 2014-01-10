package rel

import (
	"testing"
)

func TestAscendingNodeEq(t *testing.T) {
	asc1 := AscendingNode{Expr: Sql("zomg")}
	asc2 := AscendingNode{Expr: Sql("zomg")}
	if !asc1.Eq(asc2) {
		t.Fail()
	}
}

func TestAscendingNodeDirection(t *testing.T) {
	asc := AscendingNode{Expr: Sql("zomg")}
	if asc.Direction() != "ASC" {
		t.Fail()
	}
}

func TestAscendingNodeReverse(t *testing.T) {
	asc := AscendingNode{Expr: Sql("zomg")}
	desc := asc.Reverse()
	if desc.Expr != Sql("zomg") {
		t.Fail()
	}
}
