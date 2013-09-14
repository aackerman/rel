package arel

type TreeManager struct {
	Ast    *SqlStatement
	Engine *Engine
	ctx    Context
}

func NewTreeManager(e *Engine) *TreeManager {
	return &TreeManager{
		Engine: e,
	}
}

func (t *TreeManager) ToSql() string {
	return t.Visitor().Accept(t.Ast)
}

func (t *TreeManager) Visitor() *Visitor {
	return t.Engine.Connection().Visitor()
}
