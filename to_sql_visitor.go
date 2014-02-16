package rel

import (
	"log"
	"runtime/debug"
)

type ToSqlVisitor struct {
	Conn Connector
}

func (v *ToSqlVisitor) Accept(visitable Visitable) string {
	return v.Visit(visitable)
}
func (v *ToSqlVisitor) Visit(visitable Visitable) string {
	switch node := visitable.(type) {
	case nil:
		return visitationNil()
	case *SelectStatementNode:
		return visitationSelectStatementNode(v, node)
	case *InNode:
		return visitationInNode(v, node)
	case SqlLiteralNode:
		return visitationSqlLiteralNode(v, node)
	case *SqlLiteralNode:
		return visitationSqlLiteralNode(v, *node)
	case *JoinSource:
		return visitationJoinSourceNode(v, node)
	case *EqualityNode:
		return visitationEqualityNode(v, node)
	case *HavingNode:
		return visitationHavingNode(v, node)
	case *AttributeNode:
		return visitationAttributeNode(v, node)
	case *GroupNode:
		return visitationGroupNode(v, node)
	case *ExistsNode:
		return visitationExistsNode(v, node)
	case *AsNode:
		return visitationAsNode(v, node)
	case *LessThanNode:
		return visitationLessThanNode(v, node)
	case *UnionNode:
		return visitationUnionNode(v, node)
	case *UnionAllNode:
		return visitationUnionAllNode(v, node)
	case *SelectManager:
		return visitationSelectManager(v, node)
	case *GreaterThanNode:
		return visitationGreaterThanNode(v, node)
	case *IntersectNode:
		return visitationIntersectNode(v, node)
	case *ExceptNode:
		return visitationExceptNode(v, node)
	case *OnNode:
		return visitationOnNode(v, node)
	case *AscendingNode:
		return visitationAscendingNode(v, node)
	case *DescendingNode:
		return visitationDescendingNode(v, node)
	case *CountNode:
		return visitationCountNode(v, node)
	case *AndNode:
		return visitationAndNode(v, node)
	case *TableAliasNode:
		return visitationTableAliasNode(v, node)
	case *InnerJoinNode:
		return visitationInnerJoinNode(v, node)
	case *GroupingNode:
		return visitationGroupingNode(v, node)
	case *NamedWindowNode:
		return visitationNamedWindowNode(v, node)
	case *WindowNode:
		return visitationWindowNode(v, node)
	case *RowsNode:
		return visitationRowsNode(v, node)
	case *LockNode:
		return visitationLockNode(v, node)
	case *PrecedingNode:
		return visitationPrecedingNode(v, node)
	case *FollowingNode:
		return visitationFollowingNode(v, node)
	case *CurrentRowNode:
		return visitationCurrentRowNode(v, node)
	case *BetweenNode:
		return visitationBetweenNode(v, node)
	case *RangeNode:
		return visitationRangeNode(v, node)
	case *DistinctNode:
		return visitationDistinctNode(v, node)
	case *WithNode:
		return visitationWithNode(v, node)
	case *WithRecursiveNode:
		return visitationWithRecursiveNode(v, node)
	case *Table:
		if node == nil {
			return visitationNil()
		}
		return visitationTable(v, node)
	case *MultiStatementManager:
		return visitationMultiStatementManager(v, node)
	case *InsertStatementNode:
		return visitationInsertStatementNode(v, node)
	case *SelectCoreNode:
		return visitationSelectCoreNode(v, node)
	case *NotEqualNode:
		return visitationNotEqualNode(v, node)
	case *NotNode:
		return visitationNotNode(v, node)
	case *GreaterThanOrEqualNode:
		return visitationGreaterThanOrEqualNode(v, node)
	case *LessThanOrEqualNode:
		return visitationLessThanOrEqualNode(v, node)
	case *OrNode:
		return visitationOrNode(v, node)
	case *AvgNode:
		return visitationAvgNode(v, node)
	case *NamedFunctionNode:
		return visitationNamedFunctionNode(v, node)
	case *SumNode:
		return visitationSumNode(v, node)
	case *MinNode:
		return visitationMinNode(v, node)
	case *MaxNode:
		return visitationMaxNode(v, node)
	case *MatchesNode:
		return visitationMatchesNode(v, node)
	case *DoesNotMatchNode:
		return visitationDoesNotMatchNode(v, node)
	case *NotInNode:
		return visitationNotInNode(v, node)
	case *BinNode:
		return visitationBinNode(v, node)
	case *ExtractNode:
		return visitationExtractNode(v, node)
	case *InfixOperationNode:
		return visitationInfixOperationNode(v, node)
	case *QuotedNode:
		return visitationQuotedNode(v, node)
	case *OverNode:
		return visitationOverNode(v, node)
	case *AssignmentNode:
		return visitationAssignmentNode(v, node)
	case *UnqualifiedColumnNode:
		return visitationUnqualifiedColumnNode(v, node)
	case *DistinctOnNode:
		return visitationDistinctOnNode(v, node)
	case *OuterJoinNode:
		return visitationOuterJoinNode(v, node)
	case *OffsetNode:
		return visitationOffsetNode(v, node)
	case *LimitNode:
		return visitationLimitNode(v, node)
	case *UpdateStatementNode:
		return visitationUpdateStatementNode(v, node)
	case *DeleteStatementNode:
		return visitationDeleteStatementNode(v, node)
	case *FalseNode:
		return visitationFalseNode(v, node)
	case *TrueNode:
		return visitationTrueNode(v, node)
	case *ValuesNode:
		return visitationValuesNode(v, node)
	case *OrderingNode:
		return visitationOrderingNode(v, node)
	case *TopNode:
		return visitationTopNode(v, node)
	default:
		debug.PrintStack()
		log.Fatalf("ToSqlVisitor#Visit unable to handle type %T", visitable)
		return ""
	}
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
