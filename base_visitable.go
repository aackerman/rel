package rel

// BaseVisitable satisfies the Visitable Interface
// All other nodes should have an embedded BaseVisitable
type BaseVisitable struct{}

func (v BaseVisitable) String() string {
	return ""
}

func (v BaseVisitable) NewTrueNode() TrueNode {
	return TrueNode{}
}

func (v BaseVisitable) NewFalseNode() FalseNode {
	return FalseNode{}
}

func (v BaseVisitable) NewTableAliasNode(t *Table, name SqlLiteralNode) *TableAliasNode {
	return &TableAliasNode{Relation: t, Name: name}
}

func (v BaseVisitable) NewStringJoinNode() StringJoinNode {
	return StringJoinNode{}
}

func (v BaseVisitable) NewInnerJoinNode() InnerJoinNode {
	return InnerJoinNode{}
}

func (v BaseVisitable) NewOuterJoinNode() OuterJoinNode {
	return OuterJoinNode{}
}

func (v BaseVisitable) NewAndNode(n ...Visitable) *AndNode {
	return &AndNode{Children: &n}
}

func (v BaseVisitable) NewOnNode(visitable Visitable) *OnNode {
	return &OnNode{Expr: visitable}
}

func (v BaseVisitable) NewNotNode() NotNode {
	return NotNode{}
}

func (v BaseVisitable) NewGroupingNode() GroupingNode {
	return GroupingNode{}
}

func (v BaseVisitable) NewNamedFunctionNode() NamedFunctionNode {
	return NamedFunctionNode{}
}
