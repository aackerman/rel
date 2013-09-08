package arel

type Attribute struct {
	Name  string
	Table *Table
}

func AttributeNew(t *Table, name string) *Attribute {
	return &Attribute{
		Name:  name,
		Table: t,
	}
}

func (a *Attribute) IsEqual(name string) *EqualityNode {
	return EqualityNodeNew()
}
