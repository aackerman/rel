package rel

type ExtractNode struct {
	Expressions []Visitable
	Field       *SqlLiteralNode
	Alias       *SqlLiteralNode
	BaseVisitable
}

func (node *ExtractNode) As(n SqlLiteralNode) *ExtractNode {
	node.Alias = &n
	return node
}
