package arel

type UnaryNode struct {
	expr AstNode
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

func NewUnaryNode(n AstNode) UnaryNode {
	return UnaryNode{expr: n}
}

func NewBinNode(n AstNode) BinNode {
	return BinNode{expr: n}
}

func NewGroupNode(n AstNode) GroupNode {
	return GroupNode{expr: n}
}

func NewHavingNode(n AstNode) HavingNode {
	return HavingNode{expr: n}
}

func NewLimitNode(n AstNode) LimitNode {
	return LimitNode{expr: n}
}

func NewNotNode(n AstNode) NotNode {
	return NotNode{expr: n}
}

func NewOffsetNode(n AstNode) OffsetNode {
	return OffsetNode{expr: n}
}

func NewOnNode(n AstNode) OnNode {
	return OnNode{expr: n}
}

func NewOrderingNode(n AstNode) OrderingNode {
	return OrderingNode{expr: n}
}

func NewTopNode(n AstNode) TopNode {
	return TopNode{expr: n}
}

func NewLockNode(n AstNode) LockNode {
	return LockNode{expr: n}
}

func NewDistinctOnNode(n AstNode) DistinctOnNode {
	return DistinctOnNode{expr: n}
}

func NewWithNode(n AstNode) WithNode {
	return WithNode{expr: n}
}

func NewWithRecursiveNode(n AstNode) WithRecursiveNode {
	return WithRecursiveNode{expr: n}
}
