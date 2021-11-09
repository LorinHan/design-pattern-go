package template

import "fmt"

type reminder interface {
	GetUser(id int) string
	Send(user string) bool
	SendTo(id int)
}

// RemindTemp 提醒发送的模板类
type RemindTemp struct {
	Reminder reminder
}

func (rt *RemindTemp) SendTo(id int) {
	reminder := rt.Reminder
	user := reminder.GetUser(id)
	res := reminder.Send(user)
	fmt.Printf("记录日志，提醒用户：%s，提醒结果：%t\n", user, res)
	fmt.Printf("插入数据库，提醒用户：%s，插入结果：%t\n", user, res)
}

type EmailReminder struct {
	RemindTemp
}

func (e *EmailReminder) GetUser(id int) string {
	if id == 1 {
		return "2450978570@qq.com"
	}
	return "123456@163.com"
}

func (e *EmailReminder) Send(user string) bool {
	fmt.Println("发送邮件咯，目标：", user)
	fmt.Println("发送成功")
	return true
}

type PhoneReminder struct {
	RemindTemp
}

func (p *PhoneReminder) GetUser(id int) string {
	if id == 1 {
		return "15345922954"
	}
	return "10086"
}

func (p *PhoneReminder) Send(user string) bool {
	fmt.Println("发送短信咯，目标：", user)
	fmt.Println("发送成功")
	return true
}