package arel

import (
	"bytes"
	"strconv"
)

type Table struct {
	Name    string
	Engine  Engine
	Aliases *[]TableAliasNode
}

func NewTable(name string, e Engine) Table {
	table := Table{Name: name, Engine: e}
	return table
}

func (t *Table) From() *SelectManager {
	return t.SelectManager()
}

func (t *Table) Project(a ...AstNode) *SelectManager {
	return t.From().Project(a...)
}

func (t *Table) Skip(i int) *SelectManager {
	return t.From().Skip(i)
}

func (t *Table) Offset(i int) *SelectManager {
	return t.From().Offset(i)
}

func (t *Table) Having(a ...AstNode) *SelectManager {
	return t.From().Having(a...)
}

func (t *Table) Group(a ...AstNode) *SelectManager {
	return t.From().Group(a...)
}

func (t *Table) Order(a ...AstNode) *SelectManager {
	return t.From().Order(a...)
}

func (t *Table) CreateStringJoin(left string) StringJoinNode {
	return StringJoinNode{
		Left: left,
	}
}

func (t *Table) CreateInnerJoin(left *Table, right *Table) InnerJoinNode {
	return InnerJoinNode{
		Left:  left,
		Right: right,
	}
}

func (t *Table) CreateOuterJoin(left *Table, right *Table) OuterJoinNode {
	return OuterJoinNode{
		Left:  left,
		Right: right,
	}
}

func (t *Table) SelectManager() *SelectManager {
	manager := NewSelectManager(t)
	return &manager
}

func (t *Table) InsertManager() *InsertManager {
	manager := NewInsertManager(t)
	return &manager
}

func (t *Table) Alias() TableAliasNode {
	var buf bytes.Buffer
	if t.Aliases == nil {
		aliases := make([]TableAliasNode, 0)
		t.Aliases = &aliases
	}
	n := len(*t.Aliases)
	buf.WriteString(t.Name)
	buf.WriteString("_")
	buf.WriteString(strconv.Itoa(n + 2))
	alias := NewTableAliasNode(t, buf.String())
	*t.Aliases = append(*t.Aliases, alias)
	return alias
}

func (t *Table) Attr(name string) AttributeNode {
	return NewAttributeNode(name, t)
}
