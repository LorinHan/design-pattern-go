package factory

import (
	"design-pattern/factory"
	"fmt"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	i13Factory := &factory.IPhone13Factory{}
	iPhone13 := i13Factory.CreatePhone()
	iPhone13.Call()
	lightning := i13Factory.CreateCharger()
	iPhone13.ConnectCharger(lightning)

	hwFactory := &factory.HuaWeiP30Factory{}
	hwP30 := hwFactory.CreatePhone()
	hwP30.Call()
	typeC := hwFactory.CreateCharger()
	hwP30.ConnectCharger(typeC)

	fmt.Println("========= GetPhoneAndDoFunc ==========")

	p1, c1 := factory.GetPhone(1)
	p1.Call()
	p1.ConnectCharger(c1)

	p2, c2 := factory.GetPhone(2)
	p2.Call()
	p2.ConnectCharger(c2)
}
