package rel

type Windower interface {
	Over(Visitable) *OverNode
	Visitable
}

func windowPredicationOver(node Windower, visitable Visitable) *OverNode {
	return &OverNode{
		Left:  node,
		Right: visitable,
	}
}
