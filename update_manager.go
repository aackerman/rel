package arel

type UpdateManager struct {
	BaseNode
}

func NewUpdateManager() *UpdateManager {
	return new(UpdateManager)
}
