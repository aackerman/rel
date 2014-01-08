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

func (node AttributeNode) Eq(v Visitable) EqualityNode {
	return NewEqualityNode(node, v)
}

func (node AttributeNode) Lt(v Visitable) LessThanNode {
	return LessThanNode{Left: node, Right: v}
}

func (node AttributeNode) LtEq(v Visitable) *LessThanOrEqualNode {
	return &LessThanOrEqualNode{Left: node, Right: v}
}

func (node AttributeNode) Gt(v Visitable) GreaterThanNode {
	return GreaterThanNode{Left: node, Right: v}
}

func (node AttributeNode) GtEq(v Visitable) *GreaterThanOrEqualNode {
	return &GreaterThanOrEqualNode{Left: node, Right: v}
}

func (node AttributeNode) Desc() DescendingNode {
	return DescendingNode{Expr: node}
}

func (node AttributeNode) Asc() AscendingNode {
	return AscendingNode{Expr: node}
}

func (node AttributeNode) Count() CountNode {
	return CountNode{Expressions: node}
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

func (node AttributeNode) NotEqAny(visitable ...Visitable) *GroupingNode {
	var nodes []*NotEqualNode
	grouping := new(GroupingNode)
	for _, v := range visitable {
		nodes = append(nodes, node.NotEq(v))
	}
	if len(nodes) > 0 {
		// unshift first node
		m, nodes := nodes[0], nodes[1:]
		var memo Visitable = m
		for _, n := range nodes {
			memo = &OrNode{Left: memo, Right: n}
		}
		grouping.Expr = append(grouping.Expr, memo)
	}
	return grouping
}

func (node AttributeNode) NotEqAll(visitable ...Visitable) *GroupingNode {
	var nodes []Visitable
	grouping := new(GroupingNode)
	for _, v := range visitable {
		nodes = append(nodes, node.NotEq(v))
	}
	grouping.Expr = append(grouping.Expr, &AndNode{Children: &nodes})
	return grouping
}
