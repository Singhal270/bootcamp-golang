package main

type employee interface {
	salary() int
}

type FullTime struct {
	basis int
}
type Contract struct {
	basis int
}
type Freelance struct {
	basis int
	hour  int
}

func (f FullTime) salary() int {
	return f.basis * 28
}
func (c Contract) salary() int {
	return c.basis * 28
}
func (f Freelance) salary() int {
	return f.basis * f.hour
}

func main() {
	var emp1, emp2, emp3 employee
	emp1 = FullTime{500}
	emp2 = Contract{100}
	emp3 = Freelance{10, 50}

	println(emp1.salary())
	println(emp2.salary())
	println(emp3.salary())

}
