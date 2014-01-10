package rel

type ExtractNode struct {
	Expr  Visitable
	Field Visitable
	Alias *SqlLiteralNode
	BaseVisitable
}

func (node *ExtractNode) As(n SqlLiteralNode) *ExtractNode {
	node.Alias = &n
	return node
}
