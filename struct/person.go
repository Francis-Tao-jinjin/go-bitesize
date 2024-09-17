package main

type Person struct {
	Name string
	Age  int
}

func (p *Person) Greet() string {
	return "Hello, " + p.Name
}

func (p *Person) GrowUp() {
	p.Age++
}

func ToGrowUp(p Growable) {
	p.GrowUp()
}
