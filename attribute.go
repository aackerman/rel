package arel

type Attribute struct {
	Name  string
	Table *Table
}

func NewAttribute(name string, t *Table) Attribute {
	return Attribute{
		Name:  name,
		Table: t,
	}
}

func (a Attribute) Eq(b interface{}) EqualityNode {
	return NewEqualityNode(&a, Sql(b))
}
