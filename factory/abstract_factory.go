package factory

import "fmt"

type phone interface {
	Call()
	ConnectCharger(charger)
}

type charger interface {
	Charge()
}

// IPhone13 is phone的实现类
type IPhone13 struct {
}

func (i IPhone13) Call() {
	fmt.Println("IPhone13 is calling...")
}

func (i IPhone13) ConnectCharger(c charger) {
	fmt.Println("IPhone13 is charging...")
	c.Charge()
}

// HuaWeiP30 is phone的实现类 华为P30
type HuaWeiP30 struct {
}

func (h HuaWeiP30) Call() {
	fmt.Println("HuaWeiP30 is calling...")
}

func (h HuaWeiP30) ConnectCharger(c charger) {
	fmt.Println("HuaWeiP30 is charging...")
	c.Charge()
}

type LightningCharger struct {
}

func (l LightningCharger) Charge() {
	fmt.Println("Charge by Lightning Charger")
}

type TypeCCharger struct {
}

func (t TypeCCharger) Charge() {
	fmt.Println("Charge by Type-C Charger")
}

type phoneFactory interface {
	CreatePhone() phone
	CreateCharger() charger
}

type IPhone13Factory struct {
}

func (f IPhone13Factory) CreatePhone() phone {
	return &IPhone13{}
}

func (f IPhone13Factory) CreateCharger() charger {
	return &LightningCharger{}
}

type HuaWeiP30Factory struct {
}

func (h HuaWeiP30Factory) CreatePhone() phone {
	return &HuaWeiP30{}
}

func (h HuaWeiP30Factory) CreateCharger() charger {
	return &TypeCCharger{}
}

// GetPhone 根据类型获取不同的工厂，然后面向接口编程即可
func GetPhone(t int) (phone, charger) {
	var pf phoneFactory

	switch t {
	case 1:
		pf = &IPhone13Factory{}
	case 2:
		pf = &HuaWeiP30Factory{}
	default:
		return nil, nil
	}

	return pf.CreatePhone(), pf.CreateCharger()
}
