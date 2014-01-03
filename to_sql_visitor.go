package rel

import (
	"bytes"
	"log"
	"runtime/debug"
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

func (v ToSqlVisitor) Accept(a Visitable) string {
	return v.Visit(a)
}

func (v ToSqlVisitor) Visit(a Visitable) string {
	ret := ""
	switch val := a.(type) {
	case SelectStatementNode:
		ret = v.VisitSelectStatementNode(val)
	case AndNode:
		ret = v.VisitAndNode(val)
	case InNode:
		ret = v.VisitInNode(val)
	case SqlLiteralNode:
		ret = v.VisitSqlLiteralNode(val)
	case JoinSource:
		ret = v.VisitJoinSourceNode(val)
	case EqualityNode:
		ret = v.VisitEqualityNode(val)
	case HavingNode:
		ret = v.VisitHavingNode(val)
	case AttributeNode:
		ret = v.VisitAttributeNode(val)
	case GroupNode:
		ret = v.VisitGroupNode(val)
	case ExistsNode:
		ret = v.VisitExistsNode(val)
	case AsNode:
		ret = v.VisitAsNode(val)
	case Table:
		ret = v.VisitTable(val)
	case *Table:
		ret = v.VisitTable(*val)
	case LessThanNode:
		ret = v.VisitLessThanNode(val)
	case UnionNode:
		ret = v.VisitUnionNode(val)
	case SelectManager:
		ret = v.VisitSelectManager(val)
	case GreaterThanNode:
		ret = v.VisitGreaterThanNode(val)
	default:
		debug.PrintStack()
		log.Fatalf("ToSqlVisitor#Visit %T not handled", a)
	}
	return ret
}

func (v ToSqlVisitor) VisitTopNode(a TopNode) string {
	return "TopNode"
}

func (v ToSqlVisitor) VisitLimitNode(a LimitNode) string {
	var buf bytes.Buffer
	buf.WriteString("LIMIT ")
	buf.WriteString(v.Visit(a.Expr))
	return buf.String()
}

func (v ToSqlVisitor) VisitLockNode(a LockNode) string {
	return "LockNode"
}

func (v ToSqlVisitor) VisitOffsetNode(n OffsetNode) string {
	return "OFFSET " + v.Visit(n.Expr)
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

func (v ToSqlVisitor) VisitSelectManager(a SelectManager) string {
	var buf bytes.Buffer
	buf.WriteString("(")
	buf.WriteString(a.ToSql())
	buf.WriteString(")")
	return buf.String()
}

func (v ToSqlVisitor) VisitUnionNode(a UnionNode) string {
	var buf bytes.Buffer
	buf.WriteString("( ")
	buf.WriteString(v.Visit(a.Left))
	buf.WriteString(" UNION ")
	buf.WriteString(v.Visit(a.Right))
	buf.WriteString(" )")
	return buf.String()
}

func (v ToSqlVisitor) VisitLessThanNode(a LessThanNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(a.Left))
	buf.WriteString(" < ")
	buf.WriteString(v.Visit(a.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitGreaterThanNode(a GreaterThanNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(a.Left))
	buf.WriteString(" > ")
	buf.WriteString(v.Visit(a.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitAsNode(a AsNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(a.Left))
	buf.WriteString(" AS ")
	buf.WriteString(v.Visit(*a.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitGroupNode(n GroupNode) string {
	return v.Visit(n.Expr)
}

func (v ToSqlVisitor) VisitHavingNode(n HavingNode) string {
	var buf bytes.Buffer
	buf.WriteString("HAVING ")
	buf.WriteString(v.Visit(n.Expr))
	return buf.String()
}

func (v ToSqlVisitor) VisitExistsNode(n ExistsNode) string {
	var buf bytes.Buffer
	buf.WriteString("EXISTS (")
	buf.WriteString(v.Visit(n.Expressions))
	buf.WriteString(")")
	if n.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(n.Alias))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitAttributeNode(n AttributeNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.QuoteTableName(n.Table.Name))
	buf.WriteString(".")
	buf.WriteString(v.QuoteColumnName(n.Name))
	return buf.String()
}

func (v ToSqlVisitor) VisitEqualityNode(n EqualityNode) string {
	var buf bytes.Buffer
	if n.Right == nil {
		buf.WriteString(v.Visit(n.Left))
		buf.WriteString(" IS NULL")
	} else {
		buf.WriteString(v.Visit(n.Left))
		buf.WriteString(" = ")
		buf.WriteString(v.Visit(*n.Right))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitTable(t Table) string {
	var buf bytes.Buffer
	if t.TableAlias != "" {
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

func (v ToSqlVisitor) QuoteColumnName(name string) string {
	return v.conn.QuoteColumnName(name)
}

func (v ToSqlVisitor) VisitJoinSourceNode(a JoinSource) string {
	var buf bytes.Buffer
	if a.Left != nil {
		buf.WriteString(v.Visit(a.Left))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitSqlLiteralNode(a SqlLiteralNode) string {
	if a.Raw != "" {
		return a.Raw
	} else {
		return ""
	}
}

func (v ToSqlVisitor) VisitSelectCoreNode(s SelectCoreNode) string {
	var buf bytes.Buffer

	buf.WriteString("SELECT")

	// Add TOP statement to the buffer
	if s.Top != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.VisitTopNode(*s.Top))
	}

	if s.SetQuanifier != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.Visit(*s.SetQuanifier))
	}

	// add select projections
	if s.Projections != nil {
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

	// add FROM statement to the buffer
	if s.Source != nil && s.Source.Left != nil {
		if t, ok := s.Source.Left.(Table); ok && t.Name != "" {
			buf.WriteString(" FROM ")
			buf.WriteString(v.Visit(*s.Source))
		} else if t, ok := s.Source.Left.(*Table); ok && t.Name != "" {
			buf.WriteString(" FROM ")
			buf.WriteString(v.Visit(*s.Source))
		}
	}

	// add WHERE statement to the buffer
	if s.Wheres != nil {
		claused := false
		for i, where := range *s.Wheres {
			// add WHERE clause if it hasn't already been added
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

	// add GROUP BY statement to the buffer
	if s.Groups != nil {
		claused := false
		for i, group := range *s.Groups {
			// add GROUP BY clause if it hasn't already been added
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

	// add HAVING statement to the buffer
	if s.Having != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.VisitHavingNode(*s.Having))
	}

	// add WINDOW statements to the buffer
	if s.Windows != nil {
		claused := false
		for i, window := range *s.Windows {
			// add WINDOW clause if is hasn't already been added
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

	return buf.String()
}

func (v ToSqlVisitor) VisitSelectStatementNode(s SelectStatementNode) string {
	var buf bytes.Buffer

	// add WITH statement to the buffer
	if s.With != nil {
		buf.WriteString(v.Visit(s.With))
	}

	// add SELECT core to the buffer
	if s.Cores != nil {
		for _, core := range s.Cores {
			if core != nil {
				buf.WriteString(v.VisitSelectCoreNode(*core))
			}
		}
	}

	// add ORDER BY clauses to the buffer
	if s.Orders != nil {
		buf.WriteString(ORDER_BY)
		for i, order := range *s.Orders {
			buf.WriteString(v.Visit(order))
			if (len(*s.Orders) - 1) != i {
				buf.WriteString(COMMA)
			}
		}
	}

	// add LIMIT clause to the buffer
	if s.Limit != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.VisitLimitNode(*s.Limit))
	}

	// add OFFSET clause to the buffer
	if s.Offset != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.VisitOffsetNode(*s.Offset))
	}

	// add LOCK clause to the buffer
	if s.Lock != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.VisitLockNode(*s.Lock))
	}

	return strings.TrimSpace(buf.String())
}
