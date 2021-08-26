package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func stepNode(t *tree.Tree, ch chan int) {
	recursiveStepNode(t, ch)
	close(ch)
}

func recursiveStepNode(t *tree.Tree, ch chan int) {
	if t != nil {
		recursiveStepNode(t.Left, ch)
		ch <- t.Value
		recursiveStepNode(t.Right, ch)
	}
}

func same(tree1, tree2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go stepNode(tree1, ch1)
	go stepNode(tree2, ch2)
	for {
		n1, check1 := <-ch1
		n2, check2 := <-ch2
		if check1 != check2 || n1 != n2 {
			//fmt.Println("no equal")
			return false
		}
		if !check1 {
			break
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go stepNode(tree.New(1), ch)
	fmt.Println(same(tree.New(1), tree.New(2)))
	fmt.Println(same(tree.New(1), tree.New(1)))
	fmt.Println(same(tree.New(2), tree.New(1)))
}
