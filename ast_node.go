package grel

type AstNode interface {
	NewTrueNode() TrueNode
	NewFalseNode() FalseNode
	NewTableAliasNode(*Table, string) TableAliasNode
	NewStringJoinNode() StringJoinNode
	NewInnerJoinNode() InnerJoinNode
	NewOuterJoinNode() OuterJoinNode
	NewAndNode(...AstNode) AndNode
	NewOnNode() OnNode
	NewNotNode() NotNode
	NewGroupingNode() GroupingNode
}

type BaseNode struct {
	NodeFactory
}

func NewBaseNode() BaseNode {
	return BaseNode{}
}
