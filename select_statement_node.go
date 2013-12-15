package arel

type SelectStatement struct {
	cores  []SelectCore
	Limit  int
	Orders []Order
	Lock   bool
	SqlStatement
	AstNode
}

func (s *SelectStatement) Cores() []*SelectCore {
	return s.cores
}

func (s *SelectStatement) IsEqual(s2 SelectStatement) bool {
	return false
}
