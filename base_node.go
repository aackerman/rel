package rel

import (
	"log"
)

// BaseNode satisfies the AstNode Interface
// All other nodes should have an embedded BaseNode
type BaseNode struct{}

func (a BaseNode) NotEq(n AstNode) NotEqualNode {
	log.Fatal("BaseNode#NotEq not implemented")
	return NotEqualNode{}
}

func (a BaseNode) NotEqAny(n AstNode) GroupingNode {
	log.Fatal("BaseNode#NotEqAny not implemented")
	return GroupingNode{}
}

func (a BaseNode) NotEqAll(n AstNode) GroupingNode {
	log.Fatal("BaseNode#NotEqAll not implemented")
	return GroupingNode{}
}

func (a BaseNode) NewTrueNode() TrueNode {
	return TrueNode{}
}

func (a BaseNode) NewFalseNode() FalseNode {
	return FalseNode{}
}

func (a BaseNode) NewTableAliasNode(t *Table, name string) TableAliasNode {
	return TableAliasNode{Name: name, Table: t}
}

func (a BaseNode) NewStringJoinNode() StringJoinNode {
	return StringJoinNode{}
}

func (a BaseNode) NewInnerJoinNode() InnerJoinNode {
	return InnerJoinNode{}
}

func (a BaseNode) NewOuterJoinNode() OuterJoinNode {
	return OuterJoinNode{}
}

func (a BaseNode) NewAndNode(n ...AstNode) AndNode {
	return AndNode{Children: &n}
}

func (a BaseNode) NewOnNode() OnNode {
	return OnNode{}
}

func (a BaseNode) NewNotNode() NotNode {
	return NotNode{}
}

func (a BaseNode) NewGroupingNode() GroupingNode {
	return GroupingNode{}
}

func (a BaseNode) NewNamedFunctionNode() NamedFunctionNode {
	return NamedFunctionNode{}
}
