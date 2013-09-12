package arel

type InsertManager struct {
	*TreeManager
}

func NewInsertManager(e Engine) *InsertManager {
	return &InsertManager{
		&TreeManager{
			Engine: e,
		},
	}
}
