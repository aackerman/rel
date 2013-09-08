package arel

type EqualityNode BinaryNode

func EqualityNodeNew() *EqualityNode {
	return &EqualityNode{}
}
