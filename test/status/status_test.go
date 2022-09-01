package status

import (
	"design-pattern/status"
	"log"
	"testing"
)

func TestStatus(t *testing.T) {
	statusCtx := &status.StatusContext{Status: &status.StatusA{}}

	for i := 0; i < 4; i++ {
		if err := statusCtx.Handle(); err != nil {
			log.Println("err:", err)
		}
	}
}
