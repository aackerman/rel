package arel

import (
	"fmt"
)

func TableNaiveTest() {
	users := TableNew("users", "postgresql")
	query := users.Select(Sql("*"))
	fmt.Println(query.ToSql)
}
