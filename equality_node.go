package arel

type EqualityNode struct {
	Left  AstNode
	Right *AstNode
	BaseNode
}

func NewEqualityNode(left AstNode, right AstNode) EqualityNode {
	return EqualityNode{
		Left:  left,
		Right: &right,
	}
}

type InNode EqualityNode
