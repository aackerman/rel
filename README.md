# Rel

A SQL AST manager for Go.

Simplifies the generation of complex SQL queries

This library could fulfill the query generation for a great ORM.

Allows you to programatically create dynamic SQL with a nice interface

## Usage

```go
package main

import (
  "fmt"
  "rel"
)

func main() {
  relation := rel.NewTable("users")
  manager := relation.Select(Sql("*"))
  fmt.Println(manager.ToSql())
}
```
