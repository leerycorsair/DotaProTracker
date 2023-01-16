package edit_widgets

import (
	"local/internal/managers"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func NewEditTab(w fyne.Window, m managers.Managers) *widget.TabContainer {
	editTab := container.NewAppTabs(
		container.NewTabItem("Игроки", (&EditPlayerBox{}).New(w, m)),
		container.NewTabItem("Команды", (&EditTeamBox{}).New(w, m)),
		container.NewTabItem("Турниры", (&EditTournamentBox{}).New(w, m)),
		container.NewTabItem("Компании", (&EditCompanyBox{}).New(w, m)),
		container.NewTabItem("Матчи", (&EditMatchBox{}).New(w, m)),
		container.NewTabItem("Компании-Команды", (&EditCompanyTeamBox{}).New(w, m)),
		container.NewTabItem("Компании-Турниры", (&EditCompanyTournamentBox{}).New(w, m)),
		container.NewTabItem("Матчи Подробно", (&EditMatchPerfomanceBox{}).New(w, m)),
		container.NewTabItem("Команды-Игроки", (&EditTeamPlayerBox{}).New(w, m)),
		container.NewTabItem("Турниры-Команды", (&EditTournamentTeamBox{}).New(w, m)),
	)
	return editTab
}
