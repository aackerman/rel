package rel

type ExistsNode FunctionNode

func NewExistsNode(v Visitable) *ExistsNode {
	return &ExistsNode{
		Expressions: []Visitable{v},
	}
}

func (node ExistsNode) As(v Visitable) *AsNode {
	return &AsNode{
		Left:  node,
		Right: v,
	}
}
