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

func (node *ExtractNode) Desc() *DescendingNode {
	return orderingDesc(node)
}

func (node *ExtractNode) Asc() *AscendingNode {
	return orderingAsc(node)
}

func (node *ExtractNode) Eq(visitable Visitable) *EqualityNode {
	return predicationEq(node, visitable)
}

func (node *ExtractNode) EqAny(visitables ...Visitable) *GroupingNode {
	return predicationEqAny(node, visitables...)
}

func (node *ExtractNode) EqAll(visitables ...Visitable) *GroupingNode {
	return predicationEqAll(node, visitables...)
}

func (node *ExtractNode) Lt(visitable Visitable) *LessThanNode {
	return predicationLt(node, visitable)
}

func (node *ExtractNode) LtAny(visitables ...Visitable) *GroupingNode {
	return predicationLtAny(node, visitables...)
}

func (node *ExtractNode) LtAll(visitables ...Visitable) *GroupingNode {
	return predicationLtAll(node, visitables...)
}

func (node *ExtractNode) LtEq(visitable Visitable) *LessThanOrEqualNode {
	return predicationLtEq(node, visitable)
}

func (node *ExtractNode) LtEqAny(visitables ...Visitable) *GroupingNode {
	return predicationLtEqAny(node, visitables...)
}

func (node *ExtractNode) LtEqAll(visitables ...Visitable) *GroupingNode {
	return predicationLtEqAll(node, visitables...)
}

func (node *ExtractNode) Gt(visitable Visitable) *GreaterThanNode {
	return predicationGt(node, visitable)
}

func (node *ExtractNode) GtAny(visitables ...Visitable) *GroupingNode {
	return predicationGtAny(node, visitables...)
}

func (node *ExtractNode) GtAll(visitables ...Visitable) *GroupingNode {
	return predicationGtAll(node, visitables...)
}

func (node *ExtractNode) GtEq(visitable Visitable) *GreaterThanOrEqualNode {
	return predicationGtEq(node, visitable)
}

func (node *ExtractNode) GtEqAny(visitables ...Visitable) *GroupingNode {
	return predicationGtEqAny(node, visitables...)
}

func (node *ExtractNode) GtEqAll(visitables ...Visitable) *GroupingNode {
	return predicationGtEqAll(node, visitables...)
}

func (node *ExtractNode) Count() *CountNode {
	return predicationCount(node)
}

func (node *ExtractNode) Extract(literal SqlLiteralNode) *ExtractNode {
	return predicationExtract(node, literal)
}

func (node *ExtractNode) In(visitables []Visitable) Visitable {
	return predicationIn(node, visitables)
}

func (node *ExtractNode) InAny(visitableslices ...[]Visitable) Visitable {
	return predicationInAny(node, visitableslices...)
}

func (node *ExtractNode) InAll(visitableslices ...[]Visitable) Visitable {
	return predicationInAll(node, visitableslices...)
}

func (node *ExtractNode) NotIn(visitables []Visitable) Visitable {
	return predicationNotIn(node, visitables)
}

func (node *ExtractNode) NotInAny(visitableslices ...[]Visitable) Visitable {
	return predicationNotInAny(node, visitableslices...)
}

func (node *ExtractNode) NotInAll(visitableslices ...[]Visitable) Visitable {
	return predicationNotInAll(node, visitableslices...)
}

func (node *ExtractNode) NotEq(visitable Visitable) *NotEqualNode {
	return predicationNotEq(node, visitable)
}

func (node *ExtractNode) NotEqAny(visitables ...Visitable) *GroupingNode {
	return predicationNotEqAny(node, visitables...)
}

func (node *ExtractNode) NotEqAll(visitables ...Visitable) *GroupingNode {
	return predicationNotEqAll(node, visitables...)
}

func (node *ExtractNode) DoesNotMatch(literal SqlLiteralNode) *DoesNotMatchNode {
	return predicationDoesNotMatch(node, literal)
}

func (node *ExtractNode) DoesNotMatchAny(literals ...SqlLiteralNode) *GroupingNode {
	return predicationDoesNotMatchAny(node, literals...)
}

func (node *ExtractNode) DoesNotMatchAll(literals ...SqlLiteralNode) *GroupingNode {
	return predicationDoesNotMatchAll(node, literals...)
}

func (node *ExtractNode) Matches(literal SqlLiteralNode) *MatchesNode {
	return predicationMatches(node, literal)
}

func (node *ExtractNode) MatchesAny(literals ...SqlLiteralNode) *GroupingNode {
	return predicationMatchesAny(node, literals...)
}

func (node *ExtractNode) MatchesAll(literals ...SqlLiteralNode) *GroupingNode {
	return predicationMatchesAll(node, literals...)
}
