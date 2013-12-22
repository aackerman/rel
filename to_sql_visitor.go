package arel

import (
	"bytes"
	"log"
	"strings"
)

type ToSqlVisitor struct {
	conn *Connection
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
	return ToSqlVisitor{conn: c}
}

func (v ToSqlVisitor) Accept(a AstNode) string {
	return v.Visit(a)
}

func (v ToSqlVisitor) Visit(a AstNode) string {
	log.Printf("%T", a)
	switch val := a.(type) {
	case SelectStatement:
		return v.VisitSelectStatement(val)
	case AndNode:
		return v.VisitAndNode(val)
	case InNode:
		return v.VisitInNode(val)
	case SqlLiteralNode:
		return v.VisitSqlLiteralNode(val)
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

func (v ToSqlVisitor) VisitSelectCoreNode(s SelectCoreNode) string {
	var buf bytes.Buffer

	if s.Top > 0 {
		buf.WriteString(SPACE)
		buf.WriteString(v.VisitTopNode(s))
	}

	if s.SetQuanifier != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.VisitDistinctOnNode(s))
	}

	if len(s.Projections) > 0 {
		buf.WriteString(SPACE)
		for i, projection := range s.Projections {
			buf.WriteString(v.Visit(projection))
			if (len(s.Projections) - 1) != i {
				buf.WriteString(COMMA)
			}
		}
	}

	if s.Source != nil {
		buf.WriteString(" FROM ")
		buf.WriteString(v.Visit(s))
	}

	if len(s.Wheres) > 0 {
		buf.WriteString(WHERE)
		for i, where := range s.Wheres {
			buf.WriteString(v.Visit(where))
			if (len(s.Wheres) - 1) != i {
				buf.WriteString(COMMA)
			}
		}
	}

	if len(s.Groups) > 0 {
		buf.WriteString(WHERE)
		for i, group := range s.Groups {
			buf.WriteString(v.Visit(group))
			if (len(s.Groups) - 1) != i {
				buf.WriteString(COMMA)
			}
		}
	}

	if s.Having != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.Visit(s.Having))
	}

	if len(s.Windows) > 0 {
		buf.WriteString(WINDOW)
		for i, window := range s.Windows {
			buf.WriteString(v.Visit(window))
			if (len(s.Windows) - 1) != i {
				buf.WriteString(COMMA)
			}
		}
	}

	return buf.String()
}

func (v ToSqlVisitor) VisitSelectStatement(s SelectStatement) string {
	var buf bytes.Buffer
	if s.With != nil {
		buf.WriteString(v.Visit(s.With))
	}

	for _, core := range s.Cores() {
		v.VisitSelectCoreNode(core)
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
