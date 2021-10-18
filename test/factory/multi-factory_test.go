package factory

import (
	"design-pattern/factory"
	"testing"
)

func TestChineseFactory_create(t *testing.T) {
	cf := &factory.ChineseFactory{}
	c := cf.Create()
	c.Talk()

	af := &factory.AmericanFactory{}
	a := af.Create()
	a.Talk()
}