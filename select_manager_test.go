package rel

import (
	"fmt"
	"testing"
)

func TestSelectManagerJoinSources(t *testing.T) {
	t.Skip("TestSelectManagerJoinSources not implemented")
}

func TestSelectManagerSkip(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	manager := table.From(&table)
	sql := manager.Skip(10).ToSql()
	if sql != "SELECT FROM \"users\" OFFSET 10" {
		t.Fail()
	}
}

func TestSelectManagerClone(t *testing.T) {
	t.Skip("TestSelectManagerClone not implemented")
}

func TestSelectManagerExists(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	manager := table.From(&table)
	manager.Project(Sql("*"))
	m2 := NewSelectManager(DefaultEngine, nil)
	m2.Project(manager.Exists())
	sql := m2.ToSql()
	expected := fmt.Sprintf("SELECT EXISTS (%s)", manager.ToSql())
	if sql != expected {
		t.Logf("TestSelectManagerExists sql: %s != %s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerExistsAs(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	manager := table.From(&table)
	manager.Project(Sql("*"))
	m2 := NewSelectManager(DefaultEngine, nil)
	m2.Project(manager.Exists().As(Sql("foo")))
	sql := m2.ToSql()
	expected := fmt.Sprintf("SELECT EXISTS (%s) AS foo", manager.ToSql())
	if sql != expected {
		t.Logf("TestSelectManagerExists sql: %s != %s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerOffset(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	manager := table.From(&table)
	sql := manager.Offset(10).ToSql()
	expected := "SELECT FROM \"users\" OFFSET 10"
	if sql != expected {
		t.Fail()
	}
}

func TestSelectManagerUnion(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	m1 := NewSelectManager(DefaultEngine, &table)
	m1.Project(Star())
	m1.Where(table.Attr("age").Lt(18))
	m2 := NewSelectManager(DefaultEngine, &table)
	m2.Project(Star())
	m2.Where(table.Attr("age").Gt(99))
	um := m1.Union(m2)
	sql := um.ToSql()
	expected := "( SELECT * FROM \"users\" WHERE \"users\".\"age\" < 18 UNION SELECT * FROM \"users\" WHERE \"users\".\"age\" > 99 )"
	if sql != expected {
		t.Logf("TestSelectManagerUnion sql: %s != %s", sql, expected)
		t.Fail()
	}
}
