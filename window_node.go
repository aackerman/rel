package rel

type WindowNode struct {
	Orders  *[]Visitable
	Framing Visitable
	BaseVisitable
}

func (node *WindowNode) Order(v Visitable) *WindowNode {
	if node.Orders == nil {
		node.Orders = &[]Visitable{}
	}
	*node.Orders = append(*node.Orders, v)
	return node
}

type NamedWindowNode struct {
	Name    SqlLiteralNode
	Orders  *[]Visitable
	Framing Visitable
	BaseVisitable
}

func (node *NamedWindowNode) Order(v Visitable) *NamedWindowNode {
	if node.Orders == nil {
		node.Orders = &[]Visitable{}
	}
	*node.Orders = append(*node.Orders, v)
	return node
}

func (node *NamedWindowNode) Rows(v Visitable) Visitable {
	return node.Frame(&RowsNode{Expr: v})
}

func (node *NamedWindowNode) Frame(v Visitable) Visitable {
	node.Framing = v
	return node.Framing
}

func (node *NamedWindowNode) Range(v Visitable) Visitable {
	return node.Frame(&RangeNode{Expr: v})
}
