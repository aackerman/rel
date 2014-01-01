package rel

type InsertManager struct {
	BaseVisitable
}

func NewInsertManager(t *Table) *InsertManager {
	return &InsertManager{}
}
