package singleton

import (
	"design-pattern/template"
	"testing"
)

func TestTemplate(t *testing.T) {
	reminder := &template.RemindTemp{
		Reminder: &template.EmailReminder{},
	}
	reminder.SendTo(1)

	reminder.Reminder = &template.PhoneReminder{}
	reminder.SendTo(1)
}
