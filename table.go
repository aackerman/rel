package arel

import (
	"fmt"
)

type Table struct {
	Name       string
	Engine     Engine
	TableAlias string
	Aliases    []string
}

func NewTable(name string, e Engine) *Table {
	return &Table{
		Name:   name,
		Engine: e,
	}
}

func (t *Table) From() *SelectManager {
	return t.SelectManager()
}

func (t *Table) Alias(name string) {
	alias := NewTableAliasNode(t, name)
	t.Aliases = append(t.Aliases, alias.Name)
}

func (t *Table) Join(relation fmt.Stringer) *SelectManager {
	return t.From().Join(relation)
}

// TODO: replace empty interface type
// TODO: add return type
func (t *Table) Group(columns ...interface{}) *SelectManager {
	return t.From().Group(columns)
}

// TODO: replace empty interface type
// TODO: add return type
func (t *Table) Order(expr ...interface{}) *SelectManager {
	return t.From().Order(expr)
}

// TODO: replace empty interface type
// TODO: add return type
func (t *Table) Where(condition interface{}) *SelectManager {
	return t.From().Where(condition)
}

// TODO: replace empty interface type
// TODO: add return type
func (t *Table) Project(things ...interface{}) *SelectManager {
	return t.From().Project(things)
}

func (t *Table) Select(things ...interface{}) *SelectManager {
	return t.From().Select(things)
}

// TODO: add return type
func (t *Table) Take(amount int) *SelectManager {
	return t.From().Take(amount)
}

func (t *Table) Limit(amount int) *SelectManager {
	return t.From().Take(amount)
}

// TODO: add return type
func (t *Table) Skip(amount int) *SelectManager {
	return t.From().Skip(amount)
}

func (t *Table) Offset(amount int) *SelectManager {
	return t.From().Skip(amount)
}

// TODO: add input type
// TODO: add return type
func (t *Table) Having(expr ...interface{}) *SelectManager {
	return t.From().Having(expr)
}

func (t *Table) SelectManager() *SelectManager {
	return NewSelectManager(t.Engine, t)
}

func (t *Table) InsertManager() *InsertManager {
	return NewInsertManager(t.Engine)
}

func (t *Table) Attr(name string) *Attribute {
	return NewAttribute(t, name)
}

// TODO: handle equality of []Aliases
func (t *Table) IsEqual(t2 *Table) bool {
	return t.Name == t2.Name &&
		t.Engine == t2.Engine &&
		// t.Aliases == t2.Aliases &&
		t.TableAlias == t2.TableAlias
}
