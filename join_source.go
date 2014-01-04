package rel

type JoinSource struct {
	Left  Visitable
	Right []Visitable
	BaseVisitable
}

type InnerJoinNode JoinNode
type OuterJoinNode JoinNode
