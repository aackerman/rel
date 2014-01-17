package rel

type UnqualifiedColumnNode struct {
	Expr *AttributeNode
	BaseVisitable
}

func (node *UnqualifiedColumnNode) Name() SqlLiteralNode {
	return node.Expr.Name
}
