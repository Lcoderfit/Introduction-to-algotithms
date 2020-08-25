package main

import (
	"GoLeetcode/stackAndQueueTag/stackAndQueue"
	"fmt"
)

func main() {
	//225.用队列实现栈.go
	myStack()
}

func myStack() {
	stack := stackAndQueue.Constructor()
	fmt.Println("Init stack: ", stack)
	stack.Push(3)
	fmt.Println("Push 3: ", stack)

	fmt.Println("Is Empty: ", stack.Empty())
	fmt.Println("stack Top: ", stack.Top())
	fmt.Println("stack Pop: ", stack.Pop())
	fmt.Println("Final stack: ", stack)
	fmt.Println("Is Empty: ", stack.Empty())
}
