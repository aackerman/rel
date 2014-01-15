package rel_test

import (
	. "."
	"testing"
)

func TestCountNodeAs(t *testing.T) {
	table := NewTable("users")
	mgr := table.Select(table.Attr("id").Count().As(Sql("foo")))
	sql := mgr.ToSql()
	expected := "SELECT COUNT(\"users\".\"id\") AS foo FROM \"users\""
	if sql != expected {
		t.Logf("TestCoundNodeAs sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestCountNodeEq(t *testing.T) {
	count1 := CountNode{Expressions: []Visitable{Sql("foo")}}
	count2 := CountNode{Expressions: []Visitable{Sql("foo")}}
	if !count1.Eq(count2) {
		t.Fail()
	}
}
