package arel

type BaseNode struct{}

func (b BaseNode) CreateTrue() TrueNode {
	return TrueNode{}
}

func (b BaseNode) CreateFalse() FalseNode {
	return FalseNode{}
}

func (b BaseNode) CreateTableAlias(t *Table, name string) TableAliasNode {
	return TableAliasNode{Name: name, Table: t}
}

func (b BaseNode) CreateStringJoin() StringJoinNode {
	return StringJoinNode{}
}

func (b BaseNode) CreateInnerJoin() InnerJoinNode {
	return InnerJoinNode{}
}

func (b BaseNode) CreateOuterJoin() OuterJoinNode {
	return OuterJoinNode{}
}

func (b BaseNode) CreateAnd() AndNode {
	return AndNode{}
}

func (b BaseNode) CreateOn() OnNode {
	return OnNode{}
}

func (b BaseNode) CreateNot() NotNode {
	return NotNode{}
}

func (b BaseNode) CreateGrouping() GroupingNode {
	return GroupingNode{}
}

func (b BaseNode) CreateLower() NamedFunctionNode {
	return NamedFunctionNode{}
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
