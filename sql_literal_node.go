package rel

import (
	"fmt"
)

type SqlLiteralNode struct {
	Raw string
	BaseVisitable
}

func Sql(raw interface{}) SqlLiteralNode {
	return SqlLiteralNode{Raw: fmt.Sprintf("%s", raw)}
}

func Star() SqlLiteralNode {
	return Sql("*")
}

func (node SqlLiteralNode) As(literal SqlLiteralNode) *AsNode {
	return aliasPredicationAs(node, literal)
}

func (node SqlLiteralNode) Desc() *DescendingNode {
	return orderingDesc(node)
}

func (node SqlLiteralNode) Asc() *AscendingNode {
	return orderingAsc(node)
}

func (node SqlLiteralNode) Eq(visitable Visitable) *EqualityNode {
	return predicationEq(node, visitable)
}

func (node SqlLiteralNode) EqAny(visitables ...Visitable) *GroupingNode {
	return predicationEqAny(node, visitables...)
}

func (node SqlLiteralNode) EqAll(visitables ...Visitable) *GroupingNode {
	return predicationEqAll(node, visitables...)
}

func (node SqlLiteralNode) Lt(visitable Visitable) *LessThanNode {
	return predicationLt(node, visitable)
}

func (node SqlLiteralNode) LtAny(visitables ...Visitable) *GroupingNode {
	return predicationLtAny(node, visitables...)
}

func (node SqlLiteralNode) LtAll(visitables ...Visitable) *GroupingNode {
	return predicationLtAll(node, visitables...)
}

func (node SqlLiteralNode) LtEq(v Visitable) *LessThanOrEqualNode {
	return &LessThanOrEqualNode{Left: node, Right: v}
}

func (node SqlLiteralNode) LtEqAny(visitables ...Visitable) *GroupingNode {
	return predicationLtEqAny(node, visitables...)
}

func (node SqlLiteralNode) LtEqAll(visitables ...Visitable) *GroupingNode {
	return predicationLtEqAll(node, visitables...)
}

func (node SqlLiteralNode) Gt(visitable Visitable) *GreaterThanNode {
	return predicationGt(node, visitable)
}

func (node SqlLiteralNode) GtAny(visitables ...Visitable) *GroupingNode {
	return predicationGtAny(node, visitables...)
}

func (node SqlLiteralNode) GtAll(visitables ...Visitable) *GroupingNode {
	return predicationGtAll(node, visitables...)
}

func (node SqlLiteralNode) GtEq(visitable Visitable) *GreaterThanOrEqualNode {
	return predicationGtEq(node, visitable)
}

func (node SqlLiteralNode) GtEqAny(visitables ...Visitable) *GroupingNode {
	return predicationGtEqAny(node, visitables...)
}

func (node SqlLiteralNode) GtEqAll(visitables ...Visitable) *GroupingNode {
	return predicationGtEqAll(node, visitables...)
}

func (node SqlLiteralNode) Count() *CountNode {
	return predicationCount(node)
}

func (node SqlLiteralNode) Extract(literal SqlLiteralNode) *ExtractNode {
	return predicationExtract(node, literal)
}

func (node SqlLiteralNode) In(visitables []Visitable) Visitable {
	return predicationIn(node, visitables)
}

func (node SqlLiteralNode) InAny(visitableslices ...[]Visitable) Visitable {
	return predicationInAny(node, visitableslices...)
}

func (node SqlLiteralNode) InAll(visitableslices ...[]Visitable) Visitable {
	return predicationInAll(node, visitableslices...)
}

func (node SqlLiteralNode) NotIn(visitables []Visitable) Visitable {
	return predicationNotIn(node, visitables)
}

func (node SqlLiteralNode) NotInAny(visitableslices ...[]Visitable) Visitable {
	return predicationNotInAny(node, visitableslices...)
}

func (node SqlLiteralNode) NotInAll(visitableslices ...[]Visitable) Visitable {
	return predicationNotInAll(node, visitableslices...)
}

func (node SqlLiteralNode) NotEq(visitable Visitable) *NotEqualNode {
	return predicationNotEq(node, visitable)
}

func (node SqlLiteralNode) NotEqAny(visitables ...Visitable) *GroupingNode {
	return predicationNotEqAny(node, visitables...)
}

func (node SqlLiteralNode) NotEqAll(visitables ...Visitable) *GroupingNode {
	return predicationNotEqAll(node, visitables...)
}

func (node SqlLiteralNode) DoesNotMatch(literal SqlLiteralNode) *DoesNotMatchNode {
	return predicationDoesNotMatch(node, literal)
}

func (node SqlLiteralNode) DoesNotMatchAny(literals ...SqlLiteralNode) *GroupingNode {
	return predicationDoesNotMatchAny(node, literals...)
}

func (node SqlLiteralNode) DoesNotMatchAll(literals ...SqlLiteralNode) *GroupingNode {
	return predicationDoesNotMatchAll(node, literals...)
}

func (node SqlLiteralNode) Matches(literal SqlLiteralNode) *MatchesNode {
	return predicationMatches(node, literal)
}

func (node SqlLiteralNode) MatchesAny(literals ...SqlLiteralNode) *GroupingNode {
	return predicationMatchesAny(node, literals...)
}

func (node SqlLiteralNode) MatchesAll(literals ...SqlLiteralNode) *GroupingNode {
	return predicationMatchesAll(node, literals...)
}

func (node SqlLiteralNode) GroupAny(visitables ...Visitable) *GroupingNode {
	return predicationGroupAny(node, visitables...)
}

func (node SqlLiteralNode) GroupAll(visitables ...Visitable) *GroupingNode {
	return predicationGroupAll(node, visitables...)
}
)
