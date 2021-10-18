package factory

import (
	"design-pattern/factory"
	"fmt"
	"testing"
)

func TestFactory(t *testing.T) {
	h1 := factory.CreateHuman("chinese")
	h2 := factory.CreateHuman("american")

	fmt.Print("h1: ")
	h1.Talk()
	fmt.Print("h2: ")
	h2.Talk()
}
