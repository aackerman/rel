package rel

type AttributeNode struct {
	Name     SqlLiteralNode
	Relation Visitable
	BaseVisitable
}

func NewAttributeNode(v Visitable, name string) AttributeNode {
	return AttributeNode{
		Name:     Sql(name),
		Relation: v,
	}
}

func (a AttributeNode) Eq(n Visitable) EqualityNode {
	return NewEqualityNode(a, n)
}

func (a AttributeNode) Lt(i int) LessThanNode {
	return LessThanNode{Left: a, Right: Sql(i)}
}

func (a AttributeNode) Gt(i int) GreaterThanNode {
	return GreaterThanNode{Left: a, Right: Sql(i)}
}

func (a AttributeNode) Desc() DescendingNode {
	return DescendingNode{Expr: a}
}

func (a AttributeNode) Asc() AscendingNode {
	return AscendingNode{Expr: a}
}

func (a AttributeNode) Count() CountNode {
	return CountNode{Expressions: a}
}

func (node AttributeNode) As(v Visitable) AsNode {
	return AsNode{
		Left:  node,
		Right: v,
	}
}

func (node AttributeNode) In(v Visitable) Visitable {
	var ret Visitable
	switch val := v.(type) {
	case SelectManager:
		ret = &InNode{Left: node, Right: val.Ast}
	default:
		ret = &InNode{Left: node, Right: v}
	}
	return ret
}

func (node AttributeNode) NotEq(v Visitable) *NotEqualNode {
	return &NotEqualNode{
		Left:  node,
		Right: v,
	}
}

func (node AttributeNode) NotEqual(v Visitable) *NotEqualNode {
	return node.NotEq(v)
}

func (node AttributeNode) NotEqAny(n Visitable) GroupingNode {
	return GroupingNode{}
}

func (a AttributeNode) NotEqAll(n Visitable) GroupingNode {
	return GroupingNode{}
}
