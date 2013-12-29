package arel

import (
	"testing"
)

func TestTableName(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	if table.Name != "users" {
		t.Fail()
	}
}

func TestTableProject(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	query := table.Project(Sql("*"))
	sql := query.ToSql()
	expected := "SELECT * FROM \"users\""
	if sql != expected {
		t.Logf("TestTable sql: %s != %s", sql, expected)
		t.Fail()
	}
}

func TestTableSkip(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	query := table.Skip(2)
	sql := query.ToSql()
	expected := "SELECT FROM \"users\" OFFSET 2"
	if sql != expected {
		t.Logf("TestTable sql: %s != %s", sql, expected)
		t.Fail()
	}
}

func TestTableOffset(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	query := table.Offset(2)
	sql := query.ToSql()
	expected := "SELECT FROM \"users\" OFFSET 2"
	if sql != expected {
		t.Logf("TestTable sql: %s != %s", sql, expected)
		t.Fail()
	}
}

func TestTableHaving(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	query := table.Offset(2)
	sql := query.ToSql()
	expected := "SELECT FROM \"users\" OFFSET 2"
	if sql != expected {
		t.Logf("TestTable sql: %s != %s", sql, expected)
		t.Fail()
	}
}

func TestTableMultipleProjections(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	query := table.Project(Sql("*"), Sql("*"))
	sql := query.ToSql()
	expected := "SELECT *, * FROM \"users\""
	if sql != expected {
		t.Logf("TestTableMultipleProjections sql: '%s' != '%s'", sql, expected)
		t.Fail()
	}
}

func TestTableCreateStringJoin(t *testing.T) {
	table := NewTable("", DefaultEngine)
	join := table.CreateStringJoin("foo")
	if join.Left != "foo" {
		t.Log("TestTableCreateStringJoin join.Left.Name != \"foo\"")
		t.Fail()
	}
}

func TestTableSelectManager(t *testing.T) {
	table := NewTable("", DefaultEngine)
	sm := table.SelectManager()
	sql := sm.ToSql()
	expected := "SELECT"
	if sql != expected {
		t.Logf("TestTableSelectManager sql: %s != %s", sql, expected)
		t.Fail()
	}
}
