package rel

import (
	"reflect"
)

type DescendingNode OrderingNode

func (node DescendingNode) Eq(other DescendingNode) bool {
	return reflect.DeepEqual(node, other)
}

func (node DescendingNode) Direction() string {
	return "DESC"
}

func (node DescendingNode) Reverse() *AscendingNode {
	return &AscendingNode{Expr: node.Expr}
}
