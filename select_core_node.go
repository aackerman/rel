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

func CreateSelectCoreNode(t *Table) *SelectCoreNode {
	core := SelectCoreNode{Source: new(JoinSource)}
	return &core
}
