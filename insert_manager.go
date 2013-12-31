package grel

type InsertManager struct {
	BaseNode
}

func NewInsertManager(t *Table) InsertManager {
	return InsertManager{}
}
