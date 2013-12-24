package arel

type UnaryNode struct {
	expr SqlLiteralNode
	BaseNode
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

func NewUnaryNode(n SqlLiteralNode) UnaryNode {
	return UnaryNode{expr: n}
}

func NewBinNode(n SqlLiteralNode) BinNode {
	return BinNode{expr: n}
}

func NewGroupNode(n SqlLiteralNode) GroupNode {
	return GroupNode{expr: n}
}

func NewHavingNode(n SqlLiteralNode) HavingNode {
	return HavingNode{expr: n}
}

func NewLimitNode(n SqlLiteralNode) LimitNode {
	return LimitNode{expr: n}
}

func NewNotNode(n SqlLiteralNode) NotNode {
	return NotNode{expr: n}
}

func NewOffsetNode(n SqlLiteralNode) OffsetNode {
	return OffsetNode{expr: n}
}

func NewOnNode(n SqlLiteralNode) OnNode {
	return OnNode{expr: n}
}

func NewOrderingNode(n SqlLiteralNode) OrderingNode {
	return OrderingNode{expr: n}
}

func NewTopNode(n SqlLiteralNode) TopNode {
	return TopNode{expr: n}
}

func NewLockNode(n SqlLiteralNode) LockNode {
	return LockNode{expr: n}
}

func NewDistinctOnNode(n SqlLiteralNode) DistinctOnNode {
	return DistinctOnNode{expr: n}
}

func NewWithNode(n SqlLiteralNode) WithNode {
	return WithNode{expr: n}
}

func NewWithRecursiveNode(n SqlLiteralNode) WithRecursiveNode {
	return WithRecursiveNode{expr: n}
}
