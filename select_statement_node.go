package arel

type SelectStatement struct {
	Cores  []SelectCoreNode
	Limit  int
	Orders []OrderNode
	Lock   bool
	AstNode
}

func (s *SelectStatement) IsEqual(s2 SelectStatement) bool {
	return false
}
