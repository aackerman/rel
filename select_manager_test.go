package rel_test

import (
	. "."
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SelectManager", func() {
	It("has a skip method", func() {
		table := NewTable("users")
		manager := table.From(table)
		sql := manager.Skip(10).ToSql()
		expected := "SELECT FROM \"users\" OFFSET 10"
		Expect(sql).To(Equal(expected))
	})

	It("has an exists method", func() {
		table := NewTable("users")
		manager := table.From(table)
		manager.Project(Sql("*"))
		m2 := NewSelectManager(RelEngine, nil)
		m2.Project(manager.Exists())
		sql := m2.ToSql()
		expected := fmt.Sprintf("SELECT EXISTS (%s)", manager.ToSql())
		Expect(sql).To(Equal(expected))
	})

	It("has an offset method", func() {
		table := NewTable("users")
		manager := table.From(table)
		sql := manager.Offset(10).ToSql()
		expected := "SELECT FROM \"users\" OFFSET 10"
		Expect(sql).To(Equal(expected))
	})

	It("has a union method", func() {
		table := NewTable("users")
		m1 := NewSelectManager(RelEngine, table)
		m1.Project(Star())
		m1.Where(table.Attr("age").Lt(Sql(18)))
		m2 := NewSelectManager(RelEngine, table)
		m2.Project(Star())
		m2.Where(table.Attr("age").Gt(Sql(99)))
		mgr := m1.Union(m1.Ast, m2.Ast)
		sql := mgr.ToSql()
		expected := "( SELECT * FROM \"users\" WHERE \"users\".\"age\" < 18 UNION SELECT * FROM \"users\" WHERE \"users\".\"age\" > 99 )"
		Expect(sql).To(Equal(expected))
	})

	It("has a unionall method", func() {
		table := NewTable("users")
		m1 := NewSelectManager(RelEngine, table)
		m1.Project(Star())
		m1.Where(table.Attr("age").Lt(Sql(18)))
		m2 := NewSelectManager(RelEngine, table)
		m2.Project(Star())
		m2.Where(table.Attr("age").Gt(Sql(99)))
		mgr := m1.UnionAll(m1.Ast, m2.Ast)
		sql := mgr.ToSql()
		expected := "( SELECT * FROM \"users\" WHERE \"users\".\"age\" < 18 UNION ALL SELECT * FROM \"users\" WHERE \"users\".\"age\" > 99 )"
		Expect(sql).To(Equal(expected))
	})

	It("has a intersect method", func() {
		table := NewTable("users")
		m1 := NewSelectManager(RelEngine, table)
		m1.Project(Star())
		m1.Where(table.Attr("age").Lt(Sql(18)))
		m2 := NewSelectManager(RelEngine, table)
		m2.Project(Star())
		m2.Where(table.Attr("age").Gt(Sql(99)))
		mgr := m1.Intersect(m1.Ast, m2.Ast)
		sql := mgr.ToSql()
		expected := "( SELECT * FROM \"users\" WHERE \"users\".\"age\" < 18 INTERSECT SELECT * FROM \"users\" WHERE \"users\".\"age\" > 99 )"
		Expect(sql).To(Equal(expected))
	})

	It("has an except method", func() {
		table := NewTable("users")
		m1 := NewSelectManager(RelEngine, table)
		m1.Project(Star())
		m1.Where(table.Attr("age").Lt(Sql(99)))
		m2 := NewSelectManager(RelEngine, table)
		m2.Project(Star())
		m2.Where(table.Attr("age").Lt(Sql(50)))
		mgr := m1.Except(m1.Ast, m2.Ast)
		sql := mgr.ToSql()
		expected := "( SELECT * FROM \"users\" WHERE \"users\".\"age\" < 99 EXCEPT SELECT * FROM \"users\" WHERE \"users\".\"age\" < 50 )"
		Expect(sql).To(Equal(expected))
	})

	It("has an join method", func() {
		left := NewTable("users")
		right := left.Alias()
		predicate := left.Attr("id").Eq(right.Attr("id"))
		mgr := left.Select(Star()).Join(right).On(predicate)
		sql := mgr.ToSql()
		expected := "SELECT * FROM \"users\" INNER JOIN \"users\" \"users_2\" ON \"users\".\"id\" = \"users_2\".\"id\""
		Expect(sql).To(Equal(expected))
	})

	It("has an order method", func() {
		table := NewTable("users")
		mgr := table.Select(Star())
		mgr.Order(table.Attr("id"))
		sql := mgr.ToSql()
		expected := "SELECT * FROM \"users\" ORDER BY \"users\".\"id\""
		Expect(sql).To(Equal(expected))
	})

	It("has an order with direction method", func() {
		table := NewTable("users")
		mgr := table.Select(Star())
		mgr.Order(table.Attr("id").Desc())
		sql := mgr.ToSql()
		expected := "SELECT * FROM \"users\" ORDER BY \"users\".\"id\" DESC"
		Expect(sql).To(Equal(expected))
	})

	It("has an join method", func() {
		left := NewTable("users")
		right := left.Alias()
		predicate := left.Attr("id").Eq(right.Attr("id"))
		mgr := left.From(left)
		mgr.Join(right).On(predicate, predicate)
		sql := mgr.ToSql()
		expected := "SELECT FROM \"users\" INNER JOIN \"users\" \"users_2\" ON \"users\".\"id\" = \"users_2\".\"id\" AND \"users\".\"id\" = \"users_2\".\"id\""
		Expect(sql).To(Equal(expected))
	})

	It("has an count method", func() {
		table := NewTable("users")
		mgr := table.Select(Star())
		mgr.Order(table.Attr("id").Count().Desc())
		sql := mgr.ToSql()
		expected := "SELECT * FROM \"users\" ORDER BY COUNT(\"users\".\"id\") DESC"
		Expect(sql).To(Equal(expected))
	})

	It("has an lock method", func() {
		table := NewTable("users")
		mgr := table.From(table)
		sql := mgr.Lock(Sql("FOR SHARE")).ToSql()
		expected := "SELECT FROM \"users\" FOR SHARE"
		Expect(sql).To(Equal(expected))
	})

	It("has an lockforupdate method", func() {
		table := NewTable("users")
		mgr := table.From(table)
		sql := mgr.LockForUpdate().ToSql()
		expected := "SELECT FROM \"users\" FOR UPDATE"
		Expect(sql).To(Equal(expected))
	})

	It("has an group method", func() {
		users := NewTable("users")
		mgr := users.Group(users.Attr("id"))
		sql := mgr.ToSql()
		expected := "SELECT FROM \"users\" GROUP BY \"users\".\"id\""
		Expect(sql).To(Equal(expected))
	})

	It("has an window method", func() {
		users := NewTable("users")
		mgr := users.From(users)
		mgr.Window(Sql("a_window"))
		sql := mgr.ToSql()
		expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS ()"
		Expect(sql).To(Equal(expected))
	})

	It("has an window method that allows an order", func() {
		users := NewTable("users")
		mgr := users.From(users)
		mgr.Window(Sql("a_window")).Order(users.Attr("foo").Asc())
		sql := mgr.ToSql()
		expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (ORDER BY \"users\".\"foo\" ASC)"
		Expect(sql).To(Equal(expected))
	})

	It("has a window with unbounded preceding", func() {
		users := NewTable("users")
		mgr := users.From(users)
		mgr.Window(Sql("a_window")).Rows(&PrecedingNode{})
		sql := mgr.ToSql()
		expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (ROWS UNBOUNDED PRECEDING)"
		Expect(sql).To(Equal(expected))
	})

	It("has a window with bounded preceding", func() {
		users := NewTable("users")
		mgr := users.From(users)
		mgr.Window(Sql("a_window")).Rows(&PrecedingNode{Expr: Sql(5)})
		sql := mgr.ToSql()
		expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (ROWS 5 PRECEDING)"
		Expect(sql).To(Equal(expected))
	})

	It("has a window with unbounded following", func() {
		users := NewTable("users")
		mgr := users.From(users)
		mgr.Window(Sql("a_window")).Rows(&FollowingNode{})
		sql := mgr.ToSql()
		expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (ROWS UNBOUNDED FOLLOWING)"
		Expect(sql).To(Equal(expected))
	})

	It("has a window with bounded following", func() {
		users := NewTable("users")
		mgr := users.From(users)
		mgr.Window(Sql("a_window")).Rows(&FollowingNode{Expr: Sql(5)})
		sql := mgr.ToSql()
		expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (ROWS 5 FOLLOWING)"
		Expect(sql).To(Equal(expected))
	})

	It("has a window with current row", func() {
		users := NewTable("users")
		mgr := users.From(users)
		mgr.Window(Sql("a_window")).Rows(&CurrentRowNode{})
		sql := mgr.ToSql()
		expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (ROWS CURRENT ROW)"
		Expect(sql).To(Equal(expected))
	})

	It("has a window frame between two delimeters", func() {
		users := NewTable("users")
		mgr := users.From(users)
		window := mgr.Window(Sql("a_window"))
		window.Frame(&BetweenNode{
			Left:  window.Rows(nil),
			Right: mgr.NewAndNode(&PrecedingNode{}, &CurrentRowNode{}),
		})
		sql := mgr.ToSql()
		expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)"
		Expect(sql).To(Equal(expected))
	})

	It("has a window with range bounded preceding", func() {
		users := NewTable("users")
		mgr := users.From(users)
		window := mgr.Window(Sql("a_window"))
		window.Range(&PrecedingNode{})
		sql := mgr.ToSql()
		expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (RANGE UNBOUNDED PRECEDING)"
		Expect(sql).To(Equal(expected))
	})

	It("has a window with range bounded preceding", func() {
		users := NewTable("users")
		mgr := users.From(users)
		window := mgr.Window(Sql("a_window"))
		window.Range(&PrecedingNode{Expr: Sql(5)})
		sql := mgr.ToSql()
		expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (RANGE 5 PRECEDING)"
		Expect(sql).To(Equal(expected))
	})

	It("has a window with range unbounded following", func() {
		users := NewTable("users")
		mgr := users.From(users)
		window := mgr.Window(Sql("a_window"))
		window.Range(&FollowingNode{})
		sql := mgr.ToSql()
		expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (RANGE UNBOUNDED FOLLOWING)"
		Expect(sql).To(Equal(expected))
	})

	It("has a window with range bounded preceding", func() {
		users := NewTable("users")
		mgr := users.From(users)
		window := mgr.Window(Sql("a_window"))
		window.Range(&FollowingNode{Expr: Sql(5)})
		sql := mgr.ToSql()
		expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (RANGE 5 FOLLOWING)"
		Expect(sql).To(Equal(expected))
	})

	It("has a window with range current row", func() {
		users := NewTable("users")
		mgr := users.From(users)
		window := mgr.Window(Sql("a_window"))
		window.Range(&CurrentRowNode{})
		sql := mgr.ToSql()
		expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (RANGE CURRENT ROW)"
		Expect(sql).To(Equal(expected))
	})

	It("has a window with range between two delimeters", func() {
		users := NewTable("users")
		mgr := users.From(users)
		window := mgr.Window(Sql("a_window"))
		window.Frame(&BetweenNode{
			Left:  window.Range(nil),
			Right: mgr.NewAndNode(&PrecedingNode{}, &CurrentRowNode{}),
		})
		sql := mgr.ToSql()
		expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (RANGE BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)"
		Expect(sql).To(Equal(expected))
	})

	It("supports joins to multiple tables", func() {
		users := NewTable("users")
		comments := NewTable("comments")
		counts := comments.From(comments).Group(comments.Attr("user_id")).Project(
			comments.Attr("user_id").As(Sql("user_id")),
			comments.Attr("user_id").Count().As(Sql("count")),
		).As(Sql("counts"))
		mgr := users.Join(counts).On(counts.Attr("user_id").Eq(Sql(10)))
		sql := mgr.ToSql()
		expected := "SELECT FROM \"users\" INNER JOIN (SELECT \"comments\".\"user_id\" AS user_id, COUNT(\"comments\".\"user_id\") AS count FROM \"comments\" GROUP BY \"comments\".\"user_id\") counts ON counts.\"user_id\" = 10"
		Expect(sql).To(Equal(expected))
	})

	It("has a Distinct method", func() {
		users := NewTable("users")
		mgr := users.From(users)
		mgr.Distinct()
		sql := mgr.ToSql()
		expected := "SELECT DISTINCT FROM \"users\""
		Expect(sql).To(Equal(expected))
	})

	It("has a NotDistinct method", func() {
		users := NewTable("users")
		mgr := users.From(users)
		mgr.Distinct()
		mgr.NotDistinct()
		sql := mgr.ToSql()
		expected := "SELECT FROM \"users\""
		Expect(sql).To(Equal(expected))
	})

	It("has a WithRecursive method", func() {
		comments := NewTable("comments")
		commentsId := comments.Attr("id")
		commentsParentId := comments.Attr("parent_id")

		replies := NewTable("replies")
		repliesId := replies.Attr("id")

		recursiveTerm := NewSelectManager(RelEngine, nil)
		recursiveTerm.From(comments).Project(commentsId, commentsParentId).Where(commentsId.Eq(Sql(42)))

		nonRecursiveTerm := NewSelectManager(RelEngine, nil)
		nonRecursiveTerm.From(comments).Project(commentsId, commentsParentId).Join(replies).On(commentsParentId.Eq(repliesId))

		union := recursiveTerm.Union(recursiveTerm.Ast, nonRecursiveTerm.Ast)

		asStmt := &AsNode{Left: replies, Right: union}

		mgr := NewSelectManager(RelEngine, nil)
		mgr.WithRecursive(asStmt).From(replies).Project(Star())

		sql := mgr.ToSql()
		expected := "WITH RECURSIVE \"replies\" AS ( SELECT \"comments\".\"id\", \"comments\".\"parent_id\" FROM \"comments\" WHERE \"comments\".\"id\" = 42 UNION SELECT \"comments\".\"id\", \"comments\".\"parent_id\" FROM \"comments\" INNER JOIN \"replies\" ON \"comments\".\"parent_id\" = \"replies\".\"id\" ) SELECT * FROM \"replies\""
		Expect(sql).To(Equal(expected))
	})
})
