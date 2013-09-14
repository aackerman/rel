package arel

import (
	"fmt"
)

func TableNaiveTest() {
	engine := Engine{}
	users := NewTable("users", engine)
	query := users.Select(Sql("*"))
	fmt.Println(query.ToSql)
}
