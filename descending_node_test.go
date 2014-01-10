package rel

import (
	"testing"
)

func TestDescendingNodeEq(t *testing.T) {
	desc1 := DescendingNode{Expr: Sql("zomg")}
	desc2 := DescendingNode{Expr: Sql("zomg")}
	if !desc1.Eq(desc2) {
		t.Fail()
	}
}

func TestDescendingNodeDirection(t *testing.T) {
	desc := DescendingNode{Expr: Sql("zomg")}
	if desc.Direction() != "DESC" {
		t.Fail()
	}
}

func TestDescendingNodeReverse(t *testing.T) {
	desc := DescendingNode{Expr: Sql("zomg")}
	asc := desc.Reverse()
	if asc.Expr != Sql("zomg") {
		t.Fail()
	}
}
