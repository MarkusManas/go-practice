package scoping

const scoped_var string = "TEST ME"

func Foo() {
	var scoped_var string = "HaHA I was declared inside foo"
	println(scoped_var)
}
