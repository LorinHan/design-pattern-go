package status

import (
	"fmt"
	"log"
)

type Status interface {
	Handle(ctx *StatusContext) error
}

type StatusContext struct {
	Status Status
}

func (s *StatusContext) Handle() error {
	return s.Status.Handle(s)
}

type StatusA struct {
}

func (s *StatusA) Handle(ctx *StatusContext) error {
	log.Println("status A 处理完毕")
	ctx.Status = &StatusB{}
	log.Println("status A => B")
	return nil
}

type StatusB struct {
	count int
}

func (s *StatusB) Handle(ctx *StatusContext) error {
	s.count ++
	if s.count % 2 != 0 {
		return fmt.Errorf("status B 处理失败")
	}

	ctx.Status = &StatusC{}
	log.Println("status B => C")
	return nil
}

type StatusC struct {
}

func (s *StatusC) Handle(ctx *StatusContext) error {
	log.Println("status C is the final status")
	return nil
}
