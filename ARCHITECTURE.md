# Rel Architecture

The `Table` is the most central component to a SQL statement. From a table we can create several different types of statements.

```
* SELECT
* UPDATE
* DELETE
* INSERT
* UNION
* UNION ALL
* EXCEPT
* INTERSECT
```

A Table is used as a starting interface. From several of the Table instance methods a SelectManager is returned. Or an UpdateManager, an InsertManager, DeleteManger, or a MultiStatementManager. Each with a specific purpose to add visitable nodes to the structure of the query.
