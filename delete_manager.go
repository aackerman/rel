package arel

type DeleteManager struct {
	Ast SqlStatement
	TreeManager
}

func NewDeleteManager() *DeleteManager {
	return &DeleteManager{
		Ast: DeleteStatement{},
	}
}
