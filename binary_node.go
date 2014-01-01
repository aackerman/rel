package rel

type BinaryNode struct {
	Left  *Table
	Right *Table
	BaseVisitable
}

func NewBinaryNode(left *Table, right *Table) BinaryNode {
	return BinaryNode{
		Left:  left,
		Right: right,
	}
}

type AssignmentNode BinaryNode
type DoesNotMatchNode BinaryNode
type GreaterThanNode BinaryNode
type GreaterThanOrEqualNode BinaryNode
type JoinNode BinaryNode
type LessThanNode BinaryNode
type LessThanOrEqualNode BinaryNode
type MatchesNode BinaryNode
type NotEqualNode BinaryNode
type NotInNode BinaryNode
type OrNode BinaryNode
type UnionNode BinaryNode
type UnionAllNode BinaryNode
type IntersectNode BinaryNode
type ExceptNode BinaryNode
type InnerJoinNode JoinNode
type OuterJoinNode JoinNode
