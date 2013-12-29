package arel

type UnaryNode struct {
	Expr AstNode
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
	return UnaryNode{Expr: n}
}

func NewBinNode(n AstNode) BinNode {
	return BinNode{Expr: n}
}

func NewGroupNode(n AstNode) GroupNode {
	return GroupNode{Expr: n}
}

func NewHavingNode(n AstNode) HavingNode {
	return HavingNode{Expr: n}
}

func NewLimitNode(n AstNode) LimitNode {
	return LimitNode{Expr: n}
}

func NewNotNode(n AstNode) NotNode {
	return NotNode{Expr: n}
}

func NewOffsetNode(n AstNode) OffsetNode {
	return OffsetNode{Expr: n}
}

func NewOnNode(n AstNode) OnNode {
	return OnNode{Expr: n}
}

func NewOrderingNode(n AstNode) OrderingNode {
	return OrderingNode{Expr: n}
}

func NewTopNode(n AstNode) TopNode {
	return TopNode{Expr: n}
}

func NewLockNode(n AstNode) LockNode {
	return LockNode{Expr: n}
}

func NewDistinctOnNode(n AstNode) DistinctOnNode {
	return DistinctOnNode{Expr: n}
}

func NewWithNode(n AstNode) WithNode {
	return WithNode{Expr: n}
}

func NewWithRecursiveNode(n AstNode) WithRecursiveNode {
	return WithRecursiveNode{Expr: n}
}
