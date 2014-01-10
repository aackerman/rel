package rel

import (
	"reflect"
)

type BinNode UnaryNode

func (node BinNode) Eq(other BinNode) bool {
	return reflect.DeepEqual(node, other)
}
