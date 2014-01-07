package rel

type SelectStatementNode struct {
	Cores  []*SelectCoreNode
	Limit  *LimitNode
	Orders *[]Visitable
	Lock   *LockNode
	With   Visitable // WithNode or WithRecursiveNode
	Offset *OffsetNode
	Visitable
}

func NewSelectStatementNode() SelectStatementNode {
	stmt := SelectStatementNode{Cores: make([]*SelectCoreNode, 0)}
	core := NewSelectCoreNode()
	stmt.Cores = append(stmt.Cores, &core)
	return stmt
}

func (s *SelectStatementNode) IsEqual(s2 SelectStatementNode) bool {
	return false
}
