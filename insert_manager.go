package rel

type InsertManager struct {
	Ast *InsertStatementNode
	BaseVisitable
}

func NewInsertManager(engine Engine) *InsertManager {
	return &InsertManager{
		Ast: &InsertStatementNode{},
	}
}

func (mgr *InsertManager) Into(table *Table) *InsertManager {
	mgr.Ast.Relation = table
	return mgr
}

// func (mgr *InsertManager) Columns() {
// 	return mgr.Ast.columns
// }
