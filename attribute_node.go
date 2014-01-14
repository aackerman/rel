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

func (node AttributeNode) Eq(visitable Visitable) *EqualityNode {
	return &EqualityNode{Left: node, Right: visitable}
}

func (node AttributeNode) EqAny(visitable ...Visitable) *GroupingNode {
	for i, v := range visitable {
		visitable[i] = node.Eq(v)
	}
	return node.GroupAny(visitable...)
}

func (node AttributeNode) EqAll(visitable ...Visitable) *GroupingNode {
	for i, v := range visitable {
		visitable[i] = node.Eq(v)
	}
	return node.GroupAll(visitable...)
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

func (node AttributeNode) Desc() *DescendingNode {
	return &DescendingNode{Expr: node}
}

func (node AttributeNode) Asc() *AscendingNode {
	return &AscendingNode{Expr: node}
}

func (node AttributeNode) Count() *CountNode {
	return &CountNode{Expressions: []Visitable{node}}
}

func (node AttributeNode) As(v Visitable) *AsNode {
	return &AsNode{
		Left:  node,
		Right: v,
	}
}

func (node AttributeNode) In(visitables []Visitable) Visitable {
	in := &InNode{Left: node}
	for _, v := range visitables {
		switch val := v.(type) {
		case SelectManager:
			in.Right = append(in.Right, val.Ast)
		default:
			in.Right = append(in.Right, v)
		}
	}
	return in
}

func (node AttributeNode) NotIn(visitables []Visitable) Visitable {
	notin := &NotInNode{Left: node}
	for _, v := range visitables {
		switch val := v.(type) {
		case SelectManager:
			notin.Right = append(notin.Right, val.Ast)
		default:
			notin.Right = append(notin.Right, v)
		}
	}
	return notin
}

func (node AttributeNode) NotInAny(visitableslices ...[]Visitable) Visitable {
	visitables := make([]Visitable, len(visitableslices))
	for i, visitableslice := range visitableslices {
		visitables[i] = node.NotIn(visitableslice)
	}
	return node.GroupAny(visitables...)
}

func (node AttributeNode) NotInAll(visitableslices ...[]Visitable) Visitable {
	visitables := make([]Visitable, len(visitableslices))
	for i, visitableslice := range visitableslices {
		visitables[i] = node.NotIn(visitableslice)
	}
	return node.GroupAll(visitables...)
}

func (node AttributeNode) InAny(visitableslices ...[]Visitable) Visitable {
	visitables := make([]Visitable, len(visitableslices))
	for i, visitableslice := range visitableslices {
		visitables[i] = node.In(visitableslice)
	}
	return node.GroupAny(visitables...)
}

func (node AttributeNode) InAll(visitableslices ...[]Visitable) Visitable {
	visitables := make([]Visitable, len(visitableslices))
	for i, visitableslice := range visitableslices {
		visitables[i] = node.In(visitableslice)
	}
	return node.GroupAll(visitables...)
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

func (node AttributeNode) DoesNotMatch(literal SqlLiteralNode) *DoesNotMatchNode {
	var v Visitable = &QuotedNode{Raw: literal.Raw}
	return &DoesNotMatchNode{
		Left:  node,
		Right: v,
	}
}

func (node AttributeNode) DoesNotMatchAny(literals ...SqlLiteralNode) *GroupingNode {
	visitables := make([]Visitable, len(literals))
	for i, literal := range literals {
		visitables[i] = node.DoesNotMatch(literal)
	}
	return node.GroupAny(visitables...)
}

func (node AttributeNode) DoesNotMatchAll(literals ...SqlLiteralNode) *GroupingNode {
	visitables := make([]Visitable, len(literals))
	for i, literal := range literals {
		visitables[i] = node.DoesNotMatch(literal)
	}
	return node.GroupAll(visitables...)
}

func (node AttributeNode) Matches(literal SqlLiteralNode) *MatchesNode {
	var v Visitable = &QuotedNode{Raw: literal.Raw}
	return &MatchesNode{
		Left:  node,
		Right: v,
	}
}

func (node AttributeNode) MatchesAny(literals ...SqlLiteralNode) *GroupingNode {
	visitables := make([]Visitable, len(literals))
	for i, literal := range literals {
		visitables[i] = node.Matches(literal)
	}
	return node.GroupAny(visitables...)
}

func (node AttributeNode) MatchesAll(literals ...SqlLiteralNode) *GroupingNode {
	visitables := make([]Visitable, len(literals))
	for i, literal := range literals {
		visitables[i] = node.Matches(literal)
	}
	return node.GroupAll(visitables...)
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
	return &GroupingNode{
		Expr: []Visitable{
			&AndNode{Children: &visitable},
		},
	}
}
