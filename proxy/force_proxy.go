package proxy

import (
	"fmt"
)

/**
force_proxy 强制代理模式
调用者无法直接调用角色，必须通过角色的代理者来调用
*/

type IForceGamePlayer interface {
	Login(name, password string)
	KillBoss()
	UpGrade()
	GetProxy() IForceGamePlayer
}

type ForceGamePlayer struct {
	Name  string
	level int
	Proxy IForceGamePlayer
}

func (g *ForceGamePlayer) GetProxy() IForceGamePlayer {
	if g.Proxy == nil {
		g.Proxy = &ForceGamePlayerProxy{player: g} // 传入自己
	}
	return g.Proxy
}

func (g *ForceGamePlayer) Login(name, password string) {
	if g.IsProxy() {
		fmt.Printf("用户 %s 登陆成功！\n", name)
	} else {
		fmt.Println("请使用指定的代理访问")
	}
}

func (g *ForceGamePlayer) KillBoss() {
	if g.IsProxy() {
		fmt.Printf("%s正在打Boss\n", g.Name)
	} else {
		fmt.Println("请使用指定的代理访问")
	}
}

func (g *ForceGamePlayer) UpGrade() {
	if g.IsProxy() {
		g.level += 1
		fmt.Printf("%s升级了，当前等级为：%d\n", g.Name, g.level)
	} else {
		fmt.Println("请使用指定的代理访问")
	}
}

func (g *ForceGamePlayer) IsProxy() bool {
	if g.Proxy != nil {
		return true
	}
	return false
}

type ForceGamePlayerProxy struct {
	player IGamePlayer
}

func (gp *ForceGamePlayerProxy) Login(name, password string) {
	gp.player.Login(name, password)
}

func (gp *ForceGamePlayerProxy) KillBoss() {
	gp.player.KillBoss()
}

func (gp *ForceGamePlayerProxy) UpGrade() {
	gp.player.UpGrade()
}

// GetProxy 目前没有代理的代理，返回自己
func (gp *ForceGamePlayerProxy) GetProxy() IForceGamePlayer {
	return gp
}
