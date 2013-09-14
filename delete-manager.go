package arel

type DeleteManager struct {
	TreeManager
}

func NewDeleteManager() *DeleteManager {
	return &DeleteManager{}
}
