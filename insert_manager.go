package rel

import (
	"log"
)

type InsertManager struct {
	Engine Engine
	Ast    *InsertStatementNode
	BaseVisitable
}

func NewInsertManager(engine Engine) *InsertManager {
	if engine == nil {
		log.Fatal("Cannot accept a nil Engine")
	}
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

func (mgr *InsertManager) Insert(column AttributeNode, value interface{}) *InsertManager {
	if mgr.Ast.Values == nil {
		mgr.Ast.Values = &ValuesNode{
			Values:  make([]interface{}, 0),
			Columns: make([]AttributeNode, 0),
		}
	}

	if mgr.Ast.Columns == nil {
		slice := make([]AttributeNode, 0)
		mgr.Ast.Columns = &slice
	}
	*mgr.Ast.Columns = append(*mgr.Ast.Columns, column)
	mgr.Ast.Values.Columns = append(mgr.Ast.Values.Columns, column)
	mgr.Ast.Values.Values = append(mgr.Ast.Values.Values, value)
	return mgr
}

func (mgr *InsertManager) SetValues(values *ValuesNode) {
	mgr.Ast.Values = values
}

func (mgr *InsertManager) CreateValues(values []interface{}, columns []AttributeNode) *ValuesNode {
	return &ValuesNode{
		Values:  values,
		Columns: columns,
	}
}

// func (mgr *InsertManager) Columns() {
// 	return mgr.Ast.columns
// }
