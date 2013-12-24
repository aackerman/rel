package arel

type NodeFactory struct{}

func (n NodeFactory) NewTrueNode() TrueNode {
	return TrueNode{}
}

func (n NodeFactory) NewFalseNode() FalseNode {
	return FalseNode{}
}

func (n NodeFactory) NewTableAliasNode(t *Table, name string) TableAliasNode {
	return TableAliasNode{Name: name, Table: t}
}

func (n NodeFactory) NewStringJoinNode() StringJoinNode {
	return StringJoinNode{}
}

func (n NodeFactory) NewInnerJoinNode() InnerJoinNode {
	return InnerJoinNode{}
}

func (n NodeFactory) NewOuterJoinNode() OuterJoinNode {
	return OuterJoinNode{}
}

func (n NodeFactory) NewAndNode() AndNode {
	return AndNode{}
}

func (n NodeFactory) NewOnNode() OnNode {
	return OnNode{}
}

func (n NodeFactory) NewNotNode() NotNode {
	return NotNode{}
}

func (n NodeFactory) NewGroupingNode() GroupingNode {
	return GroupingNode{}
}

func (n NodeFactory) NewNamedFunctionNode() NamedFunctionNode {
	return NamedFunctionNode{}
}
