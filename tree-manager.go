package arel

type TreeManager struct {
	Engine *Engine
	ctx    ArelNode
	Ast    *SqlStatement
}

func (t *TreeManager) ToSql() string {
	return t.Visitor().Accept(t.Ast)
}

func (t *TreeManager) Visitor() Visitor {
	return t.Engine.Connection().Visitor()
}
