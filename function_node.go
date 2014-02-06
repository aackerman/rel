package rel

type FunctionNode struct {
	Expressions []Visitable
	Alias       *SqlLiteralNode
	Distinct    bool
	BaseVisitable
}

type SumNode FunctionNode
type MaxNode FunctionNode
type MinNode FunctionNode
type AvgNode FunctionNode

func Sum(attr *AttributeNode) *SumNode {
	return &SumNode{Expressions: []Visitable{attr}}
}

func Max(attr *AttributeNode) *MaxNode {
	return &MaxNode{Expressions: []Visitable{attr}}
}

func Min(attr *AttributeNode) *MinNode {
	return &MinNode{Expressions: []Visitable{attr}}
}

func Avg(attr *AttributeNode) *AvgNode {
	return &AvgNode{Expressions: []Visitable{attr}}
}

func (node *FunctionNode) Over(visitable Visitable) *OverNode {
	return windowPredicationOver(node, visitable)
}

func (node *FunctionNode) Eq(visitable Visitable) *EqualityNode {
	return predicationEq(node, visitable)
}

func (node *FunctionNode) EqAny(visitables ...Visitable) *GroupingNode {
	return predicationEqAny(node, visitables...)
}

func (node *FunctionNode) EqAll(visitables ...Visitable) *GroupingNode {
	return predicationEqAll(node, visitables...)
}

func (node *FunctionNode) Lt(visitable Visitable) *LessThanNode {
	return predicationLt(node, visitable)
}

func (node *FunctionNode) LtAny(visitables ...Visitable) *GroupingNode {
	return predicationLtAny(node, visitables...)
}

func (node *FunctionNode) LtAll(visitables ...Visitable) *GroupingNode {
	return predicationLtAll(node, visitables...)
}

func (node *FunctionNode) LtEq(visitable Visitable) *LessThanOrEqualNode {
	return predicationLtEq(node, visitable)
}

func (node *FunctionNode) LtEqAny(visitables ...Visitable) *GroupingNode {
	return predicationLtEqAny(node, visitables...)
}

func (node *FunctionNode) LtEqAll(visitables ...Visitable) *GroupingNode {
	return predicationLtEqAll(node, visitables...)
}

func (node *FunctionNode) Gt(visitable Visitable) *GreaterThanNode {
	return predicationGt(node, visitable)
}

func (node *FunctionNode) GtAny(visitables ...Visitable) *GroupingNode {
	return predicationGtAny(node, visitables...)
}

func (node *FunctionNode) GtAll(visitables ...Visitable) *GroupingNode {
	return predicationGtAll(node, visitables...)
}

func (node *FunctionNode) GtEq(visitable Visitable) *GreaterThanOrEqualNode {
	return predicationGtEq(node, visitable)
}

func (node *FunctionNode) GtEqAny(visitables ...Visitable) *GroupingNode {
	return predicationGtEqAny(node, visitables...)
}

func (node *FunctionNode) GtEqAll(visitables ...Visitable) *GroupingNode {
	return predicationGtEqAll(node, visitables...)
}

func (node *FunctionNode) Count() *CountNode {
	return predicationCount(node)
}

func (node *FunctionNode) Extract(literal SqlLiteralNode) *ExtractNode {
	return predicationExtract(node, literal)
}

func (node *FunctionNode) In(visitables []Visitable) Visitable {
	return predicationIn(node, visitables)
}

func (node *FunctionNode) InAny(visitableslices ...[]Visitable) Visitable {
	return predicationInAny(node, visitableslices...)
}

func (node *FunctionNode) InAll(visitableslices ...[]Visitable) Visitable {
	return predicationInAll(node, visitableslices...)
}

func (node *FunctionNode) NotIn(visitables []Visitable) Visitable {
	return predicationNotIn(node, visitables)
}

func (node *FunctionNode) NotInAny(visitableslices ...[]Visitable) Visitable {
	return predicationNotInAny(node, visitableslices...)
}

func (node *FunctionNode) NotInAll(visitableslices ...[]Visitable) Visitable {
	return predicationNotInAll(node, visitableslices...)
}

func (node *FunctionNode) NotEq(visitable Visitable) *NotEqualNode {
	return predicationNotEq(node, visitable)
}

func (node *FunctionNode) NotEqAny(visitables ...Visitable) *GroupingNode {
	return predicationNotEqAny(node, visitables...)
}

func (node *FunctionNode) NotEqAll(visitables ...Visitable) *GroupingNode {
	return predicationNotEqAll(node, visitables...)
}

func (node *FunctionNode) DoesNotMatch(literal SqlLiteralNode) *DoesNotMatchNode {
	return predicationDoesNotMatch(node, literal)
}

func (node *FunctionNode) DoesNotMatchAny(literals ...SqlLiteralNode) *GroupingNode {
	return predicationDoesNotMatchAny(node, literals...)
}

func (node *FunctionNode) DoesNotMatchAll(literals ...SqlLiteralNode) *GroupingNode {
	return predicationDoesNotMatchAll(node, literals...)
}

func (node *FunctionNode) Matches(literal SqlLiteralNode) *MatchesNode {
	return predicationMatches(node, literal)
}

func (node *FunctionNode) MatchesAny(literals ...SqlLiteralNode) *GroupingNode {
	return predicationMatchesAny(node, literals...)
}

func (node *FunctionNode) MatchesAll(literals ...SqlLiteralNode) *GroupingNode {
	return predicationMatchesAll(node, literals...)
}
