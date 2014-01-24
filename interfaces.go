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
	Visit(Visitable) string
	VisitNil() string
	VisitSelectStatementNode(*SelectStatementNode) string
	VisitInNode(*InNode) string
	VisitSqlLiteralNode(SqlLiteralNode) string
	VisitJoinSourceNode(*JoinSource) string
	VisitEqualityNode(*EqualityNode) string
	VisitHavingNode(*HavingNode) string
	VisitAttributeNode(*AttributeNode) string
	VisitGroupNode(*GroupNode) string
	VisitExistsNode(*ExistsNode) string
	VisitAsNode(*AsNode) string
	VisitLessThanNode(*LessThanNode) string
	VisitUnionNode(*UnionNode) string
	VisitUnionAllNode(*UnionAllNode) string
	VisitSelectManager(*SelectManager) string
	VisitGreaterThanNode(*GreaterThanNode) string
	VisitIntersectNode(*IntersectNode) string
	VisitExceptNode(*ExceptNode) string
	VisitOnNode(*OnNode) string
	VisitAscendingNode(*AscendingNode) string
	VisitDescendingNode(*DescendingNode) string
	VisitCountNode(*CountNode) string
	VisitAndNode(*AndNode) string
	VisitTableAliasNode(*TableAliasNode) string
	VisitInnerJoinNode(*InnerJoinNode) string
	VisitGroupingNode(*GroupingNode) string
	VisitNamedWindowNode(*NamedWindowNode) string
	VisitWindowNode(*WindowNode) string
	VisitRowsNode(*RowsNode) string
	VisitPrecedingNode(*PrecedingNode) string
	VisitFollowingNode(*FollowingNode) string
	VisitCurrentRowNode(*CurrentRowNode) string
	VisitBetweenNode(*BetweenNode) string
	VisitRangeNode(*RangeNode) string
	VisitDistinctNode(*DistinctNode) string
	VisitWithNode(*WithNode) string
	VisitWithRecursiveNode(*WithRecursiveNode) string
	VisitTable(*Table) string
	VisitMultiStatementManager(*MultiStatementManager) string
	VisitInsertStatementNode(*InsertStatementNode) string
	VisitValuesNode(*ValuesNode) string
	VisitSelectCoreNode(*SelectCoreNode) string
	VisitNotEqualNode(*NotEqualNode) string
	VisitNotNode(*NotNode) string
	VisitGreaterThanOrEqualNode(*GreaterThanOrEqualNode) string
	VisitLessThanOrEqualNode(*LessThanOrEqualNode) string
	VisitOrNode(*OrNode) string
	VisitAvgNode(*AvgNode) string
	VisitNamedFunctionNode(*NamedFunctionNode) string
	VisitSumNode(*SumNode) string
	VisitMinNode(*MinNode) string
	VisitMaxNode(*MaxNode) string
	VisitMatchesNode(*MatchesNode) string
	VisitDoesNotMatchNode(*DoesNotMatchNode) string
	VisitNotInNode(*NotInNode) string
	VisitBinNode(*BinNode) string
	VisitExtractNode(*ExtractNode) string
	VisitInfixOperationNode(*InfixOperationNode) string
	VisitQuotedNode(*QuotedNode) string
	VisitOverNode(*OverNode) string
	VisitAssignmentNode(*AssignmentNode) string
	VisitUnqualifiedColumnNode(*UnqualifiedColumnNode) string
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
