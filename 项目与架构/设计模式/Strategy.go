package main

import "fmt"

//比如，假设我们需要对 a、b 这两个整数进行计算，根据条件的不同，需要执行不同的计算方式。
//我们可以把所有的操作都封装在同一个函数中，然后通过 if … else … 的形式来调用不同的计算方式，这种方式称之为硬编码。

func main() {
	operator := Operator{}
	//使用加法策略
	operator.setStrategy(&add{})
	result := operator.calculate(1, 2)
	fmt.Println("add:", result)
	//使用减法策略
	operator.setStrategy(&reduce{})
	result = operator.calculate(2, 1)
	fmt.Println("reduce:", result)
}

// 策略模式
// 定义一个策略类
type IStrategy interface {
	do(int, int) int
}

// 策略实现：加
type add struct{}

func (*add) do(a, b int) int {
	return a + b
}

// 策略实现：减
type reduce struct{}

func (*reduce) do(a, b int) int {
	return a - b
}

// 具体策略的执行者
type Operator struct {
	strategy IStrategy
}

// 设置策略
func (operator *Operator) setStrategy(strategy IStrategy) {
	operator.strategy = strategy
}

// 调用策略中的方法
func (operator *Operator) calculate(a, b int) int {
	return operator.strategy.do(a, b)
}
