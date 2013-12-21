package arel

import "log"

type BaseVisitor struct{}

func (b BaseVisitor) Accept(a AstNode) string {
	return b.Visit(a)
}

func (b BaseVisitor) Visit(a AstNode) string {
	log.Printf("%T", a)
	switch val := a.(type) {
	case SelectStatement:
		return VisitSelectStatement(val)
	case AndNode:
		return VisitAndNode(val)
	case InNode:
		return VisitInNode(val)
	case SqlLiteralNode:
		return VisitSqlLiteralNode(val)
	}
	return ""
}

func VisitSelectStatement(s SelectStatement) string {
	return ""
}

func VisitAndNode(a AndNode) string {
	return "AndNode"
}

func VisitInNode(a InNode) string {
	return "InNode"
}

func VisitSqlLiteralNode(a SqlLiteralNode) string {
	return a.Raw
}

type Visitor interface {
	Accept(AstNode) string
	Visit(AstNode) string
}
