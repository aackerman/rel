package rel

type FunctionNode struct {
	Expressions AstNode
	Alias       *SqlLiteralNode
	BaseNode
}

type SumNode FunctionNode
type MaxNode FunctionNode
type MinNode FunctionNode
type AvgNode FunctionNode
