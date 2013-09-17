package arel

type Visitor interface {
	Accept(Visitor) string
	Visit(Visitor) string
}

// Visitors accept an AST of nodes
// A visitor is a method and/or field on the connection
