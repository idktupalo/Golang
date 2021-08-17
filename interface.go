package main

import "fmt"

type A struct {
	aRow string
}

type B struct {
	bRow string
}

type C interface {
	show()
}

func (a A) show() {
	fmt.Println(a.aRow)
}

func (b B) show() {
	fmt.Println(b.bRow)
}

func showAll(cAll ...C) (res []C) {
	for _, val := range cAll {
		res = append(res, val)
	}
	return
}

func main() {
	a := A{"stringA"}
	b := B{"stringB"}
	a.show()
	b.show()
	showAll(a, b)
}