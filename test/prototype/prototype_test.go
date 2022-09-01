package prototype

import (
	"design-pattern/prototype"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestPrototype(t *testing.T) {
	email := &prototype.Email{
		Title:   "标题",
		Content: "内容",
	}

	// 如果send开协程，会存在数据竞争
	for i := 0; i < 5; i++ {
		email.Receiver = strconv.Itoa(i) + "先生"
		email.Send()
	}

	fmt.Println("======================")

	for i := 0; i < 5; i++ {
		clone := email.Clone()
		clone.Receiver = strconv.Itoa(i) + "先生"
		go clone.Send()
	}

	time.Sleep(time.Second * 3)
}

// 测试深浅拷贝
func TestPrototype2(t *testing.T) {
	email := &prototype.Email2{
		Email: prototype.Email{
			Title:   "标题",
			Content: "内容",
		},
		Arr: make([]int, 0, 5),
	}

	for i := 0; i < 5; i++ {
		clone := email.Clone()
		clone.Receiver = strconv.Itoa(i) + "先生"
		clone.Arr = append(clone.Arr, i)  // arr是浅拷贝，虽然打印出来的arr指针地址不同，但实际上指向的底层数组地址相同，因此循环打印出来的都是最后一次改动的数值
		go clone.Send()
	}

	time.Sleep(time.Second * 2)
}