package arel

import (
	"log"
	"strconv"
)

type SqlLiteralNode struct {
	Raw string
	BaseNode
}

func Sql(raw interface{}) SqlLiteralNode {
	return NewSqlLiteralNode(raw)
}

func Star() SqlLiteralNode {
	return Sql("*")
}

func NewSqlLiteralNode(raw interface{}) SqlLiteralNode {
	var val string
	switch raw.(type) {
	case string:
		val = raw.(string)
	case int:
		val = strconv.Itoa(raw.(int))
	default:
		log.Fatalf("Cannot create SqlLiteralNode from input type %T", raw)
	}
	return SqlLiteralNode{Raw: val}
}
