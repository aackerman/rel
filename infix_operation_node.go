package rel

type InfixOperationNode struct {
	Operator SqlLiteralNode
	Left     Visitable
	Right    Visitable
}
