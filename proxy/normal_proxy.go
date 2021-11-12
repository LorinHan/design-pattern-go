package proxy

import "fmt"

/**
normal_proxy 普通代理模式
调用者只知道代理，不知道真实角色是谁，屏蔽真实角色变更对高层模块的影响
 */

// 限制外部无法创建
type normalGamePlayer struct {
	name  string
	level int
}

func (g *normalGamePlayer) Login(name, password string) {
	fmt.Printf("用户 %s 登陆成功！\n", name)
}

func (g *normalGamePlayer) KillBoss() {
	fmt.Printf("%s正在打Boss\n", g.name)
}

func (g *normalGamePlayer) UpGrade() {
	g.level += 1
	fmt.Printf("%s升级了，当前等级为：%d\n", g.name, g.level)
}

type NormalGamePlayerProxy struct {
	player IGamePlayer
}

func (gp *NormalGamePlayerProxy) Init(name string) {
	gp.player = &normalGamePlayer{name: name}
}

func (gp *NormalGamePlayerProxy) Login(name, password string) {
	gp.player.Login(name, password)
}

func (gp *NormalGamePlayerProxy) KillBoss() {
	gp.player.KillBoss()
}

func (gp *NormalGamePlayerProxy) UpGrade() {
	gp.player.UpGrade()
}