package arel

type SelectCore struct {
	Source       JoinSource
	Top          int
	Projections  []Node
	SetQuanifier interface{}
	Wheres       []Node
	Groups       []GroupingNode
	Having       interface{}
	Windows      []Node
}

func (s SelectCore) NodeInterface() {}
