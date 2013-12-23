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

func TestTableProjection(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	query := table.Project(Sql("*"))
	sql := query.ToSql()
	if sql != "SELECT * FROM \"users\"" {
		t.Logf("TestTable sql: %s", sql)
		t.Fail()
	}
}

func TestTableMultipleProjections(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	query := table.Project(Sql("*"), Sql("*"))
	sql := query.ToSql()
	if sql != "SELECT *, * FROM \"users\"" {
		t.Logf("TestTableMultipleProjections sql: %s", sql)
		t.Fail()
	}
}

func TestTableSelectManager(t *testing.T) {
	table := NewTable("", DefaultEngine)
	sm := table.SelectManager()
	sql := sm.ToSql()
	if sql != "SELECT" {
		t.Logf("TestTableSelectManager sql: %s", sql)
		t.Fail()
	}
}
