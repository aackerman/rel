package arel

type SelectStatementNode struct {
	Cores  []SelectCoreNode
	Limit  int
	Orders []OrderNode
	Lock   bool
	ArelNode
	SqlStatement
}

func NewSelectStatementNode() *SelectStatementNode {
	return &SelectStatementNode{
		Cores:        make([]SelectCoreNode, 10),
		Limit:        0,
		Orders:       make([]OrderNode, 10),
		ArelNode:     ArelNode{},
		SqlStatement: SqlStatement{},
	}
}
