package rel

type UpdateStatementNode struct {
	table  *Table
	Wheres *[]AstNode
	Values *[]AstNode
	Orders *[]AstNode
	Limit  int
	Key    string
	BaseNode
}
