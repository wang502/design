package formatter

type Formatter interface {
	Format(space []byte, kv []interface{}) []byte
}

func doSomethingWithFormatter(f Formatter) {
}

type FuncFormatter func(space []byte, kv []interface{}) []byte

func (f FuncFormatter) Format(space []byte, kv []interface{}) []byte {
	return f(space, kv)
}
