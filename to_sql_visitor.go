package rel

import (
	"bytes"
	"log"
	"runtime/debug"
	"strings"
)

type ToSqlVisitor struct {
	Conn Connector
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

func (v ToSqlVisitor) Accept(visitable Visitable) string {
	return v.Visit(visitable)
}

// FIXME: Only visit pointers to visitables
func (v ToSqlVisitor) Visit(visitable Visitable) string {
	switch val := visitable.(type) {
	case nil:
		return v.VisitNil()
	case *SelectStatementNode:
		return v.VisitSelectStatementNode(val)
	case *InNode:
		return v.VisitInNode(val)
	case SqlLiteralNode:
		return v.VisitSqlLiteralNode(val)
	case *SqlLiteralNode:
		return v.VisitSqlLiteralNode(*val)
	case JoinSource:
		return v.VisitJoinSourceNode(val)
	case *EqualityNode:
		return v.VisitEqualityNode(val)
	case *HavingNode:
		return v.VisitHavingNode(val)
	case *AttributeNode:
		return v.VisitAttributeNode(val)
	case *GroupNode:
		return v.VisitGroupNode(val)
	case *ExistsNode:
		return v.VisitExistsNode(val)
	case *AsNode:
		return v.VisitAsNode(val)
	case *LessThanNode:
		return v.VisitLessThanNode(val)
	case *UnionNode:
		return v.VisitUnionNode(val)
	case *UnionAllNode:
		return v.VisitUnionAllNode(val)
	case *SelectManager:
		return v.VisitSelectManager(val)
	case *GreaterThanNode:
		return v.VisitGreaterThanNode(val)
	case *IntersectNode:
		return v.VisitIntersectNode(val)
	case *ExceptNode:
		return v.VisitExceptNode(val)
	case *OnNode:
		return v.VisitOnNode(val)
	case *AscendingNode:
		return v.VisitAscendingNode(val)
	case *DescendingNode:
		return v.VisitDescendingNode(val)
	case *CountNode:
		return v.VisitCountNode(val)
	case *AndNode:
		return v.VisitAndNode(val)
	case *TableAliasNode:
		return v.VisitTableAliasNode(val)
	case *InnerJoinNode:
		return v.VisitInnerJoinNode(val)
	case *GroupingNode:
		return v.VisitGroupingNode(val)
	case *NamedWindowNode:
		return v.VisitNamedWindowNode(val)
	case *WindowNode:
		return v.VisitWindowNode(val)
	case *RowsNode:
		return v.VisitRowsNode(val)
	case *PrecedingNode:
		return v.VisitPrecedingNode(val)
	case *FollowingNode:
		return v.VisitFollowingNode(val)
	case *CurrentRowNode:
		return v.VisitCurrentRowNode(val)
	case *BetweenNode:
		return v.VisitBetweenNode(val)
	case *RangeNode:
		return v.VisitRangeNode(val)
	case *DistinctNode:
		return v.VisitDistinctNode(val)
	case *WithNode:
		return v.VisitWithNode(val)
	case *WithRecursiveNode:
		return v.VisitWithRecursiveNode(val)
	case *Table:
		if val == nil {
			return v.VisitNil()
		}
		return v.VisitTable(val)
	case *MultiStatementManager:
		return v.VisitMultiStatementManager(val)
	case *InsertStatementNode:
		return v.VisitInsertStatementNode(val)
	case *ValuesNode:
		return v.VisitValuesNode(val)
	case *SelectCoreNode:
		return v.VisitSelectCoreNode(val)
	case *NotEqualNode:
		return v.VisitNotEqualNode(val)
	case *NotNode:
		return v.VisitNotNode(val)
	case *GreaterThanOrEqualNode:
		return v.VisitGreaterThanOrEqualNode(val)
	case *LessThanOrEqualNode:
		return v.VisitLessThanOrEqualNode(val)
	case *OrNode:
		return v.VisitOrNode(val)
	case *AvgNode:
		return v.VisitAvgNode(val)
	case *NamedFunctionNode:
		return v.VisitNamedFunctionNode(val)
	case *SumNode:
		return v.VisitSumNode(val)
	case *MinNode:
		return v.VisitMinNode(val)
	case *MaxNode:
		return v.VisitMaxNode(val)
	case *MatchesNode:
		return v.VisitMatchesNode(val)
	case *DoesNotMatchNode:
		return v.VisitDoesNotMatchNode(val)
	case *NotInNode:
		return v.VisitNotInNode(val)
	case *BinNode:
		return v.VisitBinNode(val)
	case *ExtractNode:
		return v.VisitExtractNode(val)
	case *InfixOperationNode:
		return v.VisitInfixOperationNode(val)
	case *QuotedNode:
		return v.VisitQuotedNode(val)
	case *OverNode:
		return v.VisitOverNode(val)
	case *AssignmentNode:
		return v.VisitAssignmentNode(val)
	case *UnqualifiedColumnNode:
		return v.VisitUnqualifiedColumnNode(val)
	default:
		debug.PrintStack()
		log.Fatalf("ToSqlVisitor#Visit unable to handle type %T", visitable)
		return ""
	}
}

func (v ToSqlVisitor) VisitTopNode(node TopNode) string {
	log.Fatal("NOT IMPLEMENTED FOR THIS DB")
	return ""
}

func (v ToSqlVisitor) VisitOrderingNode(node *OrderingNode) string {
	log.Fatal("NOT IMPLEMENTED")
	return ""
}

func (v ToSqlVisitor) VisitUnqualifiedColumnNode(node *UnqualifiedColumnNode) string {
	return v.QuoteColumnName(node.Name())
}

func (v ToSqlVisitor) VisitAssignmentNode(node *AssignmentNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" = ")
	buf.WriteString(v.Quote(node.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitOverNode(node *OverNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" OVER ")
	if node.Right == nil {
		buf.WriteString("()")
	} else {
		buf.WriteString(v.Visit(node.Right))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitQuotedNode(node *QuotedNode) string {
	return strings.Join([]string{"'", node.Raw, "'"}, "")
}

func (v ToSqlVisitor) VisitInfixOperationNode(node *InfixOperationNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(SPACE)
	buf.WriteString(v.Visit(node.Operator))
	buf.WriteString(SPACE)
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitExtractNode(node *ExtractNode) string {
	var buf bytes.Buffer
	buf.WriteString("EXTRACT(")
	buf.WriteString(strings.ToUpper(node.Field.Raw))
	buf.WriteString(" FROM ")
	for i, expression := range node.Expressions {
		buf.WriteString(v.Visit(expression))
		if i != len(node.Expressions)-1 {
			buf.WriteString(COMMA)
		}
	}
	buf.WriteString(")")
	if node.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(node.Alias))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitBinNode(node *BinNode) string {
	return v.Visit(node.Expr)
}

func (v ToSqlVisitor) VisitNotNode(node *NotNode) string {
	var buf bytes.Buffer
	buf.WriteString("NOT ")
	buf.WriteString(v.Visit(node.Expr))
	return buf.String()
}

func (v ToSqlVisitor) VisitNotInNode(node *NotInNode) string {
	if len(node.Right) == 0 {
		return "1=1"
	}
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" NOT IN (")
	for i, expr := range node.Right {
		buf.WriteString(v.Visit(expr))
		// Join on ", "
		if i != len(node.Right)-1 {
			buf.WriteString(COMMA)
		}
	}
	buf.WriteString(")")
	return buf.String()
}

func (v ToSqlVisitor) VisitInNode(node *InNode) string {
	if len(node.Right) == 0 {
		return "1=0"
	}
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" IN (")
	for i, expr := range node.Right {
		buf.WriteString(v.Visit(expr))
		// Join on ", "
		if i != len(node.Right)-1 {
			buf.WriteString(COMMA)
		}
	}
	buf.WriteString(")")
	return buf.String()
}

func (v ToSqlVisitor) VisitDoesNotMatchNode(node *DoesNotMatchNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" NOT LIKE ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitMatchesNode(node *MatchesNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" LIKE ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitNamedFunctionNode(node *NamedFunctionNode) string {
	var buf bytes.Buffer
	buf.WriteString(node.Name.Raw)
	buf.WriteString("(")
	if node.Distinct {
		buf.WriteString("DISINCT ")
	}
	for i, expr := range node.Expressions {
		buf.WriteString(v.Visit(expr))
		// Join on ", "
		if i != len(node.Expressions)-1 {
			buf.WriteString(COMMA)
		}
	}
	buf.WriteString(")")

	if node.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(node.Alias))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitSumNode(node *SumNode) string {
	var buf bytes.Buffer
	buf.WriteString("SUM(")
	if node.Distinct {
		buf.WriteString("DISINCT ")
	}
	for i, expr := range node.Expressions {
		buf.WriteString(v.Visit(expr))
		// Join on ", "
		if i != len(node.Expressions)-1 {
			buf.WriteString(COMMA)
		}
	}
	buf.WriteString(")")

	if node.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(node.Alias))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitAvgNode(node *AvgNode) string {
	var buf bytes.Buffer
	buf.WriteString("AVG(")
	if node.Distinct {
		buf.WriteString("DISINCT ")
	}
	for i, expr := range node.Expressions {
		buf.WriteString(v.Visit(expr))
		// Join on ", "
		if i != len(node.Expressions)-1 {
			buf.WriteString(COMMA)
		}
	}
	buf.WriteString(")")

	if node.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(node.Alias))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitMinNode(node *MinNode) string {
	var buf bytes.Buffer
	buf.WriteString("MIN(")
	if node.Distinct {
		buf.WriteString("DISINCT ")
	}
	for i, expr := range node.Expressions {
		buf.WriteString(v.Visit(expr))
		// Join on ", "
		if i != len(node.Expressions)-1 {
			buf.WriteString(COMMA)
		}
	}
	buf.WriteString(")")

	if node.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(node.Alias))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitMaxNode(node *MaxNode) string {
	var buf bytes.Buffer
	buf.WriteString("MAX(")
	if node.Distinct {
		buf.WriteString("DISINCT ")
	}
	for i, expr := range node.Expressions {
		buf.WriteString(v.Visit(expr))
		// Join on ", "
		if i != len(node.Expressions)-1 {
			buf.WriteString(COMMA)
		}
	}
	buf.WriteString(")")

	if node.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(node.Alias))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitOrNode(node *OrNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" OR ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitGreaterThanOrEqualNode(node *GreaterThanOrEqualNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" >= ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitLessThanOrEqualNode(node *LessThanOrEqualNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" <= ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitNotEqualNode(node *NotEqualNode) string {
	var buf bytes.Buffer
	if node.Right == nil {
		buf.WriteString(v.Visit(node.Left))
		buf.WriteString(" IS NOT NULL")
	} else {
		buf.WriteString(v.Visit(node.Left))
		buf.WriteString(" != ")
		buf.WriteString(v.Visit(node.Right))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitValuesNode(node *ValuesNode) string {
	var buf bytes.Buffer
	buf.WriteString("VALUES (")
	for i, value := range node.Values {
		buf.WriteString(v.Quote(value))
		// Join on ", "
		if i != len(node.Values)-1 {
			buf.WriteString(COMMA)
		}
	}
	buf.WriteString(")")
	return buf.String()
}

func (v ToSqlVisitor) VisitTrueNode(node TrueNode) string {
	return "TRUE"
}

func (v ToSqlVisitor) VisitFalseNode(node FalseNode) string {
	return "FALSE"
}

func (v ToSqlVisitor) VisitDeleteStatementNode(node DeleteStatementNode) string {
	var buf bytes.Buffer

	buf.WriteString("DELETE FROM ")
	buf.WriteString(v.Visit(node.Relation))

	if node.Wheres != nil {
		buf.WriteString(WHERE)
		for i, where := range *node.Wheres {
			buf.WriteString(v.Visit(where))
			// Join on " AND "
			if i != len(*node.Wheres)-1 {
				buf.WriteString(AND)
			}
		}
	}

	return buf.String()
}

func (v ToSqlVisitor) VisitUpdateStatementNode(node UpdateStatementNode) string {
	var buf bytes.Buffer

	var wheres []Visitable

	if node.Orders == nil && node.Limit == nil {
		wheres = *node.Wheres
	} else {
		stmt := NewSelectStatementNode()
		core := stmt.Cores[0]
		core.SetFrom(node.Relation)
		core.Wheres = &wheres

		if core.Projections == nil {
			slice := make([]Visitable, 0)
			core.Projections = &slice
		}
		*core.Projections = append(*core.Projections, node.Key)
		stmt.Limit = node.Limit
		stmt.Orders = node.Orders

		wheres = append(wheres, &InNode{
			Left:  node.Key,
			Right: []Visitable{stmt},
		})
	}
	buf.WriteString("UPDATE ")
	buf.WriteString(v.Visit(node.Relation))

	if node.Values != nil {
		buf.WriteString("VALUES ")
		for i, value := range *node.Values {
			buf.WriteString(v.Visit(value))
			// Join on ", "
			if i != len(*node.Values)-1 {
				buf.WriteString(COMMA)
			}
		}
	}

	if node.Wheres != nil {
		buf.WriteString("WHERE ")
		for i, where := range *node.Wheres {
			buf.WriteString(v.Visit(where))
			// Join on " AND "
			if i != len(*node.Wheres)-1 {
				buf.WriteString(AND)
			}
		}
	}

	return buf.String()
}

func (v ToSqlVisitor) VisitInsertStatementNode(node *InsertStatementNode) string {
	var buf bytes.Buffer
	buf.WriteString("INSERT INTO ")
	buf.WriteString(v.Visit(node.Relation))

	if node.Columns != nil && len(*node.Columns) > 0 {
		buf.WriteString(" (")
		for i, column := range *node.Columns {
			buf.WriteString(v.QuoteColumnName(column.Name))
			// Join on ", "
			if i != len(*node.Columns)-1 {
				buf.WriteString(COMMA)
			}
		}
		buf.WriteString(")")
	}

	if node.Values != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.Visit(node.Values))
	}

	return buf.String()
}

func (v ToSqlVisitor) VisitNil() string {
	return "NULL"
}

func (v ToSqlVisitor) VisitWithNode(node *WithNode) string {
	var buf bytes.Buffer
	buf.WriteString("WITH ")
	buf.WriteString(v.Visit(node.Expr))
	return buf.String()
}

func (v ToSqlVisitor) VisitWithRecursiveNode(node *WithRecursiveNode) string {
	var buf bytes.Buffer
	buf.WriteString("WITH RECURSIVE ")
	buf.WriteString(v.Visit(node.Expr))
	return buf.String()
}

func (v ToSqlVisitor) VisitDistinctOnNode(node DistinctOnNode) string {
	log.Fatal("NOT IMPLEMENTED FOR THIS DB")
	return ""
}

func (v ToSqlVisitor) VisitDistinctNode(node *DistinctNode) string {
	return DISTINCT
}

func (v ToSqlVisitor) VisitRangeNode(node *RangeNode) string {
	var buf bytes.Buffer
	buf.WriteString("RANGE")
	if node.Expr != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.Visit(node.Expr))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitBetweenNode(node *BetweenNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" BETWEEN ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitCurrentRowNode(node *CurrentRowNode) string {
	return "CURRENT ROW"
}

func (v ToSqlVisitor) VisitPrecedingNode(node *PrecedingNode) string {
	var buf bytes.Buffer
	if node.Expr != nil {
		buf.WriteString(v.Visit(node.Expr))
	} else {
		buf.WriteString("UNBOUNDED")
	}
	buf.WriteString(" PRECEDING")
	return buf.String()
}

func (v ToSqlVisitor) VisitFollowingNode(node *FollowingNode) string {
	var buf bytes.Buffer
	if node.Expr != nil {
		buf.WriteString(v.Visit(node.Expr))
	} else {
		buf.WriteString("UNBOUNDED")
	}
	buf.WriteString(" FOLLOWING")
	return buf.String()
}

func (v ToSqlVisitor) VisitRowsNode(node *RowsNode) string {
	var buf bytes.Buffer
	buf.WriteString("ROWS")
	if node.Expr != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.Visit(node.Expr))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitNamedWindowNode(node *NamedWindowNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.QuoteColumnName(node.Name))
	buf.WriteString(" AS (")
	visitOrders := (node.Orders != nil && len(*node.Orders) > 0)
	visitFraming := (node.Framing != nil)
	if visitOrders {
		buf.WriteString("ORDER BY ")
		for i, order := range *node.Orders {
			buf.WriteString(v.Visit(order))
			// Join on ", "
			if i != len(*node.Orders)-1 {
				buf.WriteString(COMMA)
			}
		}
	}

	if visitOrders && visitFraming {
		buf.WriteString(SPACE)
	}

	if visitFraming {
		buf.WriteString(v.Visit(node.Framing))
	}

	buf.WriteString(")")
	return buf.String()
}

func (v ToSqlVisitor) VisitWindowNode(node *WindowNode) string {
	var buf bytes.Buffer
	buf.WriteString("(")
	visitOrders := (node.Orders != nil && len(*node.Orders) > 0)
	visitFraming := (node.Framing != nil)
	if visitOrders {
		buf.WriteString("ORDER BY ")
		for i, order := range *node.Orders {
			buf.WriteString(v.Visit(order))
			// Join on ", "
			if i != len(*node.Orders)-1 {
				buf.WriteString(COMMA)
			}
		}
	}

	if visitOrders && visitFraming {
		buf.WriteString(SPACE)
	}

	if visitFraming {
		buf.WriteString(v.Visit(node.Framing))
	}

	buf.WriteString(")")
	return buf.String()
}

func (v ToSqlVisitor) VisitGroupingNode(node *GroupingNode) string {
	var buf bytes.Buffer
	buf.WriteString("(")
	for _, expr := range node.Expr {
		if expr != nil {
			buf.WriteString(v.Visit(expr))
		}
	}
	buf.WriteString(")")
	return buf.String()
}

func (v ToSqlVisitor) VisitLimitNode(node LimitNode) string {
	var buf bytes.Buffer
	buf.WriteString("LIMIT ")
	buf.WriteString(v.Visit(node.Expr))
	return buf.String()
}

func (v ToSqlVisitor) VisitLockNode(node LockNode) string {
	return v.Visit(node.Expr)
}

func (v ToSqlVisitor) VisitOffsetNode(node OffsetNode) string {
	var buf bytes.Buffer
	buf.WriteString("OFFSET ")
	buf.WriteString(v.Visit(node.Expr))
	return buf.String()
}

func (v ToSqlVisitor) VisitAndNode(node *AndNode) string {
	var buf bytes.Buffer
	if node.Children != nil {
		children := *node.Children
		for i, child := range children {
			buf.WriteString(v.Visit(child))
			if i != len(children)-1 {
				buf.WriteString(" AND ")
			}
		}
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitCountNode(node *CountNode) string {
	var buf bytes.Buffer
	buf.WriteString("COUNT(")
	if node.Distinct {
		buf.WriteString("DISINCT ")
	}
	for i, expr := range node.Expressions {
		buf.WriteString(v.Visit(expr))
		// Join on ", "
		if i != len(node.Expressions)-1 {
			buf.WriteString(COMMA)
		}
	}
	buf.WriteString(")")

	if node.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(node.Alias))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitAscendingNode(node *AscendingNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Expr))
	buf.WriteString(" ASC")
	return buf.String()
}

func (v ToSqlVisitor) VisitDescendingNode(node *DescendingNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Expr))
	buf.WriteString(" DESC")
	return buf.String()
}

func (v ToSqlVisitor) VisitOnNode(node *OnNode) string {
	var buf bytes.Buffer
	buf.WriteString("ON ")
	buf.WriteString(v.Visit(node.Expr))
	return buf.String()
}

func (v ToSqlVisitor) VisitExceptNode(node *ExceptNode) string {
	var buf bytes.Buffer
	buf.WriteString("( ")
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" EXCEPT ")
	buf.WriteString(v.Visit(node.Right))
	buf.WriteString(" )")
	return buf.String()
}

func (v ToSqlVisitor) VisitIntersectNode(node *IntersectNode) string {
	var buf bytes.Buffer
	buf.WriteString("( ")
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" INTERSECT ")
	buf.WriteString(v.Visit(node.Right))
	buf.WriteString(" )")
	return buf.String()
}

func (v ToSqlVisitor) VisitSelectManager(node *SelectManager) string {
	var buf bytes.Buffer
	buf.WriteString("(")
	buf.WriteString(node.ToSql())
	buf.WriteString(")")
	return buf.String()
}

func (v ToSqlVisitor) VisitMultiStatementManager(node *MultiStatementManager) string {
	var buf bytes.Buffer
	buf.WriteString(node.ToSql())
	return buf.String()
}

func (v ToSqlVisitor) VisitUnionNode(node *UnionNode) string {
	var buf bytes.Buffer
	buf.WriteString("( ")
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" UNION ")
	buf.WriteString(v.Visit(node.Right))
	buf.WriteString(" )")
	return buf.String()
}

func (v ToSqlVisitor) VisitUnionAllNode(node *UnionAllNode) string {
	var buf bytes.Buffer
	buf.WriteString("( ")
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" UNION ALL ")
	buf.WriteString(v.Visit(node.Right))
	buf.WriteString(" )")
	return buf.String()
}

func (v ToSqlVisitor) VisitLessThanNode(node *LessThanNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" < ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitGreaterThanNode(node *GreaterThanNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" > ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitAsNode(node *AsNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" AS ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitGroupNode(node *GroupNode) string {
	return v.Visit(node.Expr)
}

func (v ToSqlVisitor) VisitHavingNode(node *HavingNode) string {
	var buf bytes.Buffer
	buf.WriteString("HAVING ")
	buf.WriteString(v.Visit(node.Expr))
	return buf.String()
}

func (v ToSqlVisitor) VisitExistsNode(node *ExistsNode) string {
	var buf bytes.Buffer
	buf.WriteString("EXISTS (")
	for i, expr := range node.Expressions {
		buf.WriteString(v.Visit(expr))
		// Join on ", "
		if i != len(node.Expressions)-1 {
			buf.WriteString(COMMA)
		}
	}
	buf.WriteString(")")

	if node.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(node.Alias))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitAttributeNode(node *AttributeNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.QuoteTableName(node.Relation))
	buf.WriteString(".")
	buf.WriteString(v.QuoteColumnName(node.Name))
	return buf.String()
}

func (v ToSqlVisitor) VisitEqualityNode(node *EqualityNode) string {
	var buf bytes.Buffer
	if node.Right == nil {
		buf.WriteString(v.Visit(node.Left))
		buf.WriteString(" IS NULL")
	} else {
		buf.WriteString(v.Visit(node.Left))
		buf.WriteString(" = ")
		buf.WriteString(v.Visit(node.Right))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitTable(table *Table) string {
	var buf bytes.Buffer
	buf.WriteString(v.QuoteTableName(table))
	if table.TableAlias != "" {
		buf.WriteString(SPACE)
		buf.WriteString(v.QuoteTableName(&TableAliasNode{
			Relation: table,
			Name:     Sql(table.TableAlias),
			Quoted:   true,
		}))
	}
	return buf.String()
}

func (v ToSqlVisitor) QuoteTableName(visitable Visitable) string {
	if alias, ok := visitable.(*TableAliasNode); ok {
		if !alias.Quoted {
			return alias.Name.Raw
		}
	}
	return v.Conn.QuoteTableName(visitable.String())
}

func (v ToSqlVisitor) Quote(thing interface{}) string {
	return v.Conn.Quote(thing)
}

func (v ToSqlVisitor) QuoteColumnName(literal SqlLiteralNode) string {
	return v.Conn.QuoteColumnName(literal.Raw)
}

func (v ToSqlVisitor) VisitJoinSourceNode(node JoinSource) string {
	var buf bytes.Buffer
	if node.Left != nil {
		buf.WriteString(v.Visit(node.Left))
	}
	for i, source := range node.Right {
		buf.WriteString(v.Visit(source))
		if i != len(node.Right)-1 {
			buf.WriteString(SPACE)
		}
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitOuterJoinNode(node OuterJoinNode) string {
	var buf bytes.Buffer
	buf.WriteString("LEFT OUTER JOIN ")
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(SPACE)
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitInnerJoinNode(node *InnerJoinNode) string {
	var buf bytes.Buffer
	buf.WriteString(" INNER JOIN ")
	buf.WriteString(v.Visit(node.Left))
	if node.Right != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.Visit(node.Right))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitSqlLiteralNode(node SqlLiteralNode) string {
	if node.Raw != "" {
		return node.Raw
	} else {
		return ""
	}
}

func (v ToSqlVisitor) VisitTableAliasNode(node *TableAliasNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Relation))
	buf.WriteString(" ")
	buf.WriteString(v.QuoteTableName(node))
	return buf.String()
}

func (v ToSqlVisitor) VisitSelectCoreNode(node *SelectCoreNode) string {
	var buf bytes.Buffer

	buf.WriteString("SELECT")

	// Add TOP statement to the buffer
	if node.Top != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.VisitTopNode(*node.Top))
	}

	if node.SetQuanifier != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.Visit(node.SetQuanifier))
	}

	// add select projections
	if node.Projections != nil && len(*node.Projections) > 0 {
		buf.WriteString(SPACE)
		for i, projection := range *node.Projections {
			if projection != nil {
				buf.WriteString(v.Visit(projection))
				// Join on ", "
				if i != len(*node.Projections)-1 {
					buf.WriteString(COMMA)
				}
			}
		}
	}

	// add FROM statement to the buffer
	if node.Source != nil && node.Source.Left != nil {
		// assert the source is a *Table and check the length of the name
		if t, ok := node.Source.Left.(*Table); ok && t.Name != "" {
			buf.WriteString(" FROM ")
			buf.WriteString(v.Visit(*node.Source))
		}
	}

	// add WHERE statement to the buffer
	if node.Wheres != nil && len(*node.Wheres) > 0 {
		buf.WriteString(WHERE)
		for i, where := range *node.Wheres {
			buf.WriteString(v.Visit(where))
			// Join on ", "
			if i != len(*node.Wheres)-1 {
				buf.WriteString(COMMA)
			}
		}
	}

	// add GROUP BY statement to the buffer
	if node.Groups != nil && len(*node.Groups) > 0 {
		buf.WriteString(GROUP_BY)
		for i, group := range *node.Groups {
			buf.WriteString(v.Visit(group))
			// Join on ", "
			if i != len(*node.Groups)-1 {
				buf.WriteString(COMMA)
			}
		}
	}

	// add HAVING statement to the buffer
	if node.Having != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.VisitHavingNode(node.Having))
	}

	// add WINDOW statements to the buffer
	if node.Windows != nil && len(*node.Windows) > 0 {
		buf.WriteString(WINDOW)
		for i, window := range *node.Windows {
			buf.WriteString(v.Visit(window))
			// Join on ", "
			if i != len(*node.Windows)-1 {
				buf.WriteString(COMMA)
			}
		}
	}

	return buf.String()
}

func (v ToSqlVisitor) VisitSelectStatementNode(node *SelectStatementNode) string {
	var buf bytes.Buffer

	// add WITH clause to the buffer
	if node.With != nil {
		buf.WriteString(v.Visit(node.With))
		buf.WriteString(SPACE)
	}

	// add core SELECT clause to the buffer
	if node.Cores != nil {
		for _, core := range node.Cores {
			if core != nil {
				buf.WriteString(v.Visit(core))
			}
		}
	}

	// add ORDER BY clauses to the buffer
	if node.Orders != nil {
		buf.WriteString(ORDER_BY)
		for i, order := range *node.Orders {
			buf.WriteString(v.Visit(order))
			if (len(*node.Orders) - 1) != i {
				buf.WriteString(COMMA)
			}
		}
	}

	// add LIMIT clause to the buffer
	if node.Limit != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.VisitLimitNode(*node.Limit))
	}

	// add OFFSET clause to the buffer
	if node.Offset != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.VisitOffsetNode(*node.Offset))
	}

	// add LOCK clause to the buffer
	if node.Lock != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.VisitLockNode(*node.Lock))
	}

	return strings.TrimSpace(buf.String())
}
