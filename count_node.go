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

func (node *CountNode) As(literal SqlLiteralNode) *AsNode {
	return aliasPredicationAs(node, literal)
}

func (node *CountNode) Over(visitable Visitable) *OverNode {
	return windowPredicationOver(node, visitable)
}

func (node CountNode) Eq(other CountNode) bool {
	return reflect.DeepEqual(node, other)
}
