package rel

import (
	"testing"
)

func TestInsertManager(t *testing.T) {
	sql := ""
	expected := ""
	if sql != expected {
		t.Logf("TestInsertManager sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestInsertManagerCreateValues(t *testing.T) {
	mgr := NewInsertManager(TableEngine)
	values := make([]Visitable, 0)
	values = append(values, Star())
	columns := make([]SqlLiteralNode, 0)
	columns = append(columns, Sql("a"))
	mgr.SetValues(mgr.CreateValues(values, columns))
	sql := mgr.ToSql()
	expected := "INSERT INTO NULL VALUES (*)"
	if sql != expected {
		t.Logf("TestInsertManager sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}
