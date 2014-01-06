package rel

type InsertStatementNode struct {
	Relation *Table
	Columns  *[]SqlLiteralNode
	Values   *ValuesNode
	BaseVisitable
}
