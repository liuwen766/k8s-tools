package main

import "fmt"

// 比如，针对不同的事件等级，发邮件、发短信、打电话等等

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
