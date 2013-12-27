package arel

type SelectCoreNode struct {
	Source       *JoinSource
	Top          *TopNode
	Projections  *[]AstNode
	SetQuanifier *AstNode
	Wheres       *[]AstNode
	Groups       *[]GroupingNode
	Having       *HavingNode
	Windows      *[]WindowNode
	AstNode
}

func NewSelectCoreNode() SelectCoreNode {
	return SelectCoreNode{Source: new(JoinSource)}
}
