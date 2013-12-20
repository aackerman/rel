package arel

type Table struct {
	Name       string
	Engine     Engine
	TableAlias string
	Aliases    []string
	SelectManager
}

func NewTable(name string, e Engine) Table {
	table := Table{Name: name, Engine: e}
	table.SelectManager = NewSelectManager(&table)
	return table
}

func (t *Table) Alias(name string) {
	alias := NewTableAliasNode(t, name)
	t.Aliases = append(t.Aliases, alias.Name)
}

func (t *Table) Attr(name string) Attribute {
	return NewAttribute(t, name)
}
