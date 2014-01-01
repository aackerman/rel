package rel

type FunctionNode struct {
	Expressions AstNode
	Alias       *SqlLiteralNode
	BaseNode
}

type SumNode FunctionNode
type ExistsNode FunctionNode
type MaxNode FunctionNode
type MinNode FunctionNode
type AvgNode FunctionNode

func NewExistsNode(n AstNode) ExistsNode {
	return ExistsNode{Expressions: n}
}

func (e ExistsNode) As(n AstNode) AsNode {
	return AsNode{
		Left:  e,
		Right: &n,
	}
}
