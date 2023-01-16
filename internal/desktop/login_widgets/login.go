package login_widgets

import (
	"local/internal/controllers"
	"local/internal/desktop"
	"local/internal/logger"
	"local/internal/managers"
	"local/internal/models"
	"local/internal/repository/postgres"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

type LoginUserBox struct {
	mainWindow  fyne.Window
	userManager managers.UserManager
	dApp        desktop.DesktopApp
}

func (obj *LoginUserBox) New(app fyne.App, mainWindow fyne.Window, mans managers.Managers, dApp desktop.DesktopApp) *fyne.Container {
	obj.mainWindow = mainWindow
	obj.userManager = mans.UserMan
	label := widget.NewLabel("Авторизация")
	loginEntry := widget.NewEntry()
	loginEntry.SetPlaceHolder("Логин...")
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Пароль...")

	handler := func() {
		info := models.UserView{
			Id:             "1",
			Name:           "Unknown",
			Birthdate:      "2000-01-01",
			Login:          loginEntry.Text,
			Password:       passwordEntry.Text,
			Email:          "Unknown",
			PrivilegeLevel: "1",
		}
		priv_level, err := obj.userManager.Login(info)
		if err != nil {
			logger.Logger.Println(err)
			dialog.NewError(err, obj.mainWindow)
		} else {
			db, _ := postgres.NewDB(postgres.CONNECTIONS_POOL[priv_level-1])

			if err := db.Ping(); err != nil {
				logger.Logger.Println(err)
				return
			}
			pgRep := postgres.Repositories{DB: db}
			mans = managers.Managers{СompanyTeamMan: managers.CompanyTeamManager{C: controllers.CompanyTeamController{Reps: &pgRep}},
				CompanyTournamentMan: managers.CompanyTournamentManager{C: controllers.CompanyTournamentController{Reps: &pgRep}},
				CompanyMan:           managers.CompanyManager{C: controllers.CompanyController{Reps: &pgRep}},
				MatchPerfomanceMan:   managers.MatchPerfomanceManager{C: controllers.MatchPerfomanceController{Reps: &pgRep}},
				MatchMan:             managers.MatchManager{C: controllers.MatchController{Reps: &pgRep}},
				PlayerMan:            managers.PlayerManager{C: controllers.PlayerController{Reps: &pgRep}},
				TeamPlayerMan:        managers.TeamPlayerManager{C: controllers.TeamPlayerController{Reps: &pgRep}},
				TeamMan:              managers.TeamManager{C: controllers.TeamController{Reps: &pgRep}},
				TournamentTeamMan:    managers.TournamentTeamManager{C: controllers.TournamentTeamController{Reps: &pgRep}},
				TournamentMan:        managers.TournamentManager{C: controllers.TournamentController{Reps: &pgRep}},
				UserMan:              managers.UserManager{C: controllers.UserController{Reps: &pgRep}}}

			app := desktop.NewMainWindow(app, "DotaProTracker", mans)
			app.Init(priv_level)
			app.Show()
		}
	}

	btn := widget.NewButton("Войти", handler)

	box := container.NewCenter(container.NewVBox(label,
		loginEntry,
		passwordEntry,
		btn))
	return box
}
