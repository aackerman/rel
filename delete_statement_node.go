package rel

import (
	"reflect"
)

type DeleteStatementNode struct {
	Relation *Table
	Wheres   *[]Visitable
	BaseVisitable
}

func (node DeleteStatementNode) Eq(other DeleteStatementNode) bool {
	return reflect.DeepEqual(node, other)
}
