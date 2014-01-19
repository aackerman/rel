package rel

import (
	"reflect"
)

type InsertStatementNode struct {
	Relation *Table
	Columns  *[]*AttributeNode
	Values   *ValuesNode
	BaseVisitable
}

func (mgr *InsertStatementNode) Eq(other *InsertStatementNode) bool {
	return reflect.DeepEqual(mgr, other)
}
