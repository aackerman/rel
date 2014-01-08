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

func (node AttributeNode) Lt(v Visitable) *LessThanNode {
	return &LessThanNode{Left: node, Right: v}
}

func (node AttributeNode) LtAny(visitable ...Visitable) *GroupingNode {
	for i, v := range visitable {
		visitable[i] = node.Lt(v)
	}
	return node.GroupAny(visitable...)
}

func (node AttributeNode) LtEqAny(visitable ...Visitable) *GroupingNode {
	for i, v := range visitable {
		visitable[i] = node.LtEq(v)
	}
	return node.GroupAny(visitable...)
}

func (node AttributeNode) LtEqAll(visitable ...Visitable) *GroupingNode {
	for i, v := range visitable {
		visitable[i] = node.LtEq(v)
	}
	return node.GroupAll(visitable...)
}

func (node AttributeNode) LtAll(visitable ...Visitable) *GroupingNode {
	for i, v := range visitable {
		visitable[i] = node.Lt(v)
	}
	return node.GroupAll(visitable...)
}

func (node AttributeNode) LtEq(v Visitable) *LessThanOrEqualNode {
	return &LessThanOrEqualNode{Left: node, Right: v}
}

func (node AttributeNode) Gt(v Visitable) *GreaterThanNode {
	return &GreaterThanNode{Left: node, Right: v}
}

func (node AttributeNode) GtAny(visitable ...Visitable) *GroupingNode {
	for i, v := range visitable {
		visitable[i] = node.Gt(v)
	}
	return node.GroupAny(visitable...)
}

func (node AttributeNode) GtEqAny(visitable ...Visitable) *GroupingNode {
	for i, v := range visitable {
		visitable[i] = node.GtEq(v)
	}
	return node.GroupAny(visitable...)
}

func (node AttributeNode) GtEqAll(visitable ...Visitable) *GroupingNode {
	for i, v := range visitable {
		visitable[i] = node.GtEq(v)
	}
	return node.GroupAll(visitable...)
}

func (node AttributeNode) GtAll(visitable ...Visitable) *GroupingNode {
	for i, v := range visitable {
		visitable[i] = node.Gt(v)
	}
	return node.GroupAll(visitable...)
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
	for i, v := range visitable {
		visitable[i] = node.NotEq(v)
	}
	return node.GroupAny(visitable...)
}

func (node AttributeNode) NotEqAll(visitable ...Visitable) *GroupingNode {
	for i, v := range visitable {
		visitable[i] = node.NotEq(v)
	}
	return node.GroupAll(visitable...)
}

func (node AttributeNode) GroupAny(visitable ...Visitable) *GroupingNode {
	grouping := new(GroupingNode)
	if len(visitable) > 0 {
		// unshift first node
		m, visitable := visitable[0], visitable[1:]
		var memo Visitable = m
		for _, n := range visitable {
			memo = &OrNode{Left: memo, Right: n}
		}
		grouping.Expr = append(grouping.Expr, memo)
	}
	return grouping
}

func (node AttributeNode) GroupAll(visitable ...Visitable) *GroupingNode {
	grouping := new(GroupingNode)
	grouping.Expr = append(grouping.Expr, &AndNode{Children: &visitable})
	return grouping
}
