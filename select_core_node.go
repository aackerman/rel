package rel

type SelectCoreNode struct {
	Source       *JoinSource
	Top          *TopNode
	Projections  *[]Visitable
	SetQuanifier *Visitable
	Wheres       *[]Visitable
	Groups       *[]GroupNode
	Having       *HavingNode
	Windows      *[]WindowNode
	BaseVisitable
}

func NewSelectCoreNode() SelectCoreNode {
	return SelectCoreNode{Source: &JoinSource{
		Right: make([]Visitable, 0),
	}}
}
