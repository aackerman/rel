package rel

import (
	"reflect"
)

type CountNode FunctionNode

func (node *CountNode) Desc() *DescendingNode {
	return &DescendingNode{Expr: node}
}

func (node *CountNode) Asc() *AscendingNode {
	return &AscendingNode{Expr: node}
}

func (node *CountNode) As(v Visitable) *AsNode {
	return &AsNode{
		Left:  node,
		Right: v,
	}
}

func (node CountNode) Eq(other CountNode) bool {
	return reflect.DeepEqual(node, other)
}
