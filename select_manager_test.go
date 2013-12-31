package arel

import (
	"testing"
)

func TestSelectManagerJoinSources(t *testing.T) {

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
	t.Fail()
}

func TestSelectManagerOffset(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	manager := table.SelectManager()
	sql := manager.Offset(10).ToSql()
	if sql != "SELECT FROM \"users\" OFFSET 10" {
		t.Fail()
	}
}
