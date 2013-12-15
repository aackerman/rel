package arel

type DeleteManager struct {
	Ast AstNode
	TreeManager
}

func NewDeleteManager() *DeleteManager {
	return &DeleteManager{
		Ast: DeleteStatement{},
	}
}
