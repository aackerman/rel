package rel

type ExistsNode FunctionNode

func NewExistsNode(v Visitable) *ExistsNode {
	exists := new(ExistsNode)
	exists.Expressions = append(exists.Expressions, v)
	return exists
}

func (node ExistsNode) As(v Visitable) AsNode {
	return AsNode{
		Left:  node,
		Right: v,
	}
}
