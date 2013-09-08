package arel

type Context struct {
	Wheres []string
}

func ContextNew() *Context {
	return &Context{}
}
