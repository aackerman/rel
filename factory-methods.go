package arel

type FactoryMethods struct{}

func (f *FactoryMethods) CreateTrue() *TrueNode {
	return NewTrueNode()
}

func (f *FactoryMethods) CreateFalse() *FalseNode {
	return NewFalseNode()
}

func (f *FactoryMethods) CreateTableAlias(relation *Table, name string) *TableAliasNode {
	return NewTableAliasNode(relation, name)
}

func (f *FactoryMethods) CreateJoin(to string) *JoinNode {
	return &JoinNode{}
}

func (f *FactoryMethods) CreateAnd(clauses []interface{}) *AndNode {
	return NewAndNode()
}

func (f *FactoryMethods) CreateOn(expr string) *OnNode {
	return NewOnNode()
}

func (f *FactoryMethods) CreateGrouping(expr string) *GroupingNode {
	return NewGroupingNode()
}
