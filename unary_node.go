package arel

type UnaryNode struct {
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
