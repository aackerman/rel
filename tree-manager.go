package arel

type TreeManager struct {
	Ast    string
	Engine string
	ctx    *Context
}

func (t *TreeManager) ToSql() string {
	return t.Visitor().Accept(t.Ast)
}

func (t *TreeManager) Visitor() {
	return
}
