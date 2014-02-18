package rel

type UnaryNode struct {
	Expr Visitable
	BaseVisitable
}

type GroupNode UnaryNode
type HavingNode UnaryNode
type LimitNode UnaryNode
type NotNode UnaryNode
type OffsetNode UnaryNode
type OnNode UnaryNode
type UsingNode UnaryNode
type OrderingNode UnaryNode
type TopNode UnaryNode
type LockNode UnaryNode
type DistinctOnNode UnaryNode
type WithNode UnaryNode
type WithRecursiveNode UnaryNode
type RowsNode UnaryNode
type RangeNode UnaryNode
type CurrentRowNode UnaryNode
type PrecedingNode UnaryNode
type FollowingNode UnaryNode

func NewUnaryNode(visitable Visitable) *UnaryNode {
	return &UnaryNode{Expr: visitable}
}

func NewBinNode(visitable Visitable) *BinNode {
	return &BinNode{Expr: visitable}
}

func NewGroupNode(visitable Visitable) *GroupNode {
	return &GroupNode{Expr: visitable}
}

func NewHavingNode(visitable Visitable) *HavingNode {
	return &HavingNode{Expr: visitable}
}

func NewLimitNode(visitable Visitable) *LimitNode {
	return &LimitNode{Expr: visitable}
}

func NewNotNode(visitable Visitable) *NotNode {
	return &NotNode{Expr: visitable}
}

func NewOffsetNode(visitable Visitable) *OffsetNode {
	return &OffsetNode{Expr: visitable}
}

func NewOrderingNode(visitable Visitable) *OrderingNode {
	return &OrderingNode{Expr: visitable}
}

func NewTopNode(visitable Visitable) *TopNode {
	return &TopNode{Expr: visitable}
}

func NewLockNode(visitable Visitable) *LockNode {
	return &LockNode{Expr: visitable}
}

func NewDistinctOnNode(visitable Visitable) *DistinctOnNode {
	return &DistinctOnNode{Expr: visitable}
}

func NewWithNode(visitable Visitable) *WithNode {
	return &WithNode{Expr: visitable}
}

func NewWithRecursiveNode(visitable Visitable) *WithRecursiveNode {
	return &WithRecursiveNode{Expr: visitable}
}
