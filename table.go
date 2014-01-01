package rel

import (
	"bytes"
	"strconv"
)

type Table struct {
	Name       string
	Engine     Engine
	TableAlias string
	Aliases    *[]TableAliasNode
	BaseVisitable
}

func NewTable(name string, e Engine) Table {
	table := Table{Name: name, Engine: e}
	return table
}

func (t *Table) Project(a ...Visitable) *SelectManager {
	return t.From(t).Project(a...)
}

func (t *Table) Select(a ...Visitable) *SelectManager {
	return t.Project(a...)
}

func (t *Table) Take(i int) *SelectManager {
	return t.From(t).Take(i)
}

func (t *Table) Where(n Visitable) *SelectManager {
	return t.From(t).Where(n)
}

func (t *Table) Skip(i int) *SelectManager {
	return t.From(t).Skip(i)
}

func (t *Table) Offset(i int) *SelectManager {
	return t.From(t).Offset(i)
}

func (t *Table) Having(a ...Visitable) *SelectManager {
	return t.From(t).Having(a...)
}

func (t *Table) Group(a ...Visitable) *SelectManager {
	return t.From(t).Group(a...)
}

func (t *Table) Order(exprs ...string) *SelectManager {
	return t.From(t).Order(exprs...)
}

func (t *Table) CreateStringJoin(left SqlLiteralNode) StringJoinNode {
	return StringJoinNode{Left: left}
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

func (t *Table) From(n *Table) *SelectManager {
	return t.SelectManager(n)
}

func (t *Table) SelectManager(n *Table) *SelectManager {
	return NewSelectManager(t.Engine, n)
}

func (t *Table) InsertManager() *InsertManager {
	manager := NewInsertManager(t)
	return &manager
}

func (t *Table) SetTableAlias(name string) {
	t.TableAlias = name
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
	alias := TableAliasNode{Name: buf.String(), Table: t}
	*t.Aliases = append(*t.Aliases, alias)
	return alias
}

func (t *Table) Attr(name string) AttributeNode {
	return NewAttributeNode(name, t)
}
