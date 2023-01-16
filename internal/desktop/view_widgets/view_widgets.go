package view_widgets

import (
	"local/internal/managers"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func NewViewTab(w fyne.Window, m managers.Managers) *widget.TabContainer {
	viewTab := container.NewAppTabs(
		container.NewTabItem("Игроки", (&ViewPlayerBox{}).New(w, m)),
		container.NewTabItem("Команды", (&ViewTeamBox{}).New(w, m)),
		container.NewTabItem("Турниры", (&ViewTournamentBox{}).New(w, m)),
		container.NewTabItem("Компании", (&ViewCompanyBox{}).New(w, m)),
		container.NewTabItem("Матчи", (&ViewMatchBox{}).New(w, m)),
		container.NewTabItem("Компании-Команды", (&ViewCompanyTeamBox{}).New(w, m)),
		container.NewTabItem("Компании-Турниры", (&ViewCompanyTournamentBox{}).New(w, m)),
		container.NewTabItem("Матчи Подробно", (&ViewMatchPerfomanceBox{}).New(w, m)),
		container.NewTabItem("Команды-Игроки", (&ViewTeamPlayerBox{}).New(w, m)),
		container.NewTabItem("Турниры-Команды", (&ViewTournamentTeamBox{}).New(w, m)),
	)
	return viewTab
}
