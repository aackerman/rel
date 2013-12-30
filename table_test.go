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

func TestTableAlias(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	alias := table.Alias()
	if alias.Name != "users_2" {
		t.Fail()
	}
}

func TestTableOrder(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	sm := table.Order("foo")
	sql := sm.ToSql()
	expected := "SELECT FROM \"users\" ORDER BY foo"
	if sql != expected {
		t.Logf("TestTableOrder sql: %s != %s", sql, expected)
		t.Fail()
	}
}

func TestTableProject(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	query := table.Project(Sql("*"))
	sql := query.ToSql()
	expected := "SELECT * FROM \"users\""
	if sql != expected {
		t.Logf("TestTableProject sql: %s != %s", sql, expected)
		t.Fail()
	}
}

func TestTableSkip(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	query := table.Skip(2)
	sql := query.ToSql()
	expected := "SELECT FROM \"users\" OFFSET 2"
	if sql != expected {
		t.Logf("TestTableSkip sql: %s != %s", sql, expected)
		t.Fail()
	}
}

func TestTableOffset(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	query := table.Offset(2)
	sql := query.ToSql()
	expected := "SELECT FROM \"users\" OFFSET 2"
	if sql != expected {
		t.Logf("TestTableOffset sql: %s != %s", sql, expected)
		t.Fail()
	}
}

func TestTableHaving(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	query := table.Having(table.Attr("id").Eq(10))
	sql := query.ToSql()
	expected := "SELECT FROM \"users\" HAVING \"users\".\"id\" = 10"
	if sql != expected {
		t.Logf("TestTableHaving sql: %s != %s", sql, expected)
		t.Fail()
	}
}

func TestTableGroup(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	query := table.Group(table.Attr("id"))
	sql := query.ToSql()
	expected := "SELECT FROM \"users\" GROUP BY \"users\".\"id\""
	if sql != expected {
		t.Logf("TestTableGroup sql: %s != %s", sql, expected)
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
