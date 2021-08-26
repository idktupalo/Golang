package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func stepNode(t *tree.Tree, ch chan int) {
	stepNodeRecursive(t, ch)
	close(ch)
}

func stepNodeRecursive(t *tree.Tree, ch chan int) {
	if t != nil {
		stepNodeRecursive(t.Left, ch)
		ch <- t.Value
		stepNodeRecursive(t.Right, ch)
	}
}

func same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go stepNode(t1, ch1)
	go stepNode(t2, ch2)
	for {
		node1, check1 := <-ch1
		node2, check2 := <-ch2
		if check1 != check2 || node1 != node2 {
			return false
		}
		if !check1 {
			break
		}
	}
	return true
}

func main() {
	fmt.Println("t:1-1 -> ", same(tree.New(1), tree.New(1)))
	fmt.Println("t:2-1 -> ", same(tree.New(2), tree.New(1)))
	fmt.Println("t:1-3 -> ", same(tree.New(1), tree.New(3)))
}
