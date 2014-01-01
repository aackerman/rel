package rel

type AsNode struct {
	Left  Visitable
	Right *Visitable
	BaseNode
}
