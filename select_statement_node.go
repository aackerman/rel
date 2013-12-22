package arel

type SelectStatement struct {
	Cores  []*SelectCoreNode
	Limit  *LimitNode
	Orders *[]OrderNode
	Lock   *LockNode
	With   *WithNode
	Offset *OffsetNode
	AstNode
}

func (s *SelectStatement) IsEqual(s2 SelectStatement) bool {
	return false
}
