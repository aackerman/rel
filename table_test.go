package arel

import (
	"testing"
)

func TestTable(t *testing.T) {
	engine := NewEngine()
	table := NewTable("users", engine)
	query := table.Project(Sql("*"))
	sql := query.ToSql()
	if sql != "SELECT * FROM users" {
		t.Logf("TestTable sql: %s", sql)
		t.Fail()
	}
}
