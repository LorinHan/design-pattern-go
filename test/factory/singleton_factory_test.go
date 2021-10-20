package factory

import (
	"design-pattern/factory"
	"testing"
)

func TestSingletonFactory_create(t *testing.T) {
	println(factory.SF.Single())
	println(factory.SF.Single())
	println(factory.SF.Single())
}