package main

import "fmt"

type Thing interface {
	Item() float64
	SetItem(float64)
}

func newThing() Thing {
	item := 0.0
	return struct {
		Item (func() float64)
		SetItem (func(float64))
	}{
		Item: func() float64 { return item },
		SetItem: func(x float64) { item = x },
	}
}

func main() {
	thing := newThing()
	fmt.Println("Hello, playground")
	fmt.Println(thing)
}