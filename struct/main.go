package main

import "fmt"

func tryStruct() {
	p1 := Person{}
	p1.Name = "John"
	p1.Age = 30

	fmt.Println(p1.Greet())
	p1.GrowUp()

	fmt.Println(p1.Name, p1.Age)
	ToGrowUp(&p1)
	fmt.Println(p1.Name, p1.Age)
}

func tryComposition() {
	e1 := Employee{}
	e1.Name = "Jane"
	e1.Age = 25
	e1.EmployeeID = "1234"

	fmt.Println(e1.Greet())
	e1.GrowUp()

	fmt.Println(e1.Name, e1.Age)
	ToGrowUp(&e1)
	fmt.Println(e1.Name, e1.Age)
}

func tryPolymorphism() {
	d := Dog{"Fido"}
	c := Cat{"Whiskers"}
	fmt.Println(MakeNoise(&d))
	fmt.Println(MakeNoise(&c))
}

func main() {
	fmt.Println("Main function")
	// tryStruct()
	// tryComposition()

	tryPolymorphism()
}
