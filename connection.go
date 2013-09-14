package arel

type Connection struct {
	visitor     *ToSqlVisitor
	Tables      []string
	primaryKeys []string
	columns     []string
}

func NewConnection(v *ToSqlVisitor) *Connection {
	return &Connection{
		visitor: v,
	}
}

func (c *Connection) TableExists(tableName string) bool {
	for _, name := range c.Tables {
		if tableName == name {
			return true
		}
	}
	return false
}

func QuoteTableName(name string) string {
	return "\"" + name + "\""
}

func QuoteColumnName(name string) string {
	return "\"" + name + "\""
}

func (c *Connection) Visitor() *ToSqlVisitor {
	return c.visitor
}

func (c *Connection) PrimaryKey(name string) {

}
