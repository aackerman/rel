package arel

type UpdateManager struct {
	AstNode
}

func NewUpdateManager() *UpdateManager {
	return new(UpdateManager)
}
