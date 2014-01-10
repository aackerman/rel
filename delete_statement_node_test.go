package rel

import (
	"testing"
)

func TestDeleteStatementNodeEq(t *testing.T) {
	ds1 := DeleteStatementNode{
		Wheres: &[]Visitable{Sql("a"), Sql("b"), Sql("c")},
	}

	ds2 := DeleteStatementNode{
		Wheres: &[]Visitable{Sql("a"), Sql("b"), Sql("c")},
	}

	if !ds1.Eq(ds2) {
		t.Fail()
	}
}

func TestDeleteStatementNodeNotEq(t *testing.T) {
	ds1 := DeleteStatementNode{
		Wheres: &[]Visitable{Sql("a"), Sql("b"), Sql("c")},
	}

	ds2 := DeleteStatementNode{
		Wheres: &[]Visitable{Sql("1"), Sql("2"), Sql("3")},
	}

	if ds1.Eq(ds2) {
		t.Fail()
	}
}
