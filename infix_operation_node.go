package rel

type InfixOperationNode struct {
	Operator SqlLiteralNode
	Left     Visitable
	Right    Visitable
	BaseVisitable
}

func (node InfixOperationNode) Desc() *DescendingNode {
	return orderingDesc(node)
}

func (node InfixOperationNode) Asc() *AscendingNode {
	return orderingAsc(node)
}

func (node InfixOperationNode) Eq(visitable Visitable) *EqualityNode {
	return predicationEq(node, visitable)
}

func (node InfixOperationNode) EqAny(visitables ...Visitable) *GroupingNode {
	return predicationEqAny(node, visitables...)
}

func (node InfixOperationNode) EqAll(visitables ...Visitable) *GroupingNode {
	return predicationEqAll(node, visitables...)
}

func (node InfixOperationNode) Lt(visitable Visitable) *LessThanNode {
	return predicationLt(node, visitable)
}

func (node InfixOperationNode) LtAny(visitables ...Visitable) *GroupingNode {
	return predicationLtAny(node, visitables...)
}

func (node InfixOperationNode) LtAll(visitables ...Visitable) *GroupingNode {
	return predicationLtAll(node, visitables...)
}

func (node InfixOperationNode) LtEq(v Visitable) *LessThanOrEqualNode {
	return &LessThanOrEqualNode{Left: node, Right: v}
}

func (node InfixOperationNode) LtEqAny(visitables ...Visitable) *GroupingNode {
	return predicationLtEqAny(node, visitables...)
}

func (node InfixOperationNode) LtEqAll(visitables ...Visitable) *GroupingNode {
	return predicationLtEqAll(node, visitables...)
}

func (node InfixOperationNode) Gt(visitable Visitable) *GreaterThanNode {
	return predicationGt(node, visitable)
}

func (node InfixOperationNode) GtAny(visitables ...Visitable) *GroupingNode {
	return predicationGtAny(node, visitables...)
}

func (node InfixOperationNode) GtAll(visitables ...Visitable) *GroupingNode {
	return predicationGtAll(node, visitables...)
}

func (node InfixOperationNode) GtEq(visitable Visitable) *GreaterThanOrEqualNode {
	return predicationGtEq(node, visitable)
}

func (node InfixOperationNode) GtEqAny(visitables ...Visitable) *GroupingNode {
	return predicationGtEqAny(node, visitables...)
}

func (node InfixOperationNode) GtEqAll(visitables ...Visitable) *GroupingNode {
	return predicationGtEqAll(node, visitables...)
}

func (node InfixOperationNode) Count() *CountNode {
	return predicationCount(node)
}

func (node InfixOperationNode) Extract(literal SqlLiteralNode) *ExtractNode {
	return predicationExtract(node, literal)
}

func (node InfixOperationNode) As(visitable Visitable) *AsNode {
	return predicationAs(node, visitable)
}

func (node InfixOperationNode) In(visitables []Visitable) Visitable {
	return predicationIn(node, visitables)
}

func (node InfixOperationNode) InAny(visitableslices ...[]Visitable) Visitable {
	return predicationInAny(node, visitableslices...)
}

func (node InfixOperationNode) InAll(visitableslices ...[]Visitable) Visitable {
	return predicationInAll(node, visitableslices...)
}

func (node InfixOperationNode) NotIn(visitables []Visitable) Visitable {
	return predicationNotIn(node, visitables)
}

func (node InfixOperationNode) NotInAny(visitableslices ...[]Visitable) Visitable {
	return predicationNotInAny(node, visitableslices...)
}

func (node InfixOperationNode) NotInAll(visitableslices ...[]Visitable) Visitable {
	return predicationNotInAll(node, visitableslices...)
}

func (node InfixOperationNode) NotEq(visitable Visitable) *NotEqualNode {
	return predicationNotEq(node, visitable)
}

func (node InfixOperationNode) NotEqAny(visitables ...Visitable) *GroupingNode {
	return predicationNotEqAny(node, visitables...)
}

func (node InfixOperationNode) NotEqAll(visitables ...Visitable) *GroupingNode {
	return predicationNotEqAll(node, visitables...)
}

func (node InfixOperationNode) DoesNotMatch(literal SqlLiteralNode) *DoesNotMatchNode {
	return predicationDoesNotMatch(node, literal)
}

func (node InfixOperationNode) DoesNotMatchAny(literals ...SqlLiteralNode) *GroupingNode {
	return predicationDoesNotMatchAny(node, literals...)
}

func (node InfixOperationNode) DoesNotMatchAll(literals ...SqlLiteralNode) *GroupingNode {
	return predicationDoesNotMatchAll(node, literals...)
}

func (node InfixOperationNode) Matches(literal SqlLiteralNode) *MatchesNode {
	return predicationMatches(node, literal)
}

func (node InfixOperationNode) MatchesAny(literals ...SqlLiteralNode) *GroupingNode {
	return predicationMatchesAny(node, literals...)
}

func (node InfixOperationNode) MatchesAll(literals ...SqlLiteralNode) *GroupingNode {
	return predicationMatchesAll(node, literals...)
}

func (node InfixOperationNode) GroupAny(visitables ...Visitable) *GroupingNode {
	return predicationGroupAny(node, visitables...)
}

func (node InfixOperationNode) GroupAll(visitables ...Visitable) *GroupingNode {
	return predicationGroupAll(node, visitables...)
}
