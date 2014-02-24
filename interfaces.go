package rel

type Visitable interface {
	NewTrueNode() *TrueNode
	NewFalseNode() *FalseNode
	NewTableAliasNode(*Table, string) *TableAliasNode
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

func RegisterEngine(engine Engine) {
	RelEngine = engine
}

func RegisterDatabase(db string) {
	switch db {
	case "postgresql":
		RelEngine = DefaultEngine{&PostgreSQLVisitor{Conn: DefaultConnector{}}}
	case "sqlite":
		RelEngine = DefaultEngine{&SQLiteVisitor{Conn: DefaultConnector{}}}
	case "mysql":
		RelEngine = DefaultEngine{&MysqlVisitor{Conn: DefaultConnector{}}}
	}
}
