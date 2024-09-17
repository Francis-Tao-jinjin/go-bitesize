package main

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d *Dog) Speak() string {
	return "Woof! Woof!"
}

type Cat struct {
	Name string
}

func (c *Cat) Speak() string {
	return "Meow! Meow!"
}

func MakeNoise(s Speaker) string {
	return s.Speak()
}
