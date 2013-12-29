package arel

type EqualityNode struct {
	Right *Attribute
	Left  *Attribute
	BaseNode
}

func NewEqualityNode(right *Attribute, left *Attribute) EqualityNode {
	return EqualityNode{
		Right: right,
		Left:  left,
	}
}

type InNode EqualityNode
