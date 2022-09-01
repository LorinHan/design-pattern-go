package medium

import "fmt"

type Stock struct {
	computerNum int
}

func (s *Stock) Increase(num int) {
	s.computerNum += num
	fmt.Println("库存数量为：", s.computerNum)
}

func (s *Stock) Decrease(num int) {
	s.computerNum -= num
	fmt.Println("库存数量为：", s.computerNum)
}

func (s *Stock) GetComputerNum() int {
	return s.computerNum
}

// Clear 清仓销售，通知采购部停止采购，通知销售部折扣卖出
func (s *Stock) Clear() {

}
