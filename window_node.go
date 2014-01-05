package rel

type WindowNode struct {
	BaseVisitable
}

type NamedWindowNode struct {
	Name    SqlLiteralNode
	Orders  *[]Visitable
	Framing Visitable
	BaseVisitable
}

func (node *NamedWindowNode) Order(v Visitable) {
	if node.Orders == nil {
		slice := make([]Visitable, 0)
		node.Orders = &slice
	}
	*node.Orders = append(*node.Orders, v)
}

func (node *NamedWindowNode) Rows(v Visitable) {
	node.Frame(&RowsNode{Expr: v})
}

func (node *NamedWindowNode) Frame(v Visitable) {
	node.Framing = v
}
