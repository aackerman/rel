package rel

type IntersectManager struct {
	Engine Engine
	Ast    Visitable
}

func NewIntersectManager(e Engine) *IntersectManager {
	return &IntersectManager{Engine: e}
}

func (i *IntersectManager) ToSql() string {
	return i.Engine.Visitor().Accept(i.Ast)
}

func (i *IntersectManager) Intersect(mgr1 SelectManager, mgr2 SelectManager) *IntersectManager {
	i.Ast = IntersectNode{
		Left:  mgr1.Ast,
		Right: mgr2.Ast,
	}
	return i
}
