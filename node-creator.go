package arel

type NodeCreator struct{}

func (n *NodeCreator) CreateTrue() *TrueNode {
	return NewTrueNode()
}

func (n *NodeCreator) CreateFalse() *FalseNode {
	return NewFalseNode()
}

func (n *NodeCreator) CreateTableAlias(relation *Table, name string) *TableAliasNode {
	return NewTableAliasNode(relation, name)
}

func (n *NodeCreator) CreateJoin(to string) *JoinNode {
	return &JoinNode{}
}

func (n *NodeCreator) CreateAnd(clauses []interface{}) *AndNode {
	return NewAndNode()
}

func (n *NodeCreator) CreateOn(expr string) *OnNode {
	return NewOnNode()
}

func (n *NodeCreator) CreateGrouping(expr string) *GroupingNode {
	return NewGroupingNode()
}
