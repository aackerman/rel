package arel

type Table struct {
	Name       string
	Engine     Engine
	TableAlias string
	Aliases    []string
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

func (t *Table) Alias(name string) {
	alias := NewTableAliasNode(t, name)
	t.Aliases = append(t.Aliases, alias.Name)
}

func (t *Table) Attr(name string) AttributeNode {
	return NewAttributeNode(name, t)
}
