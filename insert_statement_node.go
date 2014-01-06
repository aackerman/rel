package rel

type InsertStatementNode struct {
	Relation *Table
	Columns  *[]AttributeNode
	Values   *ValuesNode
	BaseVisitable
}
