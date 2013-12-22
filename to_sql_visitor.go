package arel

import (
	"bytes"
	"log"
	"strings"
)

type ToSqlVisitor struct {
	conn *Connection
	BaseVisitor
}

const (
	WHERE    = " WHERE "
	SPACE    = " "
	COMMA    = ", "
	GROUP_BY = " GROUP BY "
	ORDER_BY = " ORDER BY "
	WINDOW   = " WINDOW "
	AND      = " AND "
	DISTINCT = "DISTINCT"
)

func NewToSqlVisitor(c *Connection) ToSqlVisitor {
	return ToSqlVisitor{conn: c, BaseVisitor: BaseVisitor{}}
}

func (v ToSqlVisitor) Accept(a AstNode) string {
	return v.Visit(a)
}

func (v ToSqlVisitor) Visit(a AstNode) string {
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

func (v ToSqlVisitor) VisitAndNode(a AndNode) string {
	return "AndNode"
}

func (v ToSqlVisitor) VisitInNode(a InNode) string {
	return "InNode"
}

func (v ToSqlVisitor) VisitSqlLiteralNode(a SqlLiteralNode) string {
	return a.Raw
}

func (v ToSqlVisitor) VisitSelectStatement(s SelectStatement) string {
	var buf bytes.Buffer
	if s.With != nil {
		buf.WriteString(v.Visit(s.With))
	}

	for _, core := range s.Cores() {
		v.VisitSelectCore(core)
	}

	// if s.Orders is not empty
	if len(s.Orders) > 0 {
		buf.WriteString(SPACE)
		buf.WriteString(ORDER_BY)
		for i, order := range s.Orders {
			buf.WriteString(s.VisitOrderNode(order))
			if (len(s.Orders) - 1) != i {
				buf.WriteString(COMMA)
			}
		}
	}

	if s.Limit != nil {
		buf.WriteString(SPACE)
		buf.WriteString(s.VisitLimitNode(s.Limit))
	}

	if s.Offset != nil {
		buf.WriteString(SPACE)
		buf.WriteString(s.VisitOffsetNode(s.Offset))
	}

	if s.Lock != nil {
		buf.WriteString(SPACE)
		buf.WriteString(s.VisitLockNode(s.Lock))
	}

	return strings.TrimSpace(buf.String())
}
