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

type OtherFormatter func(space []byte, kv []interface{}) []byte

func (o OtherFormatter) Format(space []byte, kv []interface{}) []byte {
	return o(space, kv)
}
