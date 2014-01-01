package rel

type ExistsNode FunctionNode

func NewExistsNode(n Visitable) ExistsNode {
	return ExistsNode{Expressions: n}
}

func (e ExistsNode) As(n Visitable) AsNode {
	return AsNode{
		Left:  e,
		Right: &n,
	}
}
