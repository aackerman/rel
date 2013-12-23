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
	ret := ""
	switch val := a.(type) {
	case SelectStatement:
		ret = v.VisitSelectStatement(val)
	case AndNode:
		ret = v.VisitAndNode(val)
	case InNode:
		ret = v.VisitInNode(val)
	case SqlLiteralNode:
		ret = v.VisitSqlLiteralNode(val)
	case JoinSource:
		ret = v.VisitJoinSourceNode(val)
	default:
		panic("ToSqlVisitor#Visit AstNode not handled")
	}
	// log.Printf("ToSqlVisitor#Visit; type of node: %T, return: %v", a, ret)
	return ret
}

func (v ToSqlVisitor) VisitTopNode(a TopNode) string {
	return "TopNode"
}

func (v ToSqlVisitor) VisitLimitNode(a LimitNode) string {
	return "LimitNode"
}

func (v ToSqlVisitor) VisitLockNode(a LockNode) string {
	return "LockNode"
}

func (v ToSqlVisitor) VisitOffsetNode(a OffsetNode) string {
	return "OffsetNode"
}

func (v ToSqlVisitor) VisitDistinctOnNode(a DistinctOnNode) string {
	return "DistinctOnNode"
}

func (v ToSqlVisitor) VisitAndNode(a AndNode) string {
	return "AndNode"
}

func (v ToSqlVisitor) VisitInNode(a InNode) string {
	return "InNode"
}

func (v ToSqlVisitor) VisitOrderingNode(a OrderingNode) string {
	return "OrderingNode"
}

func (v ToSqlVisitor) VisitTable(t *Table) string {
	var buf bytes.Buffer
	if len(t.TableAlias) > 0 {
		buf.WriteString(v.QuoteTableName(t.Name))
		buf.WriteString(SPACE)
		buf.WriteString(v.QuoteTableName(t.TableAlias))
	} else {
		buf.WriteString(v.QuoteTableName(t.Name))
	}
	return buf.String()
}

func (v ToSqlVisitor) QuoteTableName(name string) string {
	return v.conn.QuoteTableName(name)
}

func (v ToSqlVisitor) VisitJoinSourceNode(a JoinSource) string {
	var buf bytes.Buffer
	if a.Left != nil {
		log.Printf("VisitJoinSourceNode: %v", a.Left.Name)
		buf.WriteString(v.VisitTable(a.Left))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitSqlLiteralNode(a SqlLiteralNode) string {
	if len(a.Raw) > 0 {
		return a.Raw
	} else {
		return ""
	}
}

func (v ToSqlVisitor) VisitSelectCoreNode(s SelectCoreNode) string {
	var buf bytes.Buffer

	buf.WriteString("SELECT")

	if s.Top != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.VisitTopNode(*s.Top))
	}

	if s.SetQuanifier != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.Visit(*s.SetQuanifier))
	}

	if s.Projections != nil && len(*s.Projections) > 0 {
		claused := false
		for i, projection := range *s.Projections {

			if projection != nil {
				if !claused {
					buf.WriteString(SPACE)
					claused = true
				}
				buf.WriteString(v.Visit(projection))
				if (len(*s.Projections) - 1) != i {
					buf.WriteString(COMMA)
				}
			}
		}
	}

	if s.Source != nil {
		buf.WriteString(" FROM ")
		buf.WriteString(v.Visit(*s.Source))
	}

	if s.Wheres != nil && len(*s.Wheres) > 0 {
		claused := false
		for i, where := range *s.Wheres {
			if where != nil {
				if !claused {
					buf.WriteString(WHERE)
					claused = true
				}
				buf.WriteString(v.Visit(where))
				if (len(*s.Wheres) - 1) != i {
					buf.WriteString(COMMA)
				}
			}
		}
	}

	if s.Groups != nil && len(*s.Groups) > 0 {
		claused := false
		for i, group := range *s.Groups {
			if group != nil {
				if !claused {
					buf.WriteString(GROUP_BY)
					claused = true
				}
				buf.WriteString(v.Visit(group))
				if (len(*s.Groups) - 1) != i {
					buf.WriteString(COMMA)
				}
			}
		}
	}

	if s.Having != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.Visit(s.Having))
	}

	if s.Windows != nil && len(*s.Windows) > 0 {
		claused := false
		for i, window := range *s.Windows {
			if window != nil {
				if !claused {
					buf.WriteString(WINDOW)
					claused = true
				}
				buf.WriteString(v.Visit(window))
				if (len(*s.Windows) - 1) != i {
					buf.WriteString(COMMA)
				}
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

	if s.Cores != nil {
		for _, core := range s.Cores {
			if core != nil {
				buf.WriteString(v.VisitSelectCoreNode(*core))
			}
		}
	}

	if s.Orders != nil && len(*s.Orders) > 0 {
		buf.WriteString(SPACE)
		buf.WriteString(ORDER_BY)
		for i, order := range *s.Orders {
			buf.WriteString(v.VisitOrderingNode(order))
			if (len(*s.Orders) - 1) != i {
				buf.WriteString(COMMA)
			}
		}
	}

	if s.Limit != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.VisitLimitNode(*s.Limit))
	}

	if s.Offset != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.VisitOffsetNode(*s.Offset))
	}

	if s.Lock != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.VisitLockNode(*s.Lock))
	}

	return strings.TrimSpace(buf.String())
}
