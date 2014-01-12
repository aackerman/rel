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

func (v ToSqlVisitor) Accept(a Visitable) string {
	return v.Visit(a)
}

// FIXME: Only visit pointers to visitables
func (v ToSqlVisitor) Visit(a Visitable) string {
	ret := ""
	switch val := a.(type) {
	case nil:
		ret = v.VisitNil()
	case SelectStatementNode:
		ret = v.VisitSelectStatementNode(val)
	case *InNode:
		ret = v.VisitInNode(*val)
	case SqlLiteralNode:
		ret = v.VisitSqlLiteralNode(val)
	case *SqlLiteralNode:
		ret = v.VisitSqlLiteralNode(*val)
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
	case *ExistsNode:
		ret = v.VisitExistsNode(*val)
	case *AsNode:
		ret = v.VisitAsNode(*val)
	case Table:
		ret = v.VisitTable(val)
	case *LessThanNode:
		ret = v.VisitLessThanNode(*val)
	case UnionNode:
		ret = v.VisitUnionNode(val)
	case UnionAllNode:
		ret = v.VisitUnionAllNode(val)
	case SelectManager:
		ret = v.VisitSelectManager(val)
	case *GreaterThanNode:
		ret = v.VisitGreaterThanNode(*val)
	case IntersectNode:
		ret = v.VisitIntersectNode(val)
	case ExceptNode:
		ret = v.VisitExceptNode(val)
	case TableAliasNode:
		ret = v.VisitTableAliasNode(val)
	case InnerJoinNode:
		ret = v.VisitInnerJoinNode(val)
	case OnNode:
		ret = v.VisitOnNode(val)
	case *AscendingNode:
		ret = v.VisitAscendingNode(*val)
	case *DescendingNode:
		ret = v.VisitDescendingNode(*val)
	case *CountNode:
		ret = v.VisitCountNode(*val)
	case *AndNode:
		ret = v.VisitAndNode(*val)
	case *TableAliasNode:
		ret = v.VisitTableAliasNode(*val)
	case *InnerJoinNode:
		ret = v.VisitInnerJoinNode(*val)
	case *GroupingNode:
		ret = v.VisitGroupingNode(*val)
	case *NamedWindowNode:
		ret = v.VisitNamedWindowNode(*val)
	case *WindowNode:
		ret = v.VisitWindowNode(*val)
	case *RowsNode:
		ret = v.VisitRowsNode(*val)
	case *PrecedingNode:
		ret = v.VisitPrecedingNode(*val)
	case *FollowingNode:
		ret = v.VisitFollowingNode(*val)
	case *CurrentRowNode:
		ret = v.VisitCurrentRowNode(*val)
	case *BetweenNode:
		ret = v.VisitBetweenNode(*val)
	case *RangeNode:
		ret = v.VisitRangeNode(*val)
	case *DistinctNode:
		ret = v.VisitDistinctNode(*val)
	case *WithNode:
		ret = v.VisitWithNode(*val)
	case *WithRecursiveNode:
		ret = v.VisitWithRecursiveNode(*val)
	case *Table:
		if val == nil {
			return v.VisitNil()
		}
		ret = v.VisitTable(*val)
	case *MultiStatementManager:
		ret = v.VisitMultiStatementManager(*val)
	case *InsertStatementNode:
		ret = v.VisitInsertStatementNode(*val)
	case *ValuesNode:
		ret = v.VisitValuesNode(*val)
	case *SelectCoreNode:
		ret = v.VisitSelectCoreNode(*val)
	case *NotEqualNode:
		ret = v.VisitNotEqualNode(*val)
	case *NotNode:
		ret = v.VisitNotNode(*val)
	case *GreaterThanOrEqualNode:
		ret = v.VisitGreaterThanOrEqualNode(*val)
	case *LessThanOrEqualNode:
		ret = v.VisitLessThanOrEqualNode(*val)
	case *OrNode:
		ret = v.VisitOrNode(*val)
	case *AvgNode:
		ret = v.VisitAvgNode(*val)
	case *NamedFunctionNode:
		ret = v.VisitNamedFunctionNode(*val)
	case *SumNode:
		ret = v.VisitSumNode(*val)
	case *MinNode:
		ret = v.VisitMinNode(*val)
	case *MaxNode:
		ret = v.VisitMaxNode(*val)
	case *MatchesNode:
		ret = v.VisitMatchesNode(*val)
	case *DoesNotMatchNode:
		ret = v.VisitDoesNotMatchNode(*val)
	case *NotInNode:
		ret = v.VisitNotInNode(*val)
	case *BinNode:
		ret = v.VisitBinNode(*val)
	case *ExtractNode:
		ret = v.VisitExtractNode(*val)
	case *InfixOperationNode:
		ret = v.VisitInfixOperationNode(*val)
	default:
		debug.PrintStack()
		log.Fatalf("ToSqlVisitor#Visit unable to handle type %T", a)
	}
	return ret
}

func (v ToSqlVisitor) VisitTopNode(node TopNode) string {
	log.Fatal("NOT IMPLEMENTED FOR THIS DB")
	return ""
}

func (v ToSqlVisitor) VisitOrderingNode(node OrderingNode) string {
	log.Fatal("NOT IMPLEMENTED")
	return ""
}

func (v ToSqlVisitor) VisitInfixOperationNode(node InfixOperationNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(SPACE)
	buf.WriteString(v.Visit(node.Operator))
	buf.WriteString(SPACE)
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitExtractNode(node ExtractNode) string {
	var buf bytes.Buffer
	buf.WriteString("EXTRACT(")
	buf.WriteString(strings.ToUpper(node.Field.Raw))
	buf.WriteString(" FROM ")
	buf.WriteString(v.Visit(node.Expr))
	buf.WriteString(")")
	if node.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(node.Alias))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitBinNode(node BinNode) string {
	return v.Visit(node.Expr)
}

func (v ToSqlVisitor) VisitNotNode(node NotNode) string {
	var buf bytes.Buffer
	buf.WriteString("NOT ")
	buf.WriteString(v.Visit(node.Expr))
	return buf.String()
}

func (v ToSqlVisitor) VisitNotInNode(node NotInNode) string {
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
			buf.WriteString(", ")
		}
	}
	buf.WriteString(")")
	return buf.String()
}

func (v ToSqlVisitor) VisitInNode(node InNode) string {
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
			buf.WriteString(", ")
		}
	}
	buf.WriteString(")")
	return buf.String()
}

func (v ToSqlVisitor) VisitDoesNotMatchNode(node DoesNotMatchNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" NOT LIKE ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitMatchesNode(node MatchesNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" LIKE ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitNamedFunctionNode(node NamedFunctionNode) string {
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
			buf.WriteString(", ")
		}
	}
	buf.WriteString(")")

	if node.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(node.Alias))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitSumNode(node SumNode) string {
	var buf bytes.Buffer
	buf.WriteString("SUM(")
	if node.Distinct {
		buf.WriteString("DISINCT ")
	}
	for i, expr := range node.Expressions {
		buf.WriteString(v.Visit(expr))
		// Join on ", "
		if i != len(node.Expressions)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteString(")")

	if node.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(node.Alias))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitAvgNode(node AvgNode) string {
	var buf bytes.Buffer
	buf.WriteString("AVG(")
	if node.Distinct {
		buf.WriteString("DISINCT ")
	}
	for i, expr := range node.Expressions {
		buf.WriteString(v.Visit(expr))
		// Join on ", "
		if i != len(node.Expressions)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteString(")")

	if node.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(node.Alias))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitMinNode(node MinNode) string {
	var buf bytes.Buffer
	buf.WriteString("MIN(")
	if node.Distinct {
		buf.WriteString("DISINCT ")
	}
	for i, expr := range node.Expressions {
		buf.WriteString(v.Visit(expr))
		// Join on ", "
		if i != len(node.Expressions)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteString(")")

	if node.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(node.Alias))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitMaxNode(node MaxNode) string {
	var buf bytes.Buffer
	buf.WriteString("MAX(")
	if node.Distinct {
		buf.WriteString("DISINCT ")
	}
	for i, expr := range node.Expressions {
		buf.WriteString(v.Visit(expr))
		// Join on ", "
		if i != len(node.Expressions)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteString(")")

	if node.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(node.Alias))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitOrNode(node OrNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" OR ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitGreaterThanOrEqualNode(node GreaterThanOrEqualNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" >= ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitLessThanOrEqualNode(node LessThanOrEqualNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" <= ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitNotEqualNode(node NotEqualNode) string {
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

// FIXME: find a better way to handle insert values
func (v ToSqlVisitor) VisitInsertValue(i interface{}) string {
	ret := ""
	switch val := i.(type) {
	case bool:
		if val {
			ret = "'t'"
		} else {
			ret = "'f'"
		}
	case SqlLiteralNode:
		ret = val.Raw
	}

	return ret
}

func (v ToSqlVisitor) VisitValuesNode(node ValuesNode) string {
	var buf bytes.Buffer
	buf.WriteString("VALUES (")
	for i, value := range node.Values {
		buf.WriteString(v.VisitInsertValue(value))
		// Join on ", "
		if i != len(node.Values)-1 {
			buf.WriteString(", ")
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
				buf.WriteString(", ")
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

func (v ToSqlVisitor) VisitInsertStatementNode(node InsertStatementNode) string {
	var buf bytes.Buffer
	buf.WriteString("INSERT INTO ")
	buf.WriteString(v.Visit(node.Relation))

	if node.Columns != nil && len(*node.Columns) > 0 {
		buf.WriteString(" (")
		for i, column := range *node.Columns {
			buf.WriteString(v.QuoteColumnName(column.Name))
			// Join on ", "
			if i != len(*node.Columns)-1 {
				buf.WriteString(", ")
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

func (v ToSqlVisitor) VisitWithNode(node WithNode) string {
	var buf bytes.Buffer
	buf.WriteString("WITH ")
	buf.WriteString(v.Visit(node.Expr))
	return buf.String()
}

func (v ToSqlVisitor) VisitWithRecursiveNode(node WithRecursiveNode) string {
	var buf bytes.Buffer
	buf.WriteString("WITH RECURSIVE ")
	buf.WriteString(v.Visit(node.Expr))
	return buf.String()
}

func (v ToSqlVisitor) VisitDistinctOnNode(node DistinctOnNode) string {
	log.Fatal("NOT IMPLEMENTED FOR THIS DB")
	return ""
}

func (v ToSqlVisitor) VisitDistinctNode(node DistinctNode) string {
	return DISTINCT
}

func (v ToSqlVisitor) VisitRangeNode(node RangeNode) string {
	var buf bytes.Buffer
	buf.WriteString("RANGE")
	if node.Expr != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.Visit(node.Expr))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitBetweenNode(node BetweenNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" BETWEEN ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitCurrentRowNode(node CurrentRowNode) string {
	return "CURRENT ROW"
}

func (v ToSqlVisitor) VisitPrecedingNode(node PrecedingNode) string {
	var buf bytes.Buffer
	if node.Expr != nil {
		buf.WriteString(v.Visit(node.Expr))
	} else {
		buf.WriteString("UNBOUNDED")
	}
	buf.WriteString(" PRECEDING")
	return buf.String()
}

func (v ToSqlVisitor) VisitFollowingNode(node FollowingNode) string {
	var buf bytes.Buffer
	if node.Expr != nil {
		buf.WriteString(v.Visit(node.Expr))
	} else {
		buf.WriteString("UNBOUNDED")
	}
	buf.WriteString(" FOLLOWING")
	return buf.String()
}

func (v ToSqlVisitor) VisitRowsNode(node RowsNode) string {
	var buf bytes.Buffer
	buf.WriteString("ROWS")
	if node.Expr != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.Visit(node.Expr))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitNamedWindowNode(node NamedWindowNode) string {
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
				buf.WriteString(", ")
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

func (v ToSqlVisitor) VisitWindowNode(node WindowNode) string {
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
				buf.WriteString(", ")
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

func (v ToSqlVisitor) VisitGroupingNode(node GroupingNode) string {
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

func (v ToSqlVisitor) VisitAndNode(node AndNode) string {
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

func (v ToSqlVisitor) VisitCountNode(node CountNode) string {
	var buf bytes.Buffer
	buf.WriteString("COUNT(")
	if node.Distinct {
		buf.WriteString("DISINCT ")
	}
	for i, expr := range node.Expressions {
		buf.WriteString(v.Visit(expr))
		// Join on ", "
		if i != len(node.Expressions)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteString(")")

	if node.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(node.Alias))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitAscendingNode(node AscendingNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Expr))
	buf.WriteString(" ASC")
	return buf.String()
}

func (v ToSqlVisitor) VisitDescendingNode(node DescendingNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Expr))
	buf.WriteString(" DESC")
	return buf.String()
}

func (v ToSqlVisitor) VisitOnNode(node OnNode) string {
	var buf bytes.Buffer
	buf.WriteString("ON ")
	buf.WriteString(v.Visit(node.Expr))
	return buf.String()
}

func (v ToSqlVisitor) VisitExceptNode(node ExceptNode) string {
	var buf bytes.Buffer
	buf.WriteString("( ")
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" EXCEPT ")
	buf.WriteString(v.Visit(node.Right))
	buf.WriteString(" )")
	return buf.String()
}

func (v ToSqlVisitor) VisitIntersectNode(node IntersectNode) string {
	var buf bytes.Buffer
	buf.WriteString("( ")
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" INTERSECT ")
	buf.WriteString(v.Visit(node.Right))
	buf.WriteString(" )")
	return buf.String()
}

func (v ToSqlVisitor) VisitSelectManager(node SelectManager) string {
	var buf bytes.Buffer
	buf.WriteString("(")
	buf.WriteString(node.ToSql())
	buf.WriteString(")")
	return buf.String()
}

func (v ToSqlVisitor) VisitMultiStatementManager(node MultiStatementManager) string {
	var buf bytes.Buffer
	buf.WriteString(node.ToSql())
	return buf.String()
}

func (v ToSqlVisitor) VisitUnionNode(node UnionNode) string {
	var buf bytes.Buffer
	buf.WriteString("( ")
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" UNION ")
	buf.WriteString(v.Visit(node.Right))
	buf.WriteString(" )")
	return buf.String()
}

func (v ToSqlVisitor) VisitUnionAllNode(node UnionAllNode) string {
	var buf bytes.Buffer
	buf.WriteString("( ")
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" UNION ALL ")
	buf.WriteString(v.Visit(node.Right))
	buf.WriteString(" )")
	return buf.String()
}

func (v ToSqlVisitor) VisitLessThanNode(node LessThanNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" < ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitGreaterThanNode(node GreaterThanNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" > ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitAsNode(node AsNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" AS ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v ToSqlVisitor) VisitGroupNode(node GroupNode) string {
	return v.Visit(node.Expr)
}

func (v ToSqlVisitor) VisitHavingNode(node HavingNode) string {
	var buf bytes.Buffer
	buf.WriteString("HAVING ")
	buf.WriteString(v.Visit(node.Expr))
	return buf.String()
}

func (v ToSqlVisitor) VisitExistsNode(node ExistsNode) string {
	var buf bytes.Buffer
	buf.WriteString("EXISTS (")
	for i, expr := range node.Expressions {
		buf.WriteString(v.Visit(expr))
		// Join on ", "
		if i != len(node.Expressions)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteString(")")

	if node.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(node.Alias))
	}
	return buf.String()
}

func (v ToSqlVisitor) VisitAttributeNode(node AttributeNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.QuoteTableName(node.Relation))
	buf.WriteString(".")
	buf.WriteString(v.QuoteColumnName(node.Name))
	return buf.String()
}

func (v ToSqlVisitor) VisitEqualityNode(node EqualityNode) string {
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

func (v ToSqlVisitor) VisitTable(table Table) string {
	var buf bytes.Buffer
	if table.TableAlias != "" {
		buf.WriteString(v.QuoteTableName(table))
		buf.WriteString(SPACE)
		// FIXME: table.TableAlias should be a ptr to a TableAliasNode not a string
		alias := TableAliasNode{Relation: table, Name: Sql(table.TableAlias), Quoted: true}
		buf.WriteString(v.QuoteTableName(alias))
	} else {
		buf.WriteString(v.QuoteTableName(table))
	}
	return buf.String()
}

// FIXME: far too complex
func (v ToSqlVisitor) QuoteTableName(visitable Visitable) string {
	switch rel := visitable.(type) {
	case Table:
		return v.conn.QuoteTableName(rel.Name)
	case *Table:
		return v.conn.QuoteTableName(rel.Name)
	case TableAliasNode:
		if rel.Quoted == true {
			return v.conn.QuoteTableName(rel.Name.Raw)
		} else {
			return rel.Name.Raw
		}
	case *TableAliasNode:
		if rel.Quoted == true {
			return v.conn.QuoteTableName(rel.Name.Raw)
		} else {
			return rel.Name.Raw
		}
	case SqlLiteralNode:
		return rel.Raw
	default:
		return ""
	}
}

func (v ToSqlVisitor) QuoteColumnName(literal SqlLiteralNode) string {
	return v.conn.QuoteColumnName(literal.Raw)
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

func (v ToSqlVisitor) VisitInnerJoinNode(node InnerJoinNode) string {
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

func (v ToSqlVisitor) VisitTableAliasNode(node TableAliasNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Relation))
	buf.WriteString(" ")
	buf.WriteString(v.QuoteTableName(node))
	return buf.String()
}

func (v ToSqlVisitor) VisitSelectCoreNode(node SelectCoreNode) string {
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
	if node.Projections != nil {
		claused := false
		for i, projection := range *node.Projections {

			if projection != nil {
				if !claused {
					buf.WriteString(SPACE)
					claused = true
				}
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
		if t, ok := node.Source.Left.(Table); ok && t.Name != "" {
			buf.WriteString(" FROM ")
			buf.WriteString(v.Visit(*node.Source))
		} else if t, ok := node.Source.Left.(*Table); ok && t.Name != "" {
			buf.WriteString(" FROM ")
			buf.WriteString(v.Visit(*node.Source))
		}
	}

	// add WHERE statement to the buffer
	if node.Wheres != nil {
		claused := false
		for i, where := range *node.Wheres {
			// add WHERE clause if it hasn't already been added
			if !claused {
				buf.WriteString(WHERE)
				claused = true
			}
			buf.WriteString(v.Visit(where))

			// Join on ", "
			if i != len(*node.Wheres)-1 {
				buf.WriteString(COMMA)
			}
		}
	}

	// add GROUP BY statement to the buffer
	if node.Groups != nil {
		claused := false
		for i, group := range *node.Groups {
			// add GROUP BY clause if it hasn't already been added
			if !claused {
				buf.WriteString(GROUP_BY)
				claused = true
			}
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
		buf.WriteString(v.VisitHavingNode(*node.Having))
	}

	// add WINDOW statements to the buffer
	if node.Windows != nil {
		claused := false
		for i, window := range *node.Windows {
			// add WINDOW clause if is hasn't already been added
			if !claused {
				buf.WriteString(WINDOW)
				claused = true
			}
			buf.WriteString(v.Visit(window))
			// Join on ", "
			if i != len(*node.Windows)-1 {
				buf.WriteString(COMMA)
			}
		}
	}

	return buf.String()
}

func (v ToSqlVisitor) VisitSelectStatementNode(node SelectStatementNode) string {
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
