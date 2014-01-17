package rel

type Orderer interface {
	Desc() *DescendingNode
	Asc() *AscendingNode
	Visitable
}

func orderingDesc(node Orderer) *DescendingNode {
	return &DescendingNode{Expr: node}
}

func orderingAsc(node Orderer) *AscendingNode {
	return &AscendingNode{Expr: node}
}
