package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func tryMap() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}

	fmt.Println((m["Bell Labs"]))
	fmt.Println(m)

	m2 := map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m2)

	v, ok := m2["Google"]
	fmt.Println("The value:", v, "Present?", ok)
	_, ok = m2["Amazon"]
	fmt.Println("Amazon Present?", ok)
}

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	tryMap()

	wc := WordCount("I am learning Go! Go is a great language!")
	fmt.Println(wc)
}
