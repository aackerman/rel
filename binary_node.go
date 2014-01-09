package rel

type BinaryNode struct {
	Left  Visitable
	Right Visitable
	BaseVisitable
}

func NewBinaryNode(left Visitable, right Visitable) BinaryNode {
	return BinaryNode{
		Left:  left,
		Right: right,
	}
}

type BetweenNode BinaryNode
type AssignmentNode BinaryNode
type DoesNotMatchNode BinaryNode
type GreaterThanNode BinaryNode
type GreaterThanOrEqualNode BinaryNode
type JoinNode BinaryNode
type LessThanNode BinaryNode
type LessThanOrEqualNode BinaryNode
type MatchesNode BinaryNode
type NotEqualNode BinaryNode
type OrNode BinaryNode
type UnionNode BinaryNode
type UnionAllNode BinaryNode
type IntersectNode BinaryNode
type ExceptNode BinaryNode
