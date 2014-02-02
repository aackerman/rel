package rel

import (
	"bytes"
	"log"
	"strconv"
)

type Table struct {
	Name       string
	Engine     Engine
	TableAlias string
	Aliases    *[]*TableAliasNode
	BaseVisitable
}

func NewTable(name string) *Table {
	if RelEngine == nil {
		log.Fatal("Please register an engine before proceding")
	}
	return &Table{Name: name, Engine: RelEngine}
}

func (t *Table) String() string {
	return t.Name
}

func (t *Table) Project(visitables ...Visitable) *SelectManager {
	return t.Select(visitables...)
}

func (t *Table) Select(visitables ...Visitable) *SelectManager {
	return t.From(t).Select(visitables...)
}

func (t *Table) Take(i int) *SelectManager {
	return t.From(t).Take(i)
}

func (t *Table) Where(visitable Visitable) *SelectManager {
	return t.From(t).Where(visitable)
}

func (t *Table) Skip(i int) *SelectManager {
	return t.From(t).Skip(i)
}

func (t *Table) Offset(i int) *SelectManager {
	return t.From(t).Offset(i)
}

func (t *Table) Having(visitables ...Visitable) *SelectManager {
	return t.From(t).Having(visitables...)
}

func (t *Table) Group(visitables ...Visitable) *SelectManager {
	return t.From(t).Group(visitables...)
}

func (t *Table) Order(visitables ...Visitable) *SelectManager {
	return t.From(t).Order(visitables...)
}

func (t *Table) Join(visitable Visitable) *SelectManager {
	return t.From(t).InnerJoin(visitable)
}

func (t *Table) InnerJoin(visitable Visitable) *SelectManager {
	return t.From(t).InnerJoin(visitable)
}

func (t *Table) OuterJoin(visitable Visitable) *SelectManager {
	return t.From(t).OuterJoin(visitable)
}

func (t *Table) From(relation *Table) *SelectManager {
	return t.SelectManager(relation)
}

func (t *Table) SelectManager(relation *Table) *SelectManager {
	return NewSelectManager(t.Engine, relation)
}

func (t *Table) InsertManager() *InsertManager {
	return NewInsertManager(RelEngine)
}

func (t *Table) SetTableAlias(name string) {
	t.TableAlias = name
}

func (t *Table) Alias() *TableAliasNode {
	var buf bytes.Buffer

	// create the slice for aliases if it doesn't already exist
	if t.Aliases == nil {
		aliases := make([]*TableAliasNode, 0)
		t.Aliases = &aliases
	}

	// create the alias name
	n := len(*t.Aliases)
	buf.WriteString(t.Name)
	buf.WriteString("_")
	buf.WriteString(strconv.Itoa(n + 2))

	// create the alias
	alias := &TableAliasNode{Relation: t, Name: Sql(buf.String()), Quoted: true}

	// append the new alias to the list of current aliases
	*t.Aliases = append(*t.Aliases, alias)
	return alias
}

func (t *Table) Attr(name string) *AttributeNode {
	return NewAttributeNode(t, name)
}
