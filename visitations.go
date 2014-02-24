package rel

import (
	"bytes"
	"log"
	"strings"
)

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

func visitationTopNode(v Visitor, node *TopNode) string {
	log.Fatal("NOT IMPLEMENTED FOR THIS DB")
	return ""
}

func visitationOrderingNode(v Visitor, node *OrderingNode) string {
	log.Fatal("NOT IMPLEMENTED")
	return ""
}

func visitationUsingNode(v Visitor, node *UsingNode) string {
	var buf bytes.Buffer
	buf.WriteString("USING (")
	buf.WriteString(v.Visit(node.Expr))
	buf.WriteString(")")
	return buf.String()
}

func visitationUnqualifiedColumnNode(v Visitor, node *UnqualifiedColumnNode) string {
	return v.QuoteColumnName(node.Name())
}

func visitationAssignmentNode(v Visitor, node *AssignmentNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" = ")
	buf.WriteString(v.Quote(node.Right))
	return buf.String()
}

func visitationOverNode(v Visitor, node *OverNode) string {
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

func visitationQuotedNode(v Visitor, node *QuotedNode) string {
	return strings.Join([]string{"'", node.Raw, "'"}, "")
}

func visitationInfixOperationNode(v Visitor, node *InfixOperationNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(SPACE)
	buf.WriteString(v.Visit(node.Operator))
	buf.WriteString(SPACE)
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func visitationExtractNode(v Visitor, node *ExtractNode) string {
	var buf bytes.Buffer
	buf.WriteString("EXTRACT(")
	buf.WriteString(strings.ToUpper(node.Field.Raw))
	buf.WriteString(" FROM ")
	buf.WriteString(iterateVisitAndJoinOnComma(v, node.Expressions))
	buf.WriteString(")")
	if node.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(node.Alias))
	}
	return buf.String()
}

func visitationBinNode(v Visitor, node *BinNode) string {
	return v.Visit(node.Expr)
}

func visitationNotNode(v Visitor, node *NotNode) string {
	var buf bytes.Buffer
	buf.WriteString("NOT ")
	buf.WriteString(v.Visit(node.Expr))
	return buf.String()
}

func visitationNotInNode(v Visitor, node *NotInNode) string {
	if len(node.Right) == 0 {
		return "1=1"
	}
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" NOT IN (")
	buf.WriteString(iterateVisitAndJoinOnComma(v, node.Right))
	buf.WriteString(")")
	return buf.String()
}

func visitationInNode(v Visitor, node *InNode) string {
	if len(node.Right) == 0 || node.Right == nil {
		return "1=0"
	}
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" IN (")
	buf.WriteString(iterateVisitAndJoinOnComma(v, node.Right))
	buf.WriteString(")")
	return buf.String()
}

func visitationDoesNotMatchNode(v Visitor, node *DoesNotMatchNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" NOT LIKE ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func visitationMatchesNode(v Visitor, node *MatchesNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" LIKE ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func visitationNamedFunctionNode(v Visitor, node *NamedFunctionNode) string {
	var buf bytes.Buffer
	buf.WriteString(node.Name.Raw)
	buf.WriteString("(")
	if node.Distinct {
		buf.WriteString("DISINCT ")
	}
	buf.WriteString(iterateVisitAndJoinOnComma(v, node.Expressions))
	buf.WriteString(")")

	if node.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(node.Alias))
	}
	return buf.String()
}

func visitationSumNode(v Visitor, node *SumNode) string {
	var buf bytes.Buffer
	buf.WriteString("SUM(")
	if node.Distinct {
		buf.WriteString("DISINCT ")
	}
	buf.WriteString(iterateVisitAndJoinOnComma(v, node.Expressions))
	buf.WriteString(")")

	if node.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(node.Alias))
	}
	return buf.String()
}

func visitationAvgNode(v Visitor, node *AvgNode) string {
	var buf bytes.Buffer
	buf.WriteString("AVG(")
	if node.Distinct {
		buf.WriteString("DISINCT ")
	}
	buf.WriteString(iterateVisitAndJoinOnComma(v, node.Expressions))
	buf.WriteString(")")

	if node.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(node.Alias))
	}
	return buf.String()
}

func visitationMinNode(v Visitor, node *MinNode) string {
	var buf bytes.Buffer
	buf.WriteString("MIN(")
	if node.Distinct {
		buf.WriteString("DISINCT ")
	}
	buf.WriteString(iterateVisitAndJoinOnComma(v, node.Expressions))
	buf.WriteString(")")

	if node.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(node.Alias))
	}
	return buf.String()
}

func visitationMaxNode(v Visitor, node *MaxNode) string {
	var buf bytes.Buffer
	buf.WriteString("MAX(")
	if node.Distinct {
		buf.WriteString("DISINCT ")
	}
	buf.WriteString(iterateVisitAndJoinOnComma(v, node.Expressions))
	buf.WriteString(")")

	if node.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(node.Alias))
	}
	return buf.String()
}

func visitationOrNode(v Visitor, node *OrNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" OR ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func visitationGreaterThanOrEqualNode(v Visitor, node *GreaterThanOrEqualNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" >= ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func visitationLessThanOrEqualNode(v Visitor, node *LessThanOrEqualNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" <= ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func visitationNotEqualNode(v Visitor, node *NotEqualNode) string {
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

func visitationValuesNode(v Visitor, node *ValuesNode) string {
	var buf bytes.Buffer
	buf.WriteString("VALUES (")
	rangevals := []string{}
	for _, value := range node.Values {
		rangevals = append(rangevals, v.Quote(value))
	}
	buf.WriteString(strings.Join(rangevals, COMMA))
	buf.WriteString(")")
	return buf.String()
}

func visitationTrueNode(v Visitor, node *TrueNode) string {
	return "TRUE"
}

func visitationFalseNode(v Visitor, node *FalseNode) string {
	return "FALSE"
}

func visitationDeleteStatementNode(v Visitor, node *DeleteStatementNode) string {
	var buf bytes.Buffer

	buf.WriteString("DELETE FROM ")
	buf.WriteString(v.Visit(node.Relation))

	if node.Wheres != nil && len(*node.Wheres) > 0 {
		buf.WriteString(WHERE)
		rangevals := []string{}
		for _, where := range *node.Wheres {
			rangevals = append(rangevals, v.Visit(where))
		}
		buf.WriteString(strings.Join(rangevals, AND))
	}

	return buf.String()
}

func visitationUpdateStatementNode(v Visitor, node *UpdateStatementNode) string {
	var buf bytes.Buffer

	var wheres []Visitable

	if (node.Orders == nil || len(*node.Orders) == 0) && node.Limit == nil {
		wheres = *node.Wheres
	} else {
		stmt := NewSelectStatementNode()
		core := stmt.Cores[0]

		if node.Relation != nil {
			core.SetFrom(node.Relation)
		}

		core.Wheres = node.Wheres

		if core.Selections == nil {
			core.Selections = &[]Visitable{}
		}

		*core.Selections = append(*core.Selections, node.Key)

		if node.Limit != nil {
			stmt.Limit = node.Limit
		}

		if node.Orders != nil && len(*node.Orders) > 0 {
			stmt.Orders = node.Orders
		}

		wheres = append(wheres, &InNode{
			Left:  node.Key,
			Right: []Visitable{stmt},
		})
	}
	buf.WriteString("UPDATE ")
	buf.WriteString(v.Visit(node.Relation))

	if node.Values != nil && len(*node.Values) > 0 {
		buf.WriteString(" SET ")
		buf.WriteString(iterateVisitAndJoinOnComma(v, *node.Values))
	}

	if wheres != nil && len(wheres) > 0 {
		buf.WriteString(WHERE)
		buf.WriteString(iterateVisitAndJoinOn(v, wheres, AND))
	}

	return buf.String()
}

func visitationInsertStatementNode(v Visitor, node *InsertStatementNode) string {
	var buf bytes.Buffer
	buf.WriteString("INSERT INTO ")
	buf.WriteString(v.Visit(node.Relation))

	if node.Columns != nil && len(*node.Columns) > 0 {
		buf.WriteString(" (")
		rangevals := []string{}
		for _, column := range *node.Columns {
			rangevals = append(rangevals, v.QuoteColumnName(column.Name))
		}
		buf.WriteString(strings.Join(rangevals, COMMA))
		buf.WriteString(")")
	}

	if node.Values != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.Visit(node.Values))
	}

	return buf.String()
}

func visitationNil() string {
	return "NULL"
}

func visitationWithNode(v Visitor, node *WithNode) string {
	var buf bytes.Buffer
	buf.WriteString("WITH ")
	buf.WriteString(v.Visit(node.Expr))
	return buf.String()
}

func visitationWithRecursiveNode(v Visitor, node *WithRecursiveNode) string {
	var buf bytes.Buffer
	buf.WriteString("WITH RECURSIVE ")
	buf.WriteString(v.Visit(node.Expr))
	return buf.String()
}

func visitationDistinctOnNode(v Visitor, node *DistinctOnNode) string {
	log.Fatal("NOT IMPLEMENTED FOR THIS DB")
	return ""
}

func visitationDistinctNode(v Visitor, node *DistinctNode) string {
	return DISTINCT
}

func visitationRangeNode(v Visitor, node *RangeNode) string {
	var buf bytes.Buffer
	buf.WriteString("RANGE")
	if node.Expr != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.Visit(node.Expr))
	}
	return buf.String()
}

func visitationBetweenNode(v Visitor, node *BetweenNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" BETWEEN ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func visitationCurrentRowNode(v Visitor, node *CurrentRowNode) string {
	return "CURRENT ROW"
}

func visitationPrecedingNode(v Visitor, node *PrecedingNode) string {
	var buf bytes.Buffer
	if node.Expr != nil {
		buf.WriteString(v.Visit(node.Expr))
	} else {
		buf.WriteString("UNBOUNDED")
	}
	buf.WriteString(" PRECEDING")
	return buf.String()
}

func visitationFollowingNode(v Visitor, node *FollowingNode) string {
	var buf bytes.Buffer
	if node.Expr != nil {
		buf.WriteString(v.Visit(node.Expr))
	} else {
		buf.WriteString("UNBOUNDED")
	}
	buf.WriteString(" FOLLOWING")
	return buf.String()
}

func visitationRowsNode(v Visitor, node *RowsNode) string {
	var buf bytes.Buffer
	buf.WriteString("ROWS")
	if node.Expr != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.Visit(node.Expr))
	}
	return buf.String()
}

func visitationNamedWindowNode(v Visitor, node *NamedWindowNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.QuoteColumnName(node.Name))
	buf.WriteString(" AS (")
	visitOrders := (node.Orders != nil && len(*node.Orders) > 0)
	visitFraming := (node.Framing != nil)
	if visitOrders {
		buf.WriteString("ORDER BY ")
		buf.WriteString(iterateVisitAndJoinOnComma(v, *node.Orders))
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

func visitationWindowNode(v Visitor, node *WindowNode) string {
	var buf bytes.Buffer
	buf.WriteString("(")
	visitOrders := (node.Orders != nil && len(*node.Orders) > 0)
	visitFraming := (node.Framing != nil)
	if visitOrders {
		buf.WriteString("ORDER BY ")
		buf.WriteString(iterateVisitAndJoinOnComma(v, *node.Orders))
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

func visitationGroupingNode(v Visitor, node *GroupingNode) string {
	var buf bytes.Buffer
	buf.WriteString("(")
	buf.WriteString(iterateVisitAndJoinOnComma(v, node.Expr))
	buf.WriteString(")")
	return buf.String()
}

func visitationLimitNode(v Visitor, node *LimitNode) string {
	var buf bytes.Buffer
	buf.WriteString("LIMIT ")
	buf.WriteString(v.Visit(node.Expr))
	return buf.String()
}

func visitationLockNode(v Visitor, node *LockNode) string {
	return v.Visit(node.Expr)
}

func visitationOffsetNode(v Visitor, node *OffsetNode) string {
	var buf bytes.Buffer
	buf.WriteString("OFFSET ")
	buf.WriteString(v.Visit(node.Expr))
	return buf.String()
}

func visitationAndNode(v Visitor, node *AndNode) string {
	var buf bytes.Buffer
	if node.Children != nil {
		buf.WriteString(iterateVisitAndJoinOn(v, *node.Children, AND))
	}
	return buf.String()
}

func visitationCountNode(v Visitor, node *CountNode) string {
	var buf bytes.Buffer
	buf.WriteString("COUNT(")
	if node.Distinct {
		buf.WriteString("DISINCT ")
	}
	buf.WriteString(iterateVisitAndJoinOnComma(v, node.Expressions))
	buf.WriteString(")")

	if node.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(node.Alias))
	}
	return buf.String()
}

func visitationAscendingNode(v Visitor, node *AscendingNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Expr))
	buf.WriteString(" ASC")
	return buf.String()
}

func visitationDescendingNode(v Visitor, node *DescendingNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Expr))
	buf.WriteString(" DESC")
	return buf.String()
}

func visitationOnNode(v Visitor, node *OnNode) string {
	var buf bytes.Buffer
	buf.WriteString("ON ")
	buf.WriteString(v.Visit(node.Expr))
	return buf.String()
}

func visitationExceptNode(v Visitor, node *ExceptNode) string {
	var buf bytes.Buffer
	buf.WriteString("( ")
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" EXCEPT ")
	buf.WriteString(v.Visit(node.Right))
	buf.WriteString(" )")
	return buf.String()
}

func visitationIntersectNode(v Visitor, node *IntersectNode) string {
	var buf bytes.Buffer
	buf.WriteString("( ")
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" INTERSECT ")
	buf.WriteString(v.Visit(node.Right))
	buf.WriteString(" )")
	return buf.String()
}

func visitationSelectManager(v Visitor, mgr *SelectManager) string {
	var buf bytes.Buffer
	buf.WriteString("(")
	buf.WriteString(mgr.ToSql())
	buf.WriteString(")")
	return buf.String()
}

func visitationMultiStatementManager(v Visitor, mgr *MultiStatementManager) string {
	var buf bytes.Buffer
	buf.WriteString(mgr.ToSql())
	return buf.String()
}

func visitationUnionNode(v Visitor, node *UnionNode) string {
	var buf bytes.Buffer
	buf.WriteString("( ")
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" UNION ")
	buf.WriteString(v.Visit(node.Right))
	buf.WriteString(" )")
	return buf.String()
}

func visitationUnionAllNode(v Visitor, node *UnionAllNode) string {
	var buf bytes.Buffer
	buf.WriteString("( ")
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" UNION ALL ")
	buf.WriteString(v.Visit(node.Right))
	buf.WriteString(" )")
	return buf.String()
}

func visitationLessThanNode(v Visitor, node *LessThanNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" < ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func visitationGreaterThanNode(v Visitor, node *GreaterThanNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" > ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func visitationAsNode(v Visitor, node *AsNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" AS ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func visitationGroupNode(v Visitor, node *GroupNode) string {
	return v.Visit(node.Expr)
}

func visitationHavingNode(v Visitor, node *HavingNode) string {
	var buf bytes.Buffer
	buf.WriteString("HAVING ")
	buf.WriteString(v.Visit(node.Expr))
	return buf.String()
}

func visitationExistsNode(v Visitor, node *ExistsNode) string {
	var buf bytes.Buffer
	buf.WriteString("EXISTS (")
	buf.WriteString(iterateVisitAndJoinOnComma(v, node.Expressions))
	buf.WriteString(")")

	if node.Alias != nil {
		buf.WriteString(" AS ")
		buf.WriteString(v.Visit(node.Alias))
	}
	return buf.String()
}

func visitationAttributeNode(v Visitor, node *AttributeNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.QuoteTableName(node.Relation))
	buf.WriteString(".")
	buf.WriteString(v.QuoteColumnName(node.Name))
	return buf.String()
}

func visitationEqualityNode(v Visitor, node *EqualityNode) string {
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

func visitationTable(v Visitor, table *Table) string {
	var buf bytes.Buffer
	buf.WriteString(v.QuoteTableName(table))
	if table.TableAlias != "" {
		buf.WriteString(SPACE)
		buf.WriteString(v.QuoteTableName(&TableAliasNode{
			Relation: table,
			Name:     table.TableAlias,
			Quoted:   true,
		}))
	}
	return buf.String()
}

func visitationJoinSourceNode(v Visitor, node *JoinSource) string {
	var buf bytes.Buffer
	if node.Left != nil {
		buf.WriteString(v.Visit(node.Left))
	}
	buf.WriteString(iterateVisitAndJoinOn(v, node.Right, SPACE))
	return buf.String()
}

func visitationOuterJoinNode(v Visitor, node *OuterJoinNode) string {
	var buf bytes.Buffer
	buf.WriteString("LEFT OUTER JOIN ")
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(SPACE)
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func visitationInnerJoinNode(v Visitor, node *InnerJoinNode) string {
	var buf bytes.Buffer
	buf.WriteString(" INNER JOIN ")
	buf.WriteString(v.Visit(node.Left))
	if node.Right != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.Visit(node.Right))
	}
	return buf.String()
}

func visitationSqlLiteralNode(v Visitor, node SqlLiteralNode) string {
	return node.Raw
}

func visitationTableAliasNode(v Visitor, node *TableAliasNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Relation))
	buf.WriteString(" ")
	buf.WriteString(v.QuoteTableName(node))
	return buf.String()
}

func visitationSelectCoreNode(v Visitor, node *SelectCoreNode) string {
	var buf bytes.Buffer

	buf.WriteString("SELECT")

	// Add TOP statement to the buffer
	if node.Top != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.Visit(node.Top))
	}

	if node.SetQuantifier != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.Visit(node.SetQuantifier))
	}

	// add select projections
	if node.Selections != nil && len(*node.Selections) > 0 {
		buf.WriteString(SPACE)
		buf.WriteString(iterateVisitAndJoinOnComma(v, *node.Selections))
	}

	// add FROM statement to the buffer
	if node.Source != nil && node.Source.Left != nil {
		// assert the source is a *Table and check the length of the name
		if t, ok := node.Source.Left.(*Table); ok && t.Name != "" {
			buf.WriteString(" FROM ")
			buf.WriteString(v.Visit(node.Source))
		}
	}

	// add WHERE statement to the buffer
	if node.Wheres != nil && len(*node.Wheres) > 0 {
		buf.WriteString(WHERE)
		buf.WriteString(iterateVisitAndJoinOnComma(v, *node.Wheres))
	}

	// add GROUP BY statement to the buffer
	if node.Groups != nil && len(*node.Groups) > 0 {
		buf.WriteString(GROUP_BY)
		buf.WriteString(iterateVisitAndJoinOnComma(v, *node.Groups))
	}

	// add HAVING statement to the buffer
	if node.Having != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.Visit(node.Having))
	}

	// add WINDOW statements to the buffer
	if node.Windows != nil && len(*node.Windows) > 0 {
		buf.WriteString(WINDOW)
		buf.WriteString(iterateVisitAndJoinOnComma(v, *node.Windows))
	}

	return buf.String()
}

func visitationSelectStatementNode(v Visitor, node *SelectStatementNode) string {
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
		buf.WriteString(iterateVisitAndJoinOnComma(v, *node.Orders))
	}

	// add LIMIT clause to the buffer
	if node.Limit != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.Visit(node.Limit))
	}

	// add OFFSET clause to the buffer
	if node.Offset != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.Visit(node.Offset))
	}

	// add LOCK clause to the buffer
	if node.Lock != nil {
		buf.WriteString(SPACE)
		buf.WriteString(v.Visit(node.Lock))
	}

	return strings.TrimSpace(buf.String())
}

func iterateVisitAndJoinOnComma(v Visitor, vistables []Visitable) string {
	return iterateVisitAndJoinOn(v, vistables, COMMA)
}

func iterateVisitAndJoinOn(v Visitor, visitables []Visitable, sep string) string {
	rangevals := []string{}
	for _, vis := range visitables {
		rangevals = append(rangevals, v.Visit(vis))
	}
	return strings.Join(rangevals, sep)
}
