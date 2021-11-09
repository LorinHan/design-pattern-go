package builder

import (
	"design-pattern/builder"
	"fmt"
	"testing"
)

func TestBike(t *testing.T) {
	// 建造ofo单车
	d1 := &builder.Director{B: &builder.OfoBuilder{}}
	ofoBike := d1.Bike()
	fmt.Println(ofoBike)

	// 建造摩拜单车
	d2 := &builder.Director{B: &builder.MobikeBuilder{}}
	mobike := d2.Bike()
	fmt.Println(mobike)

	emptyDirector := &builder.Director{}
	fmt.Println(emptyDirector.Bike())
}
