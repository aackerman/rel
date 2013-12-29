package arel

type EqualityNode struct {
	Left  SqlLiteralNode
	Right SqlLiteralNode
	BaseNode
}

func NewEqualityNode(right SqlLiteralNode, left SqlLiteralNode) EqualityNode {
	return EqualityNode{
		Left:  left,
		Right: right,
	}
}

type InNode EqualityNode
