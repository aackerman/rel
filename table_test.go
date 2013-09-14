package arel

import (
	"testing"
)

func TestTable(t *testing.T) {
	engine := NewEngine()
	relation := NewTable("users", engine)
	query := relation.Select(Sql("*"))
	t.Log(query.ToSql())
}
