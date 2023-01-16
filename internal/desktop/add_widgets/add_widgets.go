package add_widgets

import (
	"local/internal/managers"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func NewAddTab(w fyne.Window, m managers.Managers) *widget.TabContainer {
	addTab := container.NewAppTabs(
		container.NewTabItem("Игроки", (&AddPlayerBox{}).New(w, m)),
		container.NewTabItem("Команды", (&AddTeamBox{}).New(w, m)),
		container.NewTabItem("Турниры", (&AddTournamentBox{}).New(w, m)),
		container.NewTabItem("Компании", (&AddCompanyBox{}).New(w, m)),
		container.NewTabItem("Матчи", (&AddMatchBox{}).New(w, m)),
		container.NewTabItem("Компании-Команды", (&AddCompanyTeamBox{}).New(w, m)),
		container.NewTabItem("Компании-Турниры", (&AddCompanyTournamentBox{}).New(w, m)),
		container.NewTabItem("Матчи Подробно", (&AddMatchPerfomanceBox{}).New(w, m)),
		container.NewTabItem("Команды-Игроки", (&AddTeamPlayerBox{}).New(w, m)),
		container.NewTabItem("Турниры-Команды", (&AddTournamentTeamBox{}).New(w, m)),
	)
	return addTab
}
