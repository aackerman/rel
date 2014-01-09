package rel

import (
	"reflect"
)

type AndNode struct {
	Children *[]Visitable
	BaseVisitable
}

func (node AndNode) Eq(other AndNode) bool {
	return reflect.DeepEqual(node, other)
}
