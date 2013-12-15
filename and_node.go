package arel

type AndNode struct {
	Name  string
	Table *Table
	AstNode
}

func NewAndNode() AndNode {
	return AndNode{}
}
