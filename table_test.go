package arel

import (
	"testing"
)

func TestTable(t *testing.T) {
	engine := NewEngine()
	table := NewTable("users", engine)
	query := table.Project(Sql("*"))
	if query.ToSql() != "SELECT * FROM users" {
		t.Log(query.ToSql())
		t.Fail()
	}
}
