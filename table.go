package arel

type Table struct {
	Name       string
	engine     *Engine
	TableAlias string
	Aliases    []string
	*SelectManager
}

func NewTable(name string, e *Engine) *Table {
	table := &Table{
		Name:   name,
		engine: e,
	}

	table.SelectManager = NewSelectManager(e, table)
	return table
}

func (t *Table) Alias(name string) {
	alias := NewTableAliasNode(t, name)
	t.Aliases = append(t.Aliases, alias.Name)
}

func (t *Table) Attr(name string) Attribute {
	return NewAttribute(t, name)
}

// TODO: handle equality of []Aliases
func (t *Table) IsEqual(t2 *Table) bool {
	return t.Name == t2.Name &&
		t.engine == t2.engine &&
		// t.Aliases == t2.Aliases &&
		t.TableAlias == t2.TableAlias
}
