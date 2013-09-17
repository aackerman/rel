package arel

type DeleteManager struct {
	Ast Node
	TreeManager
}

func NewDeleteManager() *DeleteManager {
	return &DeleteManager{
		Ast: &DeleteStatement{},
	}
}
