package rel

import (
	"reflect"
)

type DeleteStatementNode struct {
	Relation *Table
	Wheres   *[]Visitable
	BaseVisitable
}

func NewDeleteStatementNode() *DeleteStatementNode {
	return &DeleteStatementNode{
		Wheres: &[]Visitable{},
	}
}

func (node DeleteStatementNode) Eq(other DeleteStatementNode) bool {
	return reflect.DeepEqual(node, other)
}
