package rel

// Handles specific cases where sql must be single quoted
// As is the case with matching statements such as LIKE and NOT LIKE
// as well as ILIKE and NOT ILIKE in PostgreSQL
type QuotedNode struct {
	Raw string
	BaseVisitable
}
