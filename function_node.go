package arel

type FunctionNode struct {
	BaseNode
}

type SumNode FunctionNode
type ExistsNode FunctionNode
type MaxNode FunctionNode
type MinNode FunctionNode
type AvgNode FunctionNode
