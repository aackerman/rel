package rel_test

import (
	. "."
	"testing"
)

func TestInsertManager(t *testing.T) {
	Register("postgresql", NewEngine())
	sql := ""
	expected := ""
	if sql != expected {
		t.Logf("TestInsertManager sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestInsertManagerCreateValues(t *testing.T) {
	mgr := NewInsertManager(DefaultEngine)
	values := make([]interface{}, 0)
	values = append(values, Star())
	columns := make([]AttributeNode, 0)
	columns = append(columns, AttributeNode{})
	mgr.SetValues(mgr.CreateValues(values, columns))
	sql := mgr.ToSql()
	expected := "INSERT INTO NULL VALUES (*)"
	if sql != expected {
		t.Logf("TestInsertManager sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestInsertManagerInsertsFalse(t *testing.T) {
	users := NewTable("users")
	mgr := NewInsertManager(DefaultEngine)
	mgr.Ast.Relation = users
	mgr.Insert(users.Attr("bool"), false)
	sql := mgr.ToSql()
	expected := "INSERT INTO \"users\" (\"bool\") VALUES ('f')"
	if sql != expected {
		t.Logf("TestInsertManager sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}
