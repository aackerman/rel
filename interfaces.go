package rel

type Visitable interface {
	NewTrueNode() TrueNode
	NewFalseNode() FalseNode
	NewTableAliasNode(*Table, SqlLiteralNode) *TableAliasNode
	NewStringJoinNode() StringJoinNode
	NewInnerJoinNode() InnerJoinNode
	NewOuterJoinNode() OuterJoinNode
	NewAndNode(...Visitable) *AndNode
	NewOnNode(Visitable) *OnNode
	NewNotNode() NotNode
	NewGroupingNode() GroupingNode
	String() string
}

type Visitor interface {
	Accept(Visitable) string
	QuoteColumnName(SqlLiteralNode) string
	QuoteTableName(Visitable) string
	Quote(interface{}) string
	Visit(Visitable) string
	VisitAndNode(*AndNode) string
	VisitAscendingNode(*AscendingNode) string
	VisitAsNode(*AsNode) string
	VisitAssignmentNode(*AssignmentNode) string
	VisitAttributeNode(*AttributeNode) string
	VisitAvgNode(*AvgNode) string
	VisitBetweenNode(*BetweenNode) string
	VisitBinNode(*BinNode) string
	VisitCountNode(*CountNode) string
	VisitCurrentRowNode(*CurrentRowNode) string
	VisitDescendingNode(*DescendingNode) string
	VisitDistinctNode(*DistinctNode) string
	VisitDoesNotMatchNode(*DoesNotMatchNode) string
	VisitEqualityNode(*EqualityNode) string
	VisitExceptNode(*ExceptNode) string
	VisitExistsNode(*ExistsNode) string
	VisitExtractNode(*ExtractNode) string
	VisitFollowingNode(*FollowingNode) string
	VisitGreaterThanNode(*GreaterThanNode) string
	VisitGreaterThanOrEqualNode(*GreaterThanOrEqualNode) string
	VisitGroupingNode(*GroupingNode) string
	VisitGroupNode(*GroupNode) string
	VisitHavingNode(*HavingNode) string
	VisitInfixOperationNode(*InfixOperationNode) string
	VisitInnerJoinNode(*InnerJoinNode) string
	VisitInNode(*InNode) string
	VisitInsertStatementNode(*InsertStatementNode) string
	VisitIntersectNode(*IntersectNode) string
	VisitJoinSourceNode(*JoinSource) string
	VisitLessThanNode(*LessThanNode) string
	VisitLessThanOrEqualNode(*LessThanOrEqualNode) string
	VisitLimitNode(*LimitNode) string
	VisitLockNode(*LockNode) string
	VisitMatchesNode(*MatchesNode) string
	VisitMaxNode(*MaxNode) string
	VisitMinNode(*MinNode) string
	VisitMultiStatementManager(*MultiStatementManager) string
	VisitNamedFunctionNode(*NamedFunctionNode) string
	VisitNamedWindowNode(*NamedWindowNode) string
	VisitNil() string
	VisitNotEqualNode(*NotEqualNode) string
	VisitNotInNode(*NotInNode) string
	VisitNotNode(*NotNode) string
	VisitOffsetNode(*OffsetNode) string
	VisitOnNode(*OnNode) string
	VisitOrNode(*OrNode) string
	VisitOverNode(*OverNode) string
	VisitPrecedingNode(*PrecedingNode) string
	VisitQuotedNode(*QuotedNode) string
	VisitRangeNode(*RangeNode) string
	VisitRowsNode(*RowsNode) string
	VisitSelectCoreNode(*SelectCoreNode) string
	VisitSelectManager(*SelectManager) string
	VisitSelectStatementNode(*SelectStatementNode) string
	VisitSqlLiteralNode(SqlLiteralNode) string
	VisitSumNode(*SumNode) string
	VisitTable(*Table) string
	VisitTableAliasNode(*TableAliasNode) string
	VisitTopNode(*TopNode) string
	VisitUnionAllNode(*UnionAllNode) string
	VisitUnionNode(*UnionNode) string
	VisitUnqualifiedColumnNode(*UnqualifiedColumnNode) string
	VisitValuesNode(*ValuesNode) string
	VisitWindowNode(*WindowNode) string
	VisitWithNode(*WithNode) string
	VisitWithRecursiveNode(*WithRecursiveNode) string
}

type TreeManager interface {
	ToSql() string
}

type Engine interface {
	Visitor() Visitor
}

type Connector interface {
	QuoteTableName(string) string
	QuoteColumnName(string) string
	Quote(interface{}) string
}

func Register(engine Engine) {
	DefaultEngine = engine
}
