package rel

type UnaryNode struct {
	Expr Visitable
	BaseVisitable
}

type BinNode UnaryNode
type GroupNode UnaryNode
type HavingNode UnaryNode
type LimitNode UnaryNode
type NotNode UnaryNode
type OffsetNode UnaryNode
type OnNode UnaryNode
type OrderingNode UnaryNode
type TopNode UnaryNode
type LockNode UnaryNode
type DistinctOnNode UnaryNode
type WithNode UnaryNode
type WithRecursiveNode UnaryNode

func NewUnaryNode(n Visitable) UnaryNode {
	return UnaryNode{Expr: n}
}

func NewBinNode(n Visitable) BinNode {
	return BinNode{Expr: n}
}

func NewGroupNode(n Visitable) GroupNode {
	return GroupNode{Expr: n}
}

func NewHavingNode(n Visitable) HavingNode {
	return HavingNode{Expr: n}
}

func NewLimitNode(n Visitable) LimitNode {
	return LimitNode{Expr: n}
}

func NewNotNode(n Visitable) NotNode {
	return NotNode{Expr: n}
}

func NewOffsetNode(n Visitable) OffsetNode {
	return OffsetNode{Expr: n}
}

func NewOnNode(n Visitable) OnNode {
	return OnNode{Expr: n}
}

func NewOrderingNode(n Visitable) OrderingNode {
	return OrderingNode{Expr: n}
}

func NewTopNode(n Visitable) TopNode {
	return TopNode{Expr: n}
}

func NewLockNode(n Visitable) LockNode {
	return LockNode{Expr: n}
}

func NewDistinctOnNode(n Visitable) DistinctOnNode {
	return DistinctOnNode{Expr: n}
}

func NewWithNode(n Visitable) WithNode {
	return WithNode{Expr: n}
}

func NewWithRecursiveNode(n Visitable) WithRecursiveNode {
	return WithRecursiveNode{Expr: n}
}
