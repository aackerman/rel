package arel

type SelectCoreNode struct {
	Source       JoinSource
	Top          TopNode
	Projections  []AstNode
	SetQuanifier AstNode
	Wheres       []AstNode
	Groups       []GroupNode
	Having       HavingNode
	Windows      []AstNode
	AstNode
}

func CreateSelectCoreNode() SelectCoreNode {
	return SelectCoreNode{
		Source:      JoinSource{BinaryNode{}},
		Projections: make([]AstNode, 10),
		Wheres:      make([]AstNode, 10),
		Windows:     make([]AstNode, 10),
		Groups:      make([]GroupNode, 10),
	}
}
