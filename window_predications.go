package rel

type Windower interface {
	Over(Visitable) *OverNode
	Visitable
}

func windowPredicationOver(left Windower, right Visitable) *OverNode {
	return &OverNode{
		Left:  left,
		Right: right,
	}
}
