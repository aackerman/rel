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
	// table = Table.new(:users)
	// manager = Arel::SelectManager.new Table.engine, table
	// manager.project SqlLiteral.new '*'
	// m2 = Arel::SelectManager.new(manager.engine)
	// m2.project manager.exists
	// m2.to_sql.must_be_like %{ SELECT EXISTS (#{manager.to_sql}) }
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

func TestSelectManagerOffset(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	manager := table.From(&table)
	sql := manager.Offset(10).ToSql()
	expected := "SELECT FROM \"users\" OFFSET 10"
	if sql != expected {
		t.Fail()
	}
}
