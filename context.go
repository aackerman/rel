package arel

type Context struct {
	Wheres []string
}

func NewContext() *Context {
	return &Context{}
}
