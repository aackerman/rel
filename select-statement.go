package arel

type SelectStatementNode struct {
	Cores  []ArelNode
	Limit  int
	Orders []ArelNode
	Lock   bool
	ArelNode
}

func NewSelectStatementNode() *SelectStatementNode {
	return &SelectStatementNode{
		make([]ArelNode, 10),
		0,
		make([]ArelNode, 10),
		false,
		FactoryMethods{},
	}
}
