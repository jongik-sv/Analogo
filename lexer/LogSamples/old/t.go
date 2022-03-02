package main

type Stringer interface {
	String() string
}

type Student struct {
}

func (s Student) String() string {
	println("hello")
	return "hello"
}

func main() {
	student := Student{}

	var stringer Stringer

	stringer = student

	stringer.String()

}
