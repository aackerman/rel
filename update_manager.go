package rel

type UpdateManager struct {
	Ast UpdateStatementNode
	BaseVisitable
}

func NewUpdateManager() UpdateManager {
	return UpdateManager{}
}
