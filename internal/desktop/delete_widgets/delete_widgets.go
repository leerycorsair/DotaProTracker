package delete_widgets

import (
	"local/internal/managers"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func NewDeleteTab(w fyne.Window, m managers.Managers) *widget.TabContainer {
	delTab := container.NewAppTabs(
		container.NewTabItem("Игроки", (&DeletePlayerBox{}).New(w, m)),
		container.NewTabItem("Команды", (&DeleteTeamBox{}).New(w, m)),
		container.NewTabItem("Турниры", (&DeleteTournamentBox{}).New(w, m)),
		container.NewTabItem("Компании", (&DeleteCompanyBox{}).New(w, m)),
		container.NewTabItem("Матчи", (&DeleteMatchBox{}).New(w, m)),
		container.NewTabItem("Компании-Команды", (&DeleteCompanyTeamBox{}).New(w, m)),
		container.NewTabItem("Компании-Турниры", (&DeleteCompanyTournamentBox{}).New(w, m)),
		container.NewTabItem("Матчи Подробно", (&DeleteMatchPerfomanceBox{}).New(w, m)),
		container.NewTabItem("Команды-Игроки", (&DeleteTeamPlayerBox{}).New(w, m)),
		container.NewTabItem("Турниры-Команды", (&DeleteTournamentTeamBox{}).New(w, m)),
	)
	return delTab
}
