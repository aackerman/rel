package arel

type SelectStatement struct {
	Cores  []SelectCore
	Limit  int
	Orders []Order
	Lock   bool
}

func (s SelectStatement) NodeInterface()         {}
func (s SelectStatement) SqlStatementInterface() {}
