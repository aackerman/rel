package rel

type DeleteStatementNode struct {
	Relation *Table
	Wheres   *[]Visitable
	BaseVisitable
}
