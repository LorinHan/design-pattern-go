package singleton

import (
	"design-pattern/singleton"
	"testing"
)

func TestSingleton(t *testing.T) {
	println("方式1")
	s := singleton.S
	s2 := singleton.S
	println(s)
	println(s2)
	println(singleton.S)

	println("方式2")
	println(singleton.GetSingleton())
	println(singleton.GetSingleton())

	println("方式2.2")
	println(singleton.GetSingletonByLock())
	println(singleton.GetSingletonByLock())

	println("方式2.3")
	println(singleton.GetSingletonByOnce())
	println(singleton.GetSingletonByOnce())
}
