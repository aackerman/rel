package arel

type ArelNode struct {
	Name  string
	Table *Table
	FactoryMethods
}
