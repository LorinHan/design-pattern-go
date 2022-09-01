package prototype

import "fmt"

/**
原型模式，适用于新建实例需要频繁的进行数据准备或者其他高消耗性操作
*/

type Email struct {
	Title    string
	Content  string
	Receiver string
}

func (e *Email) Clone() *Email {
	email := *e
	return &email
}

func (e *Email) Send() {
	fmt.Printf("%+v\n", e)
}

// Email2 测试深浅拷贝
type Email2 struct {
	Email
	Arr []int
}

func (e *Email2) Clone() *Email2 {
	email := *e
	copy(e.Arr, e.Arr)
	return &email
}

func (e *Email2) Send() {
	fmt.Printf("%+v\n", e)
	println(&e.Arr)
}