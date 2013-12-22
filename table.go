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
	manager := NewSelectManager(t)
	return &manager
}

func (t *Table) Project(a ...interface{}) *SelectManager {
	return t.From().Project(a...)
}

func (t *Table) Alias(name string) {
	alias := NewTableAliasNode(t, name)
	t.Aliases = append(t.Aliases, alias.Name)
}

func (t *Table) Attr(name string) Attribute {
	return NewAttribute(t, name)
}
