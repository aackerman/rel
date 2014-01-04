package rel

type MultiStatementManager struct {
	Engine Engine
	Ast    Visitable
}

func (mgr *MultiStatementManager) ToSql() string {
	return mgr.Engine.Visitor().Accept(mgr.Ast)
}

func NewMultiStatementManager(e Engine) *MultiStatementManager {
	return &MultiStatementManager{Engine: e}
}

func (mgr *MultiStatementManager) Intersect(stmt1 Visitable, stmt2 Visitable) *MultiStatementManager {
	mgr.Ast = IntersectNode{
		Left:  stmt1,
		Right: stmt2,
	}
	return mgr
}

func (mgr *MultiStatementManager) Union(stmt1 Visitable, stmt2 Visitable) *MultiStatementManager {
	mgr.Ast = UnionNode{
		Left:  stmt1,
		Right: stmt2,
	}
	return mgr
}

func (mgr *MultiStatementManager) UnionAll(stmt1 Visitable, stmt2 Visitable) *MultiStatementManager {
	mgr.Ast = UnionAllNode{
		Left:  stmt1,
		Right: stmt2,
	}
	return mgr
}
