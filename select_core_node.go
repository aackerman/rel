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

func CreateSelectCoreNode(t *Table) *SelectCoreNode {
	core := SelectCoreNode{Source: new(JoinSource)}
	return &core
}
