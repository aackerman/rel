package arel

// Base Visitor interface for visiting
// sql nodes and creating a buffer

type Visitor interface {
	Accept(AstNode) string
	Visit(AstNode) string
}
