package arel

type Table struct {
	Name       string
	engine     *Engine
	TableAlias string
	Aliases    []string
	SelectManager
}

func NewTable(name string, engine *Engine) *Table {
	relation := &Table{
		Name:   name,
		engine: engine,
	}

	relation.SelectManager = NewSelectManager(engine, relation)
	return relation
}

func (relation *Table) Alias(name string) {
	alias := NewTableAliasNode(relation, name)
	relation.Aliases = append(relation.Aliases, alias.Name)
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
