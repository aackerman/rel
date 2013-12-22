package arel

type InsertManager struct {
	AstNode
}

func NewInsertManager(e *Engine) *InsertManager {
	return new(InsertManager)
}
