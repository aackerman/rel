package rel

type Visitable interface {
	NewTrueNode() TrueNode
	NewFalseNode() FalseNode
	NewTableAliasNode(*Table, SqlLiteralNode) *TableAliasNode
	NewStringJoinNode() StringJoinNode
	NewInnerJoinNode() InnerJoinNode
	NewOuterJoinNode() OuterJoinNode
	NewAndNode(...Visitable) *AndNode
	NewOnNode() OnNode
	NewNotNode() NotNode
	NewGroupingNode() GroupingNode
}

type Predicator interface {
	Eq(Visitable) *EqualityNode
	EqAny(...Visitable) *GroupingNode
	EqAll(...Visitable) *GroupingNode
	Lt(Visitable) *LessThanNode
	LtEq(Visitable) *LessThanOrEqualNode
	LtAny(...Visitable) *GroupingNode
	LtAll(...Visitable) *GroupingNode
	Gt(Visitable) *GreaterThanNode
	GtEq(Visitable) *GreaterThanOrEqualNode
	GtAny(...Visitable) *GroupingNode
	GtAll(...Visitable) *GroupingNode
	In([]Visitable) Visitable
	NotIn([]Visitable) Visitable
	NotEq(Visitable) *NotEqualNode
	Matches(SqlLiteralNode) *MatchesNode
	DoesNotMatch(SqlLiteralNode) *DoesNotMatchNode
	GroupAny(...Visitable) *GroupingNode
	GroupAll(...Visitable) *GroupingNode
	Visitable
}

type Visitor interface {
	Accept(Visitable) string
	Visit(Visitable) string
}

type TreeManager interface {
	ToSql() string
}

type Connection struct{}

func (c *Connection) QuoteTableName(name string) string {
	return "\"" + name + "\""
}

func (c *Connection) QuoteColumnName(name string) string {
	return "\"" + name + "\""
}

type Engine interface {
	Visitor() Visitor
	Connection() *Connection
	QuoteTableName(string) string
	QuoteColumnName(string) string
	TableExists(string) bool
}

func Register(name string, engine Engine) {
	DefaultEngine = engine
}
