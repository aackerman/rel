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
  users := rel.NewTable("users")
  manager := users.Select(rel.Sql("*"))
  fmt.Println(manager.ToSql()) // SELECT * FROM "users"
}
```

## Where

```
users.Where(users.Attr("name").Eq(rel.Sql("amy")))
// SELECT * FROM "users" WHERE "users"."name" = "amy"
```
