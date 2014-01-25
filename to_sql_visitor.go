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
	default:
		debug.PrintStack()
		log.Fatalf("ToSqlVisitor#Visit unable to handle type %T", visitable)
		return ""
	}
}
