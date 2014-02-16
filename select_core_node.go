package rel

type SelectCoreNode struct {
	Source        *JoinSource
	Top           *TopNode
	Selections    *[]Visitable
	SetQuantifier Visitable
	Wheres        *[]Visitable
	Groups        *[]Visitable
	Having        *HavingNode
	Windows       *[]Visitable
	BaseVisitable
}

func NewSelectCoreNode() *SelectCoreNode {
	return &SelectCoreNode{
		Source: &JoinSource{
			Right: []Visitable{},
		},
	}
}

func (node *SelectCoreNode) SetFrom(v Visitable) {
	node.Source.Left = v
}
