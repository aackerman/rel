package rel

// A UnionManager manages the handling of UNIONs on SELECT statements
type UnionManager struct {
	Engine Engine
	Ast    Visitable
}

func (u *UnionManager) ToSql() string {
	return u.Engine.Visitor().Accept(u.Ast)
}

// Union appends a SelectManager to allow for more unions
func (u *UnionManager) Union(mgr1 SelectManager, mgr2 SelectManager) *UnionManager {
	u.Ast = UnionNode{
		Left:  mgr1.Ast,
		Right: mgr2.Ast,
	}
	return u
}

func NewUnionManager(e Engine) *UnionManager {
	return &UnionManager{Engine: e}
}
