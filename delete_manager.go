package arel

type DeleteManager struct {
	AstNode
}

func NewDeleteManager() *DeleteManager {
	return new(DeleteManager)
}
