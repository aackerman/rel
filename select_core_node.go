package rel

type SelectCoreNode struct {
	Source       *JoinSource
	Top          *TopNode
	Projections  *[]AstNode
	SetQuanifier *AstNode
	Wheres       *[]AstNode
	Groups       *[]GroupNode
	Having       *HavingNode
	Windows      *[]WindowNode
	AstNode
}

func NewSelectCoreNode() SelectCoreNode {
	return SelectCoreNode{Source: new(JoinSource)}
}
