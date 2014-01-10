package rel

type EqualityNode struct {
	Left  Visitable
	Right Visitable
	BaseVisitable
}

func NewEqualityNode(left Visitable, right Visitable) EqualityNode {
	return EqualityNode{
		Left:  left,
		Right: right,
	}
}

func (node EqualityNode) Or(other Visitable) *GroupingNode {
	return &GroupingNode{Expr: []Visitable{&OrNode{Left: node, Right: other}}}
}

func (node EqualityNode) And(other Visitable) *GroupingNode {
	return &GroupingNode{
		Expr: []Visitable{
			&AndNode{
				Children: &[]Visitable{node, other},
			},
		},
	}
}

type InNode struct {
	Left  Visitable
	Right []Visitable
	BaseVisitable
}

type NotInNode struct {
	Left  Visitable
	Right []Visitable
	BaseVisitable
}
