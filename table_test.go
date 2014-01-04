package rel

import (
	"testing"
)

func TestTableName(t *testing.T) {
	table := NewTable("users")
	if table.Name != "users" {
		t.Fail()
	}
}

func TestTableAlias(t *testing.T) {
	table := NewTable("users")
	alias := table.Alias()
	if alias.Name != "users_2" {
		t.Fail()
	}
}

func TestTableSetTableAlias(t *testing.T) {
	table := NewTable("users")
	table.SetTableAlias("foo")
	manager := table.From(table)
	manager.Skip(10)
	sql := manager.ToSql()
	expected := "SELECT FROM \"users\" \"foo\" OFFSET 10"
	if sql != expected {
		t.Logf("TestTableSetTableAlias sql: %s != %s", sql, expected)
		t.Fail()
	}
}

func TestTableOrder(t *testing.T) {
	table := NewTable("users")
	sm := table.Order("foo")
	sql := sm.ToSql()
	expected := "SELECT FROM \"users\" ORDER BY foo"
	if sql != expected {
		t.Logf("TestTableOrder sql: %s != %s", sql, expected)
		t.Fail()
	}
}

func TestTableTake(t *testing.T) {
	table := NewTable("users")
	sm := table.Take(1)
	sm.Project(Sql("*"))
	sql := sm.ToSql()
	expected := "SELECT * FROM \"users\" LIMIT 1"
	if sql != expected {
		t.Logf("TestTableOrder sql: %s != %s", sql, expected)
		t.Fail()
	}
}

func TestTableWhere(t *testing.T) {
	table := NewTable("users")
	sm := table.Where(table.Attr("id").Eq(Sql(1)))
	sm.Project(table.Attr("id"))
	sql := sm.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" = 1"
	if sql != expected {
		t.Logf("TestTableWhere sql: %s != %s", sql, expected)
		t.Fail()
	}
}

func TestTableProject(t *testing.T) {
	table := NewTable("users")
	query := table.Project(Sql("*"))
	sql := query.ToSql()
	expected := "SELECT * FROM \"users\""
	if sql != expected {
		t.Logf("TestTableProject sql: %s != %s", sql, expected)
		t.Fail()
	}
}

func TestTableSkip(t *testing.T) {
	table := NewTable("users")
	query := table.Skip(2)
	sql := query.ToSql()
	expected := "SELECT FROM \"users\" OFFSET 2"
	if sql != expected {
		t.Logf("TestTableSkip sql: %s != %s", sql, expected)
		t.Fail()
	}
}

func TestTableOffset(t *testing.T) {
	table := NewTable("users")
	query := table.Offset(2)
	sql := query.ToSql()
	expected := "SELECT FROM \"users\" OFFSET 2"
	if sql != expected {
		t.Logf("TestTableOffset sql: %s != %s", sql, expected)
		t.Fail()
	}
}

func TestTableHaving(t *testing.T) {
	table := NewTable("users")
	query := table.Having(table.Attr("id").Eq(Sql(10)))
	sql := query.ToSql()
	expected := "SELECT FROM \"users\" HAVING \"users\".\"id\" = 10"
	if sql != expected {
		t.Logf("TestTableHaving sql: %s != %s", sql, expected)
		t.Fail()
	}
}

func TestTableGroup(t *testing.T) {
	table := NewTable("users")
	query := table.Group(table.Attr("id"))
	sql := query.ToSql()
	expected := "SELECT FROM \"users\" GROUP BY \"users\".\"id\""
	if sql != expected {
		t.Logf("TestTableGroup sql: %s != %s", sql, expected)
		t.Fail()
	}
}

func TestTableMultipleProjections(t *testing.T) {
	table := NewTable("users")
	query := table.Project(Sql("*"), Sql("*"))
	sql := query.ToSql()
	expected := "SELECT *, * FROM \"users\""
	if sql != expected {
		t.Logf("TestTableMultipleProjections sql: '%s' != '%s'", sql, expected)
		t.Fail()
	}
}

func TestTableNewStringJoin(t *testing.T) {
	table := NewTable("")
	join := table.NewStringJoin(Sql("foo"))
	if join.Left.Raw != "foo" {
		t.Log("TestTableNewStringJoin join.Left != \"foo\"")
		t.Fail()
	}
}

func TestTableEquality(t *testing.T) {
	t.SkipNow()
}

func TestTableSelectManager(t *testing.T) {
	table := NewTable("")
	sm := table.From(table)
	sql := sm.ToSql()
	expected := "SELECT"
	if sql != expected {
		t.Logf("TestTableSelectManager sql: %s != %s", sql, expected)
		t.Fail()
	}
}
