package arel

type InsertManager struct {
	*TreeManager
}

func InsertManagerNew(engine string) *InsertManager {
	return &InsertManager{
		&TreeManager{
			Engine: engine,
		},
	}
}
