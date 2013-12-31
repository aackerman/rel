package arel

import (
	"fmt"
	"testing"
)

func TestSelectManagerJoinSources(t *testing.T) {
	t.Log("TestSelectManagerJoinSources not implemented")
	t.Fail()
}

func TestSelectManagerSkip(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	manager := table.SelectManager()
	sql := manager.Skip(10).ToSql()
	if sql != "SELECT FROM \"users\" OFFSET 10" {
		t.Fail()
	}
}

func TestSelectManagerClone(t *testing.T) {
	t.Log("TestSelectManagerClone not implemented")
	t.Fail()
}

func TestSelectManagerExists(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	manager := table.SelectManager()
	manager.Project(Sql("*"))
	m2 := table.SelectManager()
	m2.Project(manager.Exists())
	sql := m2.ToSql()
	expected := fmt.Sprintf("SELECT EXISTS (%s)", manager.ToSql())
	if sql != expected {
		t.Log("TestSelectManagerClone not implemented")
		t.Fail()
	}
}

func TestSelectManagerOffset(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	manager := table.SelectManager()
	sql := manager.Offset(10).ToSql()
	if sql != "SELECT FROM \"users\" OFFSET 10" {
		t.Fail()
	}
}
