package arel

type Attribute struct {
	Name  string
	Table *Table
}

func NewAttribute(t *Table, name string) *Attribute {
	return &Attribute{
		Name:  name,
		Table: t,
	}
}

func (a *Attribute) IsEqual(name string) *EqualityNode {
	return NewEqualityNode()
}
