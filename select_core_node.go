package arel

type SelectCoreNode struct {
	Source       *JoinSource
	Top          *TopNode
	Projections  *[]AstNode
	SetQuanifier *AstNode
	Wheres       *[]AstNode
	Groups       *[]AstNode
	Having       *HavingNode
	Windows      *[]AstNode
	AstNode
}

func CreateSelectCoreNode() *SelectCoreNode {
	return new(SelectCoreNode)
}
