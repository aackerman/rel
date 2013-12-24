package arel

import (
	"testing"
)

func TestJoinSources(t *testing.T) {

}

func TestSkip(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	manager := table.SelectManager()
	sql := manager.Skip(10).ToSql()
	if sql != "SELECT FROM \"users\" OFFSET 10" {
		t.Fail()
	}
}

func TestOffset(t *testing.T) {
	table := NewTable("users", DefaultEngine)
	manager := table.SelectManager()
	sql := manager.Offset(10).ToSql()
	if sql != "SELECT FROM \"users\" OFFSET 10" {
		t.Fail()
	}
}
