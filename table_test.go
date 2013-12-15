package arel

import (
	"testing"
)

func TestTable(t *testing.T) {
	engine := NewEngine()
	relation := NewTable("users", engine)
	query := relation.Project(Sql("*"))
	if query.ToSql() != "SELECT * FROM users" {
		t.Log(query.ToSql())
		t.Fail()
	}
}
