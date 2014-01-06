package rel

type InsertStatementNode struct {
	Relation *Table
	columns  *[]Visitable
	values   *[]interface{}
	BaseVisitable
}
