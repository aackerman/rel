package rel

type Visitable interface {
	NewTrueNode() *TrueNode
	NewFalseNode() *FalseNode
	NewTableAliasNode(*Table, SqlLiteralNode) *TableAliasNode
	NewInnerJoinNode() *InnerJoinNode
	NewOuterJoinNode() *OuterJoinNode
	NewAndNode(...Visitable) *AndNode
	NewOnNode(Visitable) *OnNode
	NewNotNode() *NotNode
	NewGroupingNode() *GroupingNode
	String() string
}

type Visitor interface {
	Accept(Visitable) string
	QuoteColumnName(SqlLiteralNode) string
	QuoteTableName(Visitable) string
	Quote(interface{}) string
	Visit(Visitable) string
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
	RelEngine = engine
}
