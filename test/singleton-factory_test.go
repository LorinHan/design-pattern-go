package test

import (
	"design-parttern/factory"
	"testing"
)

func TestSingletonFactory_create(t *testing.T) {
	println(factory.SF.Single())
	println(factory.SF.Single())
	println(factory.SF.Single())
}