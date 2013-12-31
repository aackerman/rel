package grel

type SelectStatementNode struct {
	Cores  []*SelectCoreNode
	Limit  *LimitNode
	Orders *[]AstNode
	Lock   *LockNode
	With   *WithNode
	Offset *OffsetNode
	AstNode
}

func (s *SelectStatementNode) IsEqual(s2 SelectStatementNode) bool {
	return false
}
