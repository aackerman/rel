package arel

type SelectCore struct {
	Source       JoinSource
	Top          int
	Projections  []AstNode
	SetQuanifier interface{}
	Wheres       []AstNode
	Groups       []GroupNode
	Having       interface{}
	Windows      []AstNode
	AstNode
}

func CreateSelectCore() SelectCore {
	return SelectCore{
		Source:      JoinSource{BinaryNode{}},
		Projections: make([]AstNode, 10),
		Wheres:      make([]AstNode, 10),
		Windows:     make([]AstNode, 10),
		Groups:      make([]GroupNode, 10),
	}
}
