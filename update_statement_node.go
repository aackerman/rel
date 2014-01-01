package rel

type UpdateStatementNode struct {
	table  *Table
	Wheres *[]Visitable
	Values *[]Visitable
	Orders *[]Visitable
	Limit  int
	Key    string
	BaseVisitable
}
