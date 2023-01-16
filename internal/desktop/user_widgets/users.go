package user_widgets

import (
	"local/internal/managers"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func NewUserTab(w fyne.Window, m managers.Managers) *widget.TabContainer {
	editTab := container.NewAppTabs(
		container.NewTabItem("Добавить", (&AddUserBox{}).New(w, m)),
		container.NewTabItem("Редактировать", (&EditUserBox{}).New(w, m)),
		container.NewTabItem("Удалить", (&DeleteUserBox{}).New(w, m)),
		container.NewTabItem("Просмотр", (&ViewUserBox{}).New(w, m)),
	)
	return editTab
}
