package rel

import (
	"reflect"
)

type AsNode struct {
	Left  Visitable
	Right Visitable
	BaseVisitable
}

func (node AsNode) Eq(other AsNode) bool {
	return reflect.DeepEqual(node, other)
}
