package rel

import (
	"reflect"
)

type CountNode FunctionNode

func Count() *CountNode {
	return &CountNode{Expressions: []Visitable{Sql(1)}}
}

func (node *CountNode) Desc() *DescendingNode {
	return orderingDesc(node)
}

func (node *CountNode) Asc() *AscendingNode {
	return orderingAsc(node)
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
