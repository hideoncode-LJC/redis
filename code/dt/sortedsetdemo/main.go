package main

type Student struct {
	ID int
	Name string
	Age int
	Sex string
	Score int
}

type Class struct {
	Title string
	Students []*Student
}

func main() {

}

