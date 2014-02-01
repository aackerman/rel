package rel

type UpdateStatementNode struct {
	Relation *Table
	Wheres   *[]Visitable
	Values   *[]Visitable
	Orders   *[]Visitable
	Limit    *LimitNode
	Key      Visitable // SqlLiteralNode AttributeNode
	BaseVisitable
}

func NewUpdateStatementNode() *UpdateStatementNode {
	return &UpdateStatementNode{
		Wheres: &[]Visitable{},
		Values: &[]Visitable{},
		Orders: &[]Visitable{},
	}
}
