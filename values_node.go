package rel

type ValuesNode struct {
	Values  []interface{}
	Columns []*AttributeNode
	BaseVisitable
}
