package rel

type InsertManager struct {
	Engine Engine
	Ast    *InsertStatementNode
	BaseVisitable
}

func NewInsertManager(engine Engine) *InsertManager {
	return &InsertManager{
		Engine: engine,
		Ast:    &InsertStatementNode{},
	}
}

func (mgr *InsertManager) ToSql() string {
	return mgr.Engine.Visitor().Accept(mgr.Ast)
}

func (mgr *InsertManager) Into(table *Table) *InsertManager {
	mgr.Ast.Relation = table
	return mgr
}

func (mgr *InsertManager) SetValues(values *ValuesNode) {
	mgr.Ast.Values = values
}

func (mgr *InsertManager) CreateValues(values []Visitable, columns []SqlLiteralNode) *ValuesNode {
	return &ValuesNode{
		Values:  values,
		Columns: columns,
	}
}

// func (mgr *InsertManager) Columns() {
// 	return mgr.Ast.columns
// }
