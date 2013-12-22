package arel

import (
	"testing"
)

func TestTable(t *testing.T) {
	engine := NewEngine()
	table := NewTable("users", engine)
	query := table.Project(Sql("*"))
	sql := query.ToSql()
	if sql != "SELECT * FROM \"users\"" {
		t.Logf("TestTable sql: %s", sql)
		t.Fail()
	}
}

func TestTableName(t *testing.T) {
	engine := NewEngine()
	table := NewTable("users", engine)
	if table.Name != "users" {
		t.Fail()
	}
}

func TestTableProjections(t *testing.T) {
	engine := NewEngine()
	table := NewTable("users", engine)
	query := table.Project(Sql("*"), Sql("*"))
	sql := query.ToSql()
	if sql != "SELECT *, * FROM \"users\"" {
		t.Logf("TestTable sql: %s", sql)
		t.Fail()
	}
}
