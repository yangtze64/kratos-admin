package expr

type (
	String string
	Option func(f String) String
)

func Convert(str string) String {
	return String(str)
}
func (f String) String() string {
	return f.toString()
}
func (f String) toString() string {
	return string(f)
}
func (f String) Expr(opts ...Option) String {
	return f.expr(opts...)
}
func (f String) expr(opts ...Option) String {
	for _, opt := range opts {
		f = opt(f)
	}
	return f
}
