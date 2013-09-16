package arel

type TreeManager struct {
	Engine *Engine
	ctx    Node
	Ast    SqlStatement
}

func (t *TreeManager) ToSql() string {
	return t.Visitor().Accept(t.Ast)
}

func (t *TreeManager) Visitor() Visitor {
	return t.Engine.Connection().Visitor()
}

func (t *TreeManager) Where(expr string) *TreeManager {
	append(c.ctx.Wheres, expr)
	return t
}
