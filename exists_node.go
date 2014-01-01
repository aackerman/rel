package rel

type ExistsNode FunctionNode

func NewExistsNode(n AstNode) ExistsNode {
	return ExistsNode{Expressions: n}
}

func (e ExistsNode) As(n AstNode) AsNode {
	return AsNode{
		Left:  e,
		Right: &n,
	}
}
