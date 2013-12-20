package arel

type BaseNode struct {
	NodeFactory
}

func CreateBaseNode() BaseNode {
	return BaseNode{NodeFactory{}}
}

type AstNode interface {
	CreateTrue() TrueNode
	CreateFalse() FalseNode
	CreateTableAlias(*Table, string) TableAliasNode
	CreateStringJoin() StringJoinNode
	CreateInnerJoin() InnerJoinNode
	CreateOuterJoin() OuterJoinNode
	CreateAnd() AndNode
	CreateOn() OnNode
	CreateNot() NotNode
	CreateGrouping() GroupingNode
	CreateLower() NamedFunctionNode
}
