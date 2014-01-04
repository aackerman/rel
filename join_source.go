package rel

type JoinSource struct {
	Left  Visitable
	Right []Visitable
	BaseVisitable
}
