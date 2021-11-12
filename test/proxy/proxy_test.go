package proxy

import (
	"design-pattern/proxy"
	"fmt"
	"testing"
)

func TestBaseProxy(t *testing.T) {
	player := &proxy.GamePlayer{Name: "Lorin"}
	p := &proxy.GamePlayerProxy{Player: player}

	p.Login("lorin", "123456")
	p.KillBoss()
	p.UpGrade()

	p.KillBoss()
	p.UpGrade()
}

func TestNormalProxy(t *testing.T) {
	p := &proxy.NormalGamePlayerProxy{}
	p.Init("Lorin")

	p.Login("lorin", "123456")
	p.KillBoss()
	p.UpGrade()

	p.KillBoss()
	p.UpGrade()
}

func TestForceProxy(t *testing.T) {
	player := &proxy.ForceGamePlayer{Name: "Lorin"}
	// 无法访问
	player.Login("lorin", "123456")
	player.KillBoss()
	player.UpGrade()

	fmt.Println("==================================")

	// 获取player的代理，通过代理访问
	p := player.GetProxy()
	p.Login("lorin", "123456")
	p.KillBoss()
	p.UpGrade()
}