package arel

type SqlAst interface {
	CreateTrue() *TrueNode
	CreateFalse() *FalseNode
	CreateTableAlias(*TableAliasNode, string) *TableAliasNode
	CreateJoin(string) *JoinNode
	CreateAnd([]interface{}) *AndNode
	CreateOn(string) *OnNode
	CreateGrouping(string) *GroupingNode
}
