package rel

type FunctionNode struct {
	Expressions Visitable
	Alias       *SqlLiteralNode
	BaseVisitable
}

type SumNode FunctionNode
type MaxNode FunctionNode
type MinNode FunctionNode
type AvgNode FunctionNode
