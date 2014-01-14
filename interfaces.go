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
