package arel

type NodeFactory struct{}

func (n NodeFactory) CreateTrue() TrueNode {
	return TrueNode{}
}

func (n NodeFactory) CreateFalse() FalseNode {
	return FalseNode{}
}

func (n NodeFactory) CreateTableAlias(t *Table, name string) TableAliasNode {
	return TableAliasNode{Name: name, Table: t}
}

func (n NodeFactory) CreateStringJoin() StringJoinNode {
	return StringJoinNode{}
}

func (n NodeFactory) CreateInnerJoin() InnerJoinNode {
	return InnerJoinNode{}
}

func (n NodeFactory) CreateOuterJoin() OuterJoinNode {
	return OuterJoinNode{}
}

func (n NodeFactory) CreateAnd() AndNode {
	return AndNode{}
}

func (n NodeFactory) CreateOn() OnNode {
	return OnNode{}
}

func (n NodeFactory) CreateNot() NotNode {
	return NotNode{}
}

func (n NodeFactory) CreateGrouping() GroupingNode {
	return GroupingNode{}
}

func (n NodeFactory) CreateLower() NamedFunctionNode {
	return NamedFunctionNode{}
}
