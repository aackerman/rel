package rel

type ValuesNode struct {
	Values  []Visitable
	Columns []SqlLiteralNode
	BaseVisitable
}
