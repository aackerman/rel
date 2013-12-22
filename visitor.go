package arel

type Visitor interface {
	Accept(AstNode) string
	Visit(AstNode) string
}
