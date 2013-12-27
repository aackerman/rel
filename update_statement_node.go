package arel

type UpdateStatementNode struct {
	table  *Table
	Wheres *[]AstNode
	Values *[]AstNode
	Orders *[]OrderingNode
	Limit  int
	Key    string
	BaseNode
}
