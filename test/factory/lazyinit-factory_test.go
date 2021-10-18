package factory

import (
	"design-pattern/factory"
	"testing"
)

func TestLazyInitFactory_create(t *testing.T) {
	println(factory.LF.Get("product1"))
	println(factory.LF.Get("product1"))
	println(factory.LF.Get("product1"))

	println(factory.LF.Get("product2"))
	println(factory.LF.Get("product2"))
	println(factory.LF.Get("product2"))

	println(factory.LF.Get("product1"))
}