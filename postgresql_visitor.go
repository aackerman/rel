package rel

import (
	"bytes"
	"log"
	"runtime/debug"
)

// Used to handle generating Postgres specific sql
type PostgreSQLVisitor struct {
	Conn Connector
}

func (v PostgreSQLVisitor) Accept(visitable Visitable) string {
	return v.Visit(visitable)
}

func (v PostgreSQLVisitor) Visit(visitable Visitable) string {
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
	case *JoinSource:
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
	case *DistinctOnNode:
		return v.VisitDistinctOnNode(val)
	default:
		debug.PrintStack()
		log.Fatalf("PostgreSQLVisitor#Visit unable to handle type %T", visitable)
		return ""
	}
}

func (v *PostgreSQLVisitor) QuoteTableName(visitable Visitable) string {
	if alias, ok := visitable.(*TableAliasNode); ok {
		if !alias.Quoted {
			return alias.Name.Raw
		}
	}
	return v.Conn.QuoteTableName(visitable.String())
}

func (v *PostgreSQLVisitor) Quote(thing interface{}) string {
	return v.Conn.Quote(thing)
}

func (v *PostgreSQLVisitor) QuoteColumnName(literal SqlLiteralNode) string {
	return v.Conn.QuoteColumnName(literal.Raw)
}

func (v *PostgreSQLVisitor) VisitMatchesNode(node *MatchesNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" ILIKE ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v *PostgreSQLVisitor) VisitDoesNotMatchNode(node *DoesNotMatchNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" NOT ILIKE ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v *PostgreSQLVisitor) VisitDistinctOnNode(node *DistinctOnNode) string {
	var buf bytes.Buffer
	buf.WriteString("DISTINCT ON ( ")
	buf.WriteString(v.Visit(node.Expr))
	buf.WriteString(" )")
	return buf.String()
}

func (v *PostgreSQLVisitor) VisitTopNode(node *TopNode) string {
	return visitationTopNode(v, node)
}
func (v *PostgreSQLVisitor) VisitOrderingNode(node *OrderingNode) string {
	return visitationOrderingNode(v, node)
}
func (v *PostgreSQLVisitor) VisitUnqualifiedColumnNode(node *UnqualifiedColumnNode) string {
	return visitationUnqualifiedColumnNode(v, node)
}
func (v *PostgreSQLVisitor) VisitAssignmentNode(node *AssignmentNode) string {
	return visitationAssignmentNode(v, node)
}
func (v *PostgreSQLVisitor) VisitOverNode(node *OverNode) string {
	return visitationOverNode(v, node)
}
func (v *PostgreSQLVisitor) VisitQuotedNode(node *QuotedNode) string {
	return visitationQuotedNode(v, node)
}
func (v *PostgreSQLVisitor) VisitInfixOperationNode(node *InfixOperationNode) string {
	return visitationInfixOperationNode(v, node)
}
func (v *PostgreSQLVisitor) VisitExtractNode(node *ExtractNode) string {
	return visitationExtractNode(v, node)
}
func (v *PostgreSQLVisitor) VisitBinNode(node *BinNode) string {
	return visitationBinNode(v, node)
}
func (v *PostgreSQLVisitor) VisitNotNode(node *NotNode) string {
	return visitationNotNode(v, node)
}
func (v *PostgreSQLVisitor) VisitNotInNode(node *NotInNode) string {
	return visitationNotInNode(v, node)
}
func (v *PostgreSQLVisitor) VisitInNode(node *InNode) string {
	return visitationInNode(v, node)
}
func (v *PostgreSQLVisitor) VisitNamedFunctionNode(node *NamedFunctionNode) string {
	return visitationNamedFunctionNode(v, node)
}
func (v *PostgreSQLVisitor) VisitSumNode(node *SumNode) string {
	return visitationSumNode(v, node)
}
func (v *PostgreSQLVisitor) VisitAvgNode(node *AvgNode) string {
	return visitationAvgNode(v, node)
}
func (v *PostgreSQLVisitor) VisitMinNode(node *MinNode) string {
	return visitationMinNode(v, node)
}
func (v *PostgreSQLVisitor) VisitMaxNode(node *MaxNode) string {
	return visitationMaxNode(v, node)
}
func (v *PostgreSQLVisitor) VisitOrNode(node *OrNode) string {
	return visitationOrNode(v, node)
}
func (v *PostgreSQLVisitor) VisitGreaterThanOrEqualNode(node *GreaterThanOrEqualNode) string {
	return visitationGreaterThanOrEqualNode(v, node)
}
func (v *PostgreSQLVisitor) VisitLessThanOrEqualNode(node *LessThanOrEqualNode) string {
	return visitationLessThanOrEqualNode(v, node)
}
func (v *PostgreSQLVisitor) VisitNotEqualNode(node *NotEqualNode) string {
	return visitationNotEqualNode(v, node)
}
func (v *PostgreSQLVisitor) VisitValuesNode(node *ValuesNode) string {
	return visitationValuesNode(v, node)
}
func (v *PostgreSQLVisitor) VisitTrueNode(node TrueNode) string {
	return visitationTrueNode(v, node)
}
func (v *PostgreSQLVisitor) VisitFalseNode(node FalseNode) string {
	return visitationFalseNode(v, node)
}
func (v *PostgreSQLVisitor) VisitDeleteStatementNode(node DeleteStatementNode) string {
	return visitationDeleteStatementNode(v, node)
}
func (v *PostgreSQLVisitor) VisitUpdateStatementNode(node UpdateStatementNode) string {
	return visitationUpdateStatementNode(v, node)
}
func (v *PostgreSQLVisitor) VisitInsertStatementNode(node *InsertStatementNode) string {
	return visitationInsertStatementNode(v, node)
}
func (v *PostgreSQLVisitor) VisitNil() string {
	return visitationNil()
}
func (v *PostgreSQLVisitor) VisitWithNode(node *WithNode) string {
	return visitationWithNode(v, node)
}
func (v *PostgreSQLVisitor) VisitWithRecursiveNode(node *WithRecursiveNode) string {
	return visitationWithRecursiveNode(v, node)
}
func (v *PostgreSQLVisitor) VisitDistinctNode(node *DistinctNode) string {
	return visitationDistinctNode(v, node)
}
func (v *PostgreSQLVisitor) VisitRangeNode(node *RangeNode) string {
	return visitationRangeNode(v, node)
}
func (v *PostgreSQLVisitor) VisitBetweenNode(node *BetweenNode) string {
	return visitationBetweenNode(v, node)
}
func (v *PostgreSQLVisitor) VisitCurrentRowNode(node *CurrentRowNode) string {
	return visitationCurrentRowNode(v, node)
}
func (v *PostgreSQLVisitor) VisitPrecedingNode(node *PrecedingNode) string {
	return visitationPrecedingNode(v, node)
}
func (v *PostgreSQLVisitor) VisitFollowingNode(node *FollowingNode) string {
	return visitationFollowingNode(v, node)
}
func (v *PostgreSQLVisitor) VisitRowsNode(node *RowsNode) string {
	return visitationRowsNode(v, node)
}
func (v *PostgreSQLVisitor) VisitNamedWindowNode(node *NamedWindowNode) string {
	return visitationNamedWindowNode(v, node)
}
func (v *PostgreSQLVisitor) VisitWindowNode(node *WindowNode) string {
	return visitationWindowNode(v, node)
}
func (v *PostgreSQLVisitor) VisitGroupingNode(node *GroupingNode) string {
	return visitationGroupingNode(v, node)
}
func (v *PostgreSQLVisitor) VisitLimitNode(node *LimitNode) string {
	return visitationLimitNode(v, node)
}
func (v *PostgreSQLVisitor) VisitLockNode(node *LockNode) string {
	return visitationLockNode(v, node)
}
func (v *PostgreSQLVisitor) VisitOffsetNode(node *OffsetNode) string {
	return visitationOffsetNode(v, node)
}
func (v *PostgreSQLVisitor) VisitAndNode(node *AndNode) string {
	return visitationAndNode(v, node)
}
func (v *PostgreSQLVisitor) VisitCountNode(node *CountNode) string {
	return visitationCountNode(v, node)
}
func (v *PostgreSQLVisitor) VisitAscendingNode(node *AscendingNode) string {
	return visitationAscendingNode(v, node)
}
func (v *PostgreSQLVisitor) VisitDescendingNode(node *DescendingNode) string {
	return visitationDescendingNode(v, node)
}
func (v *PostgreSQLVisitor) VisitOnNode(node *OnNode) string {
	return visitationOnNode(v, node)
}
func (v *PostgreSQLVisitor) VisitExceptNode(node *ExceptNode) string {
	return visitationExceptNode(v, node)
}
func (v *PostgreSQLVisitor) VisitIntersectNode(node *IntersectNode) string {
	return visitationIntersectNode(v, node)
}
func (v *PostgreSQLVisitor) VisitSelectManager(node *SelectManager) string {
	return visitationSelectManager(v, node)
}
func (v *PostgreSQLVisitor) VisitMultiStatementManager(node *MultiStatementManager) string {
	return visitationMultiStatementManager(v, node)
}
func (v *PostgreSQLVisitor) VisitUnionNode(node *UnionNode) string {
	return visitationUnionNode(v, node)
}
func (v *PostgreSQLVisitor) VisitUnionAllNode(node *UnionAllNode) string {
	return visitationUnionAllNode(v, node)
}
func (v *PostgreSQLVisitor) VisitLessThanNode(node *LessThanNode) string {
	return visitationLessThanNode(v, node)
}
func (v *PostgreSQLVisitor) VisitGreaterThanNode(node *GreaterThanNode) string {
	return visitationGreaterThanNode(v, node)
}
func (v *PostgreSQLVisitor) VisitAsNode(node *AsNode) string {
	return visitationAsNode(v, node)
}
func (v *PostgreSQLVisitor) VisitGroupNode(node *GroupNode) string {
	return visitationGroupNode(v, node)
}
func (v *PostgreSQLVisitor) VisitHavingNode(node *HavingNode) string {
	return visitationHavingNode(v, node)
}
func (v *PostgreSQLVisitor) VisitExistsNode(node *ExistsNode) string {
	return visitationExistsNode(v, node)
}
func (v *PostgreSQLVisitor) VisitAttributeNode(node *AttributeNode) string {
	return visitationAttributeNode(v, node)
}
func (v *PostgreSQLVisitor) VisitEqualityNode(node *EqualityNode) string {
	return visitationEqualityNode(v, node)
}
func (v *PostgreSQLVisitor) VisitTable(table *Table) string {
	return visitationTable(v, table)
}
func (v *PostgreSQLVisitor) VisitJoinSourceNode(node *JoinSource) string {
	return visitationJoinSourceNode(v, node)
}
func (v *PostgreSQLVisitor) VisitOuterJoinNode(node OuterJoinNode) string {
	return visitationOuterJoinNode(v, node)
}
func (v *PostgreSQLVisitor) VisitInnerJoinNode(node *InnerJoinNode) string {
	return visitationInnerJoinNode(v, node)
}
func (v *PostgreSQLVisitor) VisitSqlLiteralNode(node SqlLiteralNode) string {
	return visitationSqlLiteralNode(v, node)
}
func (v *PostgreSQLVisitor) VisitTableAliasNode(node *TableAliasNode) string {
	return visitationTableAliasNode(v, node)
}
func (v *PostgreSQLVisitor) VisitSelectCoreNode(node *SelectCoreNode) string {
	return visitationSelectCoreNode(v, node)
}
func (v *PostgreSQLVisitor) VisitSelectStatementNode(node *SelectStatementNode) string {
	return visitationSelectStatementNode(v, node)
}
