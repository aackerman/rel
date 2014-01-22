package rel_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AttributeNode", func() {
	var users *Table
	var mgr *SelectManager

	BeforeEach(func() {
		Register(NewTestEngine())
		users = NewTable("users")
		mgr = users.Select(users.Attr("id"))
	})

	It("implements Predicator", func() {
		// compile time test
		var _ Predicator = &AttributeNode{}
	})

	It("implements Orderer", func() {
		// compile time test
		var _ Orderer = &AttributeNode{}
	})

	It("can use the NotEq predication", func() {
		mgr.Where(users.Attr("id").NotEq(Sql(10)))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" != 10"
		Expect(sql).To(Equal(expected))
	})

	It("can use the NotEqAny predication", func() {
		mgr.Where(users.Attr("id").NotEqAny(Sql(1), Sql(2)))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" != 1 OR \"users\".\"id\" != 2)"
		Expect(sql).To(Equal(expected))
	})

	It("can use the NotEq predication allowing for nil input", func() {
		mgr.Where(users.Attr("id").NotEq(nil))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" IS NOT NULL"
		Expect(sql).To(Equal(expected))
	})

	It("can use the NotEqAll predication", func() {
		mgr.Where(users.Attr("id").NotEqAll(Sql(1), Sql(2)))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" != 1 AND \"users\".\"id\" != 2)"
		Expect(sql).To(Equal(expected))
	})

	It("can use the Gt predication", func() {
		mgr.Where(users.Attr("id").Gt(Sql(10)))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" > 10"
		Expect(sql).To(Equal(expected))
	})

	It("can use the GtAny predication", func() {
		mgr.Where(users.Attr("id").GtAny(Sql(1), Sql(2)))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" > 1 OR \"users\".\"id\" > 2)"
		Expect(sql).To(Equal(expected))
	})

	It("can use the GtAll predication", func() {
		mgr.Where(users.Attr("id").GtAll(Sql(1), Sql(2)))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" > 1 AND \"users\".\"id\" > 2)"
		Expect(sql).To(Equal(expected))
	})

	It("can use the GtEq predication", func() {
		mgr.Where(users.Attr("id").GtEq(Sql(10)))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" >= 10"
		Expect(sql).To(Equal(expected))
	})

	It("can use the GtEqAny predication", func() {
		mgr.Where(users.Attr("id").GtEqAny(Sql(1), Sql(2)))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" >= 1 OR \"users\".\"id\" >= 2)"
		Expect(sql).To(Equal(expected))
	})

	It("can use the GtEqAll predication", func() {
		mgr.Where(users.Attr("id").GtEqAll(Sql(1), Sql(2)))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" >= 1 AND \"users\".\"id\" >= 2)"
		Expect(sql).To(Equal(expected))
	})

	It("can use the Lt predication", func() {
		mgr.Where(users.Attr("id").Lt(Sql(10)))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" < 10"
		Expect(sql).To(Equal(expected))
	})

	It("can use the LtAny predication", func() {
		mgr.Where(users.Attr("id").LtAny(Sql(1), Sql(2)))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" < 1 OR \"users\".\"id\" < 2)"
		Expect(sql).To(Equal(expected))
	})

	It("can use the LtAll predication", func() {
		mgr.Where(users.Attr("id").LtAll(Sql(1), Sql(2)))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" < 1 AND \"users\".\"id\" < 2)"
		Expect(sql).To(Equal(expected))
	})

	It("can use the LtEq predication", func() {
		mgr.Where(users.Attr("id").LtEq(Sql(10)))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" <= 10"
		Expect(sql).To(Equal(expected))
	})

	It("can use the LtEqAny predication", func() {
		mgr.Where(users.Attr("id").LtEqAny(Sql(1), Sql(2)))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" <= 1 OR \"users\".\"id\" <= 2)"
		Expect(sql).To(Equal(expected))
	})

	It("can use the LtEqAll predication", func() {
		mgr.Where(users.Attr("id").LtEqAll(Sql(1), Sql(2)))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" <= 1 AND \"users\".\"id\" <= 2)"
		Expect(sql).To(Equal(expected))
	})

	It("can use the Count predication", func() {
		mgr := users.Select(users.Attr("id").Count())
		sql := mgr.ToSql()
		expected := "SELECT COUNT(\"users\".\"id\") FROM \"users\""
		Expect(sql).To(Equal(expected))
	})

	It("can use the Eq predication", func() {
		mgr.Where(users.Attr("id").Eq(Sql(10)))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" = 10"
		Expect(sql).To(Equal(expected))
	})

	It("can use the Eq predication with nil input", func() {
		mgr.Where(users.Attr("id").Eq(nil))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" IS NULL"
		Expect(sql).To(Equal(expected))
	})

	It("can use the EqAny predication", func() {
		mgr.Where(users.Attr("id").EqAny(Sql(1), Sql(2)))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" = 1 OR \"users\".\"id\" = 2)"
		Expect(sql).To(Equal(expected))
	})

	It("can use the EqAll predication", func() {
		mgr.Where(users.Attr("id").EqAll(Sql(1), Sql(2)))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" = 1 AND \"users\".\"id\" = 2)"
		Expect(sql).To(Equal(expected))
	})

	It("can use the Matches predication", func() {
		mgr.Where(users.Attr("name").Matches(Sql("%bacon%")))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"name\" LIKE '%bacon%'"
		Expect(sql).To(Equal(expected))
	})

	It("can use the MatchesAny predication", func() {
		mgr.Where(users.Attr("name").MatchesAny(Sql("%chunky%"), Sql("%bacon%")))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"name\" LIKE '%chunky%' OR \"users\".\"name\" LIKE '%bacon%')"
		Expect(sql).To(Equal(expected))
	})

	It("can use the MatchesAll predication", func() {
		mgr.Where(users.Attr("name").MatchesAll(Sql("%chunky%"), Sql("%bacon%")))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"name\" LIKE '%chunky%' AND \"users\".\"name\" LIKE '%bacon%')"
		Expect(sql).To(Equal(expected))
	})

	It("can use the DoesNotMatch predication", func() {
		mgr.Where(users.Attr("name").DoesNotMatch(Sql("%bacon%")))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"name\" NOT LIKE '%bacon%'"
		Expect(sql).To(Equal(expected))
	})

	It("can use the DoesNotMatchAny predication", func() {
		mgr.Where(users.Attr("name").DoesNotMatchAny(Sql("%chunky%"), Sql("%bacon%")))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"name\" NOT LIKE '%chunky%' OR \"users\".\"name\" NOT LIKE '%bacon%')"
		Expect(sql).To(Equal(expected))
	})

	It("can use the DoesNotMatchAll predication", func() {
		mgr.Where(users.Attr("name").DoesNotMatchAll(Sql("%chunky%"), Sql("%bacon%")))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"name\" NOT LIKE '%chunky%' AND \"users\".\"name\" NOT LIKE '%bacon%')"
		Expect(sql).To(Equal(expected))
	})

	It("can use the Asc predication", func() {
		mgr.Order(users.Attr("id").Asc())
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" ORDER BY \"users\".\"id\" ASC"
		Expect(sql).To(Equal(expected))
	})

	It("can use the NotEq predication", func() {
		mgr.Order(users.Attr("id").Desc())
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" ORDER BY \"users\".\"id\" DESC"
		Expect(sql).To(Equal(expected))
	})

	It("can use the In predication", func() {
		mgr.Where(users.Attr("id").In([]Visitable{Sql(1), Sql(2), Sql(3)}))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" IN (1, 2, 3)"
		Expect(sql).To(Equal(expected))
	})

	It("can use the InAny predication", func() {
		mgr.Where(users.Attr("id").InAny([]Visitable{Sql(1), Sql(2)}, []Visitable{Sql(3), Sql(4)}))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" IN (1, 2) OR \"users\".\"id\" IN (3, 4))"
		Expect(sql).To(Equal(expected))
	})

	It("can use the InAll predication", func() {
		mgr.Where(users.Attr("id").InAll([]Visitable{Sql(1), Sql(2)}, []Visitable{Sql(3), Sql(4)}))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" IN (1, 2) AND \"users\".\"id\" IN (3, 4))"
		Expect(sql).To(Equal(expected))
	})

	It("can use the NotIn predication", func() {
		mgr.Where(users.Attr("id").NotIn([]Visitable{Sql(1), Sql(2), Sql(3)}))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" NOT IN (1, 2, 3)"
		Expect(sql).To(Equal(expected))
	})

	It("can use the NotInAny predication", func() {
		mgr.Where(users.Attr("id").NotInAny([]Visitable{Sql(1), Sql(2)}, []Visitable{Sql(3), Sql(4)}))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" NOT IN (1, 2) OR \"users\".\"id\" NOT IN (3, 4))"
		Expect(sql).To(Equal(expected))
	})

	It("can use the NotInAll predication", func() {
		mgr.Where(users.Attr("id").NotInAll([]Visitable{Sql(1), Sql(2)}, []Visitable{Sql(3), Sql(4)}))
		sql := mgr.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" NOT IN (1, 2) AND \"users\".\"id\" NOT IN (3, 4))"
		Expect(sql).To(Equal(expected))
	})
})
