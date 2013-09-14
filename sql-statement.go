package arel

type SqlStatement struct{}

func (s *SqlStatement) CreateTrue() *TrueNode {
	return NewTrueNode()
}

func (s *SqlStatement) CreateFalse() *FalseNode {
	return NewFalseNode()
}

func (s *SqlStatement) CreateTableAlias(relation *Table, name string) *TableAliasNode {
	return NewTableAliasNode(relation, name)
}

func (s *SqlStatement) CreateJoin(to string) *JoinNode {
	return &JoinNode{}
}

func (s *SqlStatement) CreateAnd(clauses []interface{}) *AndNode {
	return NewAndNode()
}

func (s *SqlStatement) CreateOn(expr string) *OnNode {
	return NewOnNode()
}

func (s *SqlStatement) Grouping(expr string) *GroupingNode {
	return NewGroupingNode()
}

func (s *SqlStatement) Lower() {

}
