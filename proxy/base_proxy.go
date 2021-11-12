package proxy

import "fmt"

/**
base_proxy 基本的代理模式
proxy也可以嵌入after、before之类的函数实现aop
*/

type IGamePlayer interface {
	Login(name, password string)
	KillBoss()
	UpGrade()
}

type GamePlayer struct {
	Name  string
	level int
}

func (g *GamePlayer) Login(name, password string) {
	fmt.Printf("用户 %s 登陆成功！\n", name)
}

func (g *GamePlayer) KillBoss() {
	fmt.Printf("%s正在打Boss\n", g.Name)
}

func (g *GamePlayer) UpGrade() {
	g.level += 1
	fmt.Printf("%s升级了，当前等级为：%d\n", g.Name, g.level)
}

type GamePlayerProxy struct {
	Player IGamePlayer
}

func (gp *GamePlayerProxy) Login(name, password string) {
	gp.Player.Login(name, password)
}

func (gp *GamePlayerProxy) KillBoss() {
	gp.Player.KillBoss()
}

func (gp *GamePlayerProxy) UpGrade() {
	gp.Player.UpGrade()
}
