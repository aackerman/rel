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

func (t *Table) SelectManager() *SelectManager {
	manager := NewSelectManager(t)
	return &manager
}

func (t *Table) Alias(name string) {
	alias := NewTableAliasNode(t, name)
	t.Aliases = append(t.Aliases, alias.Name)
}

func (t *Table) Attr(name string) Attribute {
	return NewAttribute(t, name)
}
