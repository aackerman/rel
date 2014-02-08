package rel

import (
	"reflect"
)

type AscendingNode OrderingNode

func (node AscendingNode) Eq(other AscendingNode) bool {
	return reflect.DeepEqual(node, other)
}

func (node *AscendingNode) Direction() string {
	return "ASC"
}

func (node *AscendingNode) Reverse() *DescendingNode {
	return &DescendingNode{Expr: node.Expr}
}
