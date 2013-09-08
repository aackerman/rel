package arel

import (
	"fmt"
	"testing"
)

func TestTableProject(t *testing.T) {
	users := TableNew("users")
	query := users.Project(arel.Star())
	fmt.PrintLn(query.ToSql())
}
