package arel

type NodeCreator interface {
	CreateTrue() *TrueNode
	CreateFalse() *FalseNode
	CreateTableAlias(*Table, string) *TableAliasNode
	CreateJoin(string) *JoinNode
	CreateAnd([]interface{}) *AndNode
	CreateOn(string) *OnNode
	CreateGrouping(string) *GroupingNode
}

type BaseNodeCreator struct{}

func (n *BaseNodeCreator) CreateTrue() *TrueNode {
	return NewTrueNode()
}

func (n *BaseNodeCreator) CreateFalse() *FalseNode {
	return NewFalseNode()
}

func (n *BaseNodeCreator) CreateTableAlias(relation *Table, name string) *TableAliasNode {
	return NewTableAliasNode(relation, name)
}

func (n *BaseNodeCreator) CreateJoin(to string) *JoinNode {
	return &JoinNode{}
}

func (n *BaseNodeCreator) CreateAnd(clauses []interface{}) *AndNode {
	return NewAndNode()
}

func (n *BaseNodeCreator) CreateOn(expr string) *OnNode {
	return NewOnNode()
}

func (n *BaseNodeCreator) CreateGrouping(expr string) *GroupingNode {
	return NewGroupingNode()
}
