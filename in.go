package arel

type InNode EqualityNode

func NewInNode() *InNode {
	return &InNode{}
}
