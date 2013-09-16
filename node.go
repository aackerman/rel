package arel

type Node interface {
	NodeInterface()
}

type ArelNode struct {
	Name  string
	Table *Table
}
