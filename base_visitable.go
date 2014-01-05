package rel

import (
	"log"
)

// BaseVisitable satisfies the Visitable Interface
// All other nodes should have an embedded BaseVisitable
type BaseVisitable struct{}

func (a BaseVisitable) NotEq(n Visitable) NotEqualNode {
	log.Fatal("BaseVisitable#NotEq not implemented")
	return NotEqualNode{}
}

func (a BaseVisitable) NotEqAny(n Visitable) GroupingNode {
	log.Fatal("BaseVisitable#NotEqAny not implemented")
	return GroupingNode{}
}

func (a BaseVisitable) NotEqAll(n Visitable) GroupingNode {
	log.Fatal("BaseVisitable#NotEqAll not implemented")
	return GroupingNode{}
}

func (a BaseVisitable) NewTrueNode() TrueNode {
	return TrueNode{}
}

func (a BaseVisitable) NewFalseNode() FalseNode {
	return FalseNode{}
}

func (a BaseVisitable) NewTableAliasNode(t *Table, name SqlLiteralNode) *TableAliasNode {
	return &TableAliasNode{Relation: t, Name: name}
}

func (a BaseVisitable) NewStringJoinNode() StringJoinNode {
	return StringJoinNode{}
}

func (a BaseVisitable) NewInnerJoinNode() InnerJoinNode {
	return InnerJoinNode{}
}

func (a BaseVisitable) NewOuterJoinNode() OuterJoinNode {
	return OuterJoinNode{}
}

func (a BaseVisitable) NewAndNode(n ...Visitable) AndNode {
	return AndNode{Children: &n}
}

func (a BaseVisitable) NewOnNode() OnNode {
	return OnNode{}
}

func (a BaseVisitable) NewNotNode() NotNode {
	return NotNode{}
}

func (a BaseVisitable) NewGroupingNode() GroupingNode {
	return GroupingNode{}
}

func (a BaseVisitable) NewNamedFunctionNode() NamedFunctionNode {
	return NamedFunctionNode{}
}
