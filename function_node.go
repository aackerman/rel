package rel

type FunctionNode struct {
	Expressions []Visitable
	Alias       *SqlLiteralNode
	Distinct    bool
	BaseVisitable
}

type SumNode FunctionNode
type MaxNode FunctionNode
type MinNode FunctionNode
type AvgNode FunctionNode
type CountNode FunctionNode

func (node *CountNode) Desc() *DescendingNode {
	return &DescendingNode{Expr: node}
}

func (node *CountNode) As(v Visitable) *AsNode {
	return &AsNode{
		Left:  node,
		Right: v,
	}
}
