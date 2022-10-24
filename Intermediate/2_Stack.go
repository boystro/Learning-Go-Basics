package main

import "fmt"

type Stack10 struct {
	__arr [10]int
	__top int
}

func (self *Stack10) Init() {
	self.__top = -1
	self.__arr = [10]int{}
}

func (self *Stack10) push(val int) {
	if self.__top > 9 {
		fmt.Println("[ ERROR | Stack Overflow ]")
		return;
	}
	self.__top++
	self.__arr[self.__top] = val	
}

func (self *Stack10) top() int {
	return self.__arr[self.__top]
}

func (self *Stack10) pop() (val int) {
	if self.__top < 0 {
		fmt.Println("[ ERROR | No elements present ]")
		return 0
	}
	val = self.__arr[self.__top]
	self.__arr[self.__top] = 0
	self.__top--
	return val
}

func (self *Stack10) log() {
	fmt.Println(*self)
}

func main() {
	stack := Stack10{}
	stack.Init()

	stack.push(1)
	stack.log()
	stack.push(2)
	stack.log()
	stack.pop()
	stack.log()
	stack.push(3)
	stack.log()

}