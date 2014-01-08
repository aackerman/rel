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

func (n CountNode) Desc() DescendingNode {
	return DescendingNode{Expr: n}
}

func (n CountNode) As(v Visitable) AsNode {
	return AsNode{
		Left:  n,
		Right: v,
	}
}
