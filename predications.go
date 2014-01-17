package rel

type Predicator interface {
	Eq(Visitable) *EqualityNode
	EqAny(...Visitable) *GroupingNode
	EqAll(...Visitable) *GroupingNode
	Lt(Visitable) *LessThanNode
	LtEq(Visitable) *LessThanOrEqualNode
	LtAny(...Visitable) *GroupingNode
	LtAll(...Visitable) *GroupingNode
	Gt(Visitable) *GreaterThanNode
	GtEq(Visitable) *GreaterThanOrEqualNode
	GtAny(...Visitable) *GroupingNode
	GtAll(...Visitable) *GroupingNode
	In([]Visitable) Visitable
	NotIn([]Visitable) Visitable
	NotEq(Visitable) *NotEqualNode
	Matches(SqlLiteralNode) *MatchesNode
	DoesNotMatch(SqlLiteralNode) *DoesNotMatchNode
	GroupAny(...Visitable) *GroupingNode
	GroupAll(...Visitable) *GroupingNode
	Visitable
}

func predicationEq(node Predicator, visitable Visitable) *EqualityNode {
	return &EqualityNode{Left: node, Right: visitable}
}

func predicationEqAny(node Predicator, visitables ...Visitable) *GroupingNode {
	for i, visitable := range visitables {
		visitables[i] = node.Eq(visitable)
	}
	return node.GroupAny(visitables...)
}

func predicationEqAll(node Predicator, visitable ...Visitable) *GroupingNode {
	for i, v := range visitable {
		visitable[i] = node.Eq(v)
	}
	return node.GroupAll(visitable...)
}

func predicationLt(node Predicator, visitable Visitable) *LessThanNode {
	return &LessThanNode{Left: node, Right: visitable}
}

func predicationLtAny(node Predicator, visitable ...Visitable) *GroupingNode {
	for i, v := range visitable {
		visitable[i] = node.Lt(v)
	}
	return node.GroupAny(visitable...)
}

func predicationLtEqAny(node Predicator, visitable ...Visitable) *GroupingNode {
	for i, v := range visitable {
		visitable[i] = node.LtEq(v)
	}
	return node.GroupAny(visitable...)
}

func predicationLtEqAll(node Predicator, visitable ...Visitable) *GroupingNode {
	for i, v := range visitable {
		visitable[i] = node.LtEq(v)
	}
	return node.GroupAll(visitable...)
}

func predicationLtAll(node Predicator, visitable ...Visitable) *GroupingNode {
	for i, v := range visitable {
		visitable[i] = node.Lt(v)
	}
	return node.GroupAll(visitable...)
}

func predicationLtEq(node Predicator, v Visitable) *LessThanOrEqualNode {
	return &LessThanOrEqualNode{Left: node, Right: v}
}

func predicationGt(node Predicator, v Visitable) *GreaterThanNode {
	return &GreaterThanNode{Left: node, Right: v}
}

func predicationGtAny(node Predicator, visitable ...Visitable) *GroupingNode {
	for i, v := range visitable {
		visitable[i] = node.Gt(v)
	}
	return node.GroupAny(visitable...)
}

func predicationGtEqAny(node Predicator, visitable ...Visitable) *GroupingNode {
	for i, v := range visitable {
		visitable[i] = node.GtEq(v)
	}
	return node.GroupAny(visitable...)
}

func predicationGtEqAll(node Predicator, visitable ...Visitable) *GroupingNode {
	for i, v := range visitable {
		visitable[i] = node.GtEq(v)
	}
	return node.GroupAll(visitable...)
}

func predicationGtAll(node Predicator, visitable ...Visitable) *GroupingNode {
	for i, v := range visitable {
		visitable[i] = node.Gt(v)
	}
	return node.GroupAll(visitable...)
}

func predicationGtEq(node Predicator, v Visitable) *GreaterThanOrEqualNode {
	return &GreaterThanOrEqualNode{Left: node, Right: v}
}

func predicationCount(node Predicator) *CountNode {
	return &CountNode{Expressions: []Visitable{node}}
}

func predicationExtract(node Predicator, literal SqlLiteralNode) *ExtractNode {
	return &ExtractNode{Expressions: []Visitable{node}, Field: &literal}
}

func predicationAs(node Predicator, v Visitable) *AsNode {
	return &AsNode{
		Left:  node,
		Right: v,
	}
}

func predicationIn(node Predicator, visitables []Visitable) Visitable {
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

func predicationNotIn(node Predicator, visitables []Visitable) Visitable {
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

func predicationNotInAny(node Predicator, visitableslices ...[]Visitable) Visitable {
	visitables := make([]Visitable, len(visitableslices))
	for i, visitableslice := range visitableslices {
		visitables[i] = node.NotIn(visitableslice)
	}
	return node.GroupAny(visitables...)
}

func predicationNotInAll(node Predicator, visitableslices ...[]Visitable) Visitable {
	visitables := make([]Visitable, len(visitableslices))
	for i, visitableslice := range visitableslices {
		visitables[i] = node.NotIn(visitableslice)
	}
	return node.GroupAll(visitables...)
}

func predicationInAny(node Predicator, visitableslices ...[]Visitable) Visitable {
	visitables := make([]Visitable, len(visitableslices))
	for i, visitableslice := range visitableslices {
		visitables[i] = node.In(visitableslice)
	}
	return node.GroupAny(visitables...)
}

func predicationInAll(node Predicator, visitableslices ...[]Visitable) Visitable {
	visitables := make([]Visitable, len(visitableslices))
	for i, visitableslice := range visitableslices {
		visitables[i] = node.In(visitableslice)
	}
	return node.GroupAll(visitables...)
}

func predicationNotEq(node Predicator, v Visitable) *NotEqualNode {
	return &NotEqualNode{
		Left:  node,
		Right: v,
	}
}

func predicationNotEqual(node Predicator, v Visitable) *NotEqualNode {
	return node.NotEq(v)
}

func predicationNotEqAny(node Predicator, visitable ...Visitable) *GroupingNode {
	for i, v := range visitable {
		visitable[i] = node.NotEq(v)
	}
	return node.GroupAny(visitable...)
}

func predicationNotEqAll(node Predicator, visitable ...Visitable) *GroupingNode {
	for i, v := range visitable {
		visitable[i] = node.NotEq(v)
	}
	return node.GroupAll(visitable...)
}

func predicationDoesNotMatch(node Predicator, literal SqlLiteralNode) *DoesNotMatchNode {
	var v Visitable = &QuotedNode{Raw: literal.Raw}
	return &DoesNotMatchNode{
		Left:  node,
		Right: v,
	}
}

func predicationDoesNotMatchAny(node Predicator, literals ...SqlLiteralNode) *GroupingNode {
	visitables := make([]Visitable, len(literals))
	for i, literal := range literals {
		visitables[i] = node.DoesNotMatch(literal)
	}
	return node.GroupAny(visitables...)
}

func predicationDoesNotMatchAll(node Predicator, literals ...SqlLiteralNode) *GroupingNode {
	visitables := make([]Visitable, len(literals))
	for i, literal := range literals {
		visitables[i] = node.DoesNotMatch(literal)
	}
	return node.GroupAll(visitables...)
}

func predicationMatches(node Predicator, literal SqlLiteralNode) *MatchesNode {
	var v Visitable = &QuotedNode{Raw: literal.Raw}
	return &MatchesNode{
		Left:  node,
		Right: v,
	}
}

func predicationMatchesAny(node Predicator, literals ...SqlLiteralNode) *GroupingNode {
	visitables := make([]Visitable, len(literals))
	for i, literal := range literals {
		visitables[i] = node.Matches(literal)
	}
	return node.GroupAny(visitables...)
}

func predicationMatchesAll(node Predicator, literals ...SqlLiteralNode) *GroupingNode {
	visitables := make([]Visitable, len(literals))
	for i, literal := range literals {
		visitables[i] = node.Matches(literal)
	}
	return node.GroupAll(visitables...)
}

func predicationGroupAny(node Predicator, visitable ...Visitable) *GroupingNode {
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

func predicationGroupAll(node Predicator, visitable ...Visitable) *GroupingNode {
	return &GroupingNode{
		Expr: []Visitable{
			&AndNode{Children: &visitable},
		},
	}
}
