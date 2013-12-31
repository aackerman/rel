package rel

type UpdateManager struct {
	Ast UpdateStatementNode
	BaseNode
}

func NewUpdateManager() UpdateManager {
	return UpdateManager{}
}
