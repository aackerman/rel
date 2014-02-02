package rel

type ExistsNode FunctionNode

func NewExistsNode(v Visitable) *ExistsNode {
	return &ExistsNode{
		Expressions: []Visitable{v},
	}
}

func (node ExistsNode) Desc() *DescendingNode {
	return orderingDesc(node)
}

func (node ExistsNode) Asc() *AscendingNode {
	return orderingAsc(node)
}

func (node ExistsNode) Eq(visitable Visitable) *EqualityNode {
	return predicationEq(node, visitable)
}

func (node ExistsNode) EqAny(visitables ...Visitable) *GroupingNode {
	return predicationEqAny(node, visitables...)
}

func (node ExistsNode) EqAll(visitables ...Visitable) *GroupingNode {
	return predicationEqAll(node, visitables...)
}

func (node ExistsNode) Lt(visitable Visitable) *LessThanNode {
	return predicationLt(node, visitable)
}

func (node ExistsNode) LtAny(visitables ...Visitable) *GroupingNode {
	return predicationLtAny(node, visitables...)
}

func (node ExistsNode) LtAll(visitables ...Visitable) *GroupingNode {
	return predicationLtAll(node, visitables...)
}

func (node ExistsNode) LtEq(visitable Visitable) *LessThanOrEqualNode {
	return predicationLtEq(node, visitable)
}

func (node ExistsNode) LtEqAny(visitables ...Visitable) *GroupingNode {
	return predicationLtEqAny(node, visitables...)
}

func (node ExistsNode) LtEqAll(visitables ...Visitable) *GroupingNode {
	return predicationLtEqAll(node, visitables...)
}

func (node ExistsNode) Gt(visitable Visitable) *GreaterThanNode {
	return predicationGt(node, visitable)
}

func (node ExistsNode) GtAny(visitables ...Visitable) *GroupingNode {
	return predicationGtAny(node, visitables...)
}

func (node ExistsNode) GtAll(visitables ...Visitable) *GroupingNode {
	return predicationGtAll(node, visitables...)
}

func (node ExistsNode) GtEq(visitable Visitable) *GreaterThanOrEqualNode {
	return predicationGtEq(node, visitable)
}

func (node ExistsNode) GtEqAny(visitables ...Visitable) *GroupingNode {
	return predicationGtEqAny(node, visitables...)
}

func (node ExistsNode) GtEqAll(visitables ...Visitable) *GroupingNode {
	return predicationGtEqAll(node, visitables...)
}

func (node ExistsNode) Count() *CountNode {
	return predicationCount(node)
}

func (node ExistsNode) Extract(literal SqlLiteralNode) *ExtractNode {
	return predicationExtract(node, literal)
}

func (node ExistsNode) As(literal SqlLiteralNode) *AsNode {
	return aliasPredicationAs(node, literal)
}

func (node ExistsNode) In(visitables []Visitable) Visitable {
	return predicationIn(node, visitables)
}

func (node ExistsNode) InAny(visitableslices ...[]Visitable) Visitable {
	return predicationInAny(node, visitableslices...)
}

func (node ExistsNode) InAll(visitableslices ...[]Visitable) Visitable {
	return predicationInAll(node, visitableslices...)
}

func (node ExistsNode) NotIn(visitables []Visitable) Visitable {
	return predicationNotIn(node, visitables)
}

func (node ExistsNode) NotInAny(visitableslices ...[]Visitable) Visitable {
	return predicationNotInAny(node, visitableslices...)
}

func (node ExistsNode) NotInAll(visitableslices ...[]Visitable) Visitable {
	return predicationNotInAll(node, visitableslices...)
}

func (node ExistsNode) NotEq(visitable Visitable) *NotEqualNode {
	return predicationNotEq(node, visitable)
}

func (node ExistsNode) NotEqAny(visitables ...Visitable) *GroupingNode {
	return predicationNotEqAny(node, visitables...)
}

func (node ExistsNode) NotEqAll(visitables ...Visitable) *GroupingNode {
	return predicationNotEqAll(node, visitables...)
}

func (node ExistsNode) DoesNotMatch(literal SqlLiteralNode) *DoesNotMatchNode {
	return predicationDoesNotMatch(node, literal)
}

func (node ExistsNode) DoesNotMatchAny(literals ...SqlLiteralNode) *GroupingNode {
	return predicationDoesNotMatchAny(node, literals...)
}

func (node ExistsNode) DoesNotMatchAll(literals ...SqlLiteralNode) *GroupingNode {
	return predicationDoesNotMatchAll(node, literals...)
}

func (node ExistsNode) Matches(literal SqlLiteralNode) *MatchesNode {
	return predicationMatches(node, literal)
}

func (node ExistsNode) MatchesAny(literals ...SqlLiteralNode) *GroupingNode {
	return predicationMatchesAny(node, literals...)
}

func (node ExistsNode) MatchesAll(literals ...SqlLiteralNode) *GroupingNode {
	return predicationMatchesAll(node, literals...)
}
