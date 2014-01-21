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

func NewSelectStatementNode() *SelectStatementNode {
	return &SelectStatementNode{
		Cores: []*SelectCoreNode{NewSelectCoreNode()},
	}
}

func (s *SelectStatementNode) IsEqual(s2 SelectStatementNode) bool {
	return false
}
