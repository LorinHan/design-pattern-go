package factory

import "fmt"

type humanBase struct {
	Color string
	Language string
}

type humanI interface {
	Talk()
}

type chinese struct {
	humanBase
}

func (c *chinese) Talk() {
	fmt.Printf("我是%s种人，我说%s\n", c.Color, c.Language)
}

type american struct {
	humanBase
}

func (a *american) Talk() {
	fmt.Printf("my skin is %s. I speak %s\n", a.Color, a.Language)
}