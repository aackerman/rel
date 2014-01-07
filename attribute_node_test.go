package rel

import (
	"testing"
)

func TestAttributeNotEqSql(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").NotEq(Sql(10)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" != 10"
	if sql != expected {
		t.Logf("TestAttributeNotEqSql sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeNotEqNil(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").NotEq(nil))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" IS NOT NULL"
	if sql != expected {
		t.Logf("TestAttributeNotEqSql sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeGt(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").Gt(Sql(10)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" > 10"
	if sql != expected {
		t.Logf("TestAttributeNotEqSql sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}
