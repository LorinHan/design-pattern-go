package singleton

import (
	"design-pattern/template"
	"testing"
)

func TestTemplate(t *testing.T) {
	reminder := template.GetReminder("email")
	reminder.SendTo(1)

	reminder = template.GetReminder("shortMsg")
	reminder.SendTo(1)
}
