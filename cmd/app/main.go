package main

import (
	"fmt"
	"local/internal/controllers"
	"local/internal/desktop"
	"local/internal/desktop/login_widgets"
	"local/internal/logger"
	"local/internal/managers"
	"local/internal/repository/postgres"
	"log"
	"os"

	"fyne.io/fyne/app"
	_ "github.com/lib/pq"
)

const DESKTOP_APP int = 0
const WEB_APP int = 1

const APP_TITLE string = "DotaProTracker"
const APP_TYPE int = DESKTOP_APP

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger.Logger = log.New(file, "LOGS: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	db, _ := postgres.NewDB(postgres.INITIAL_LOGIN)
	defer db.Close()

	if err := db.Ping(); err != nil {
		logger.Logger.Println(err)
		return
	}

	a := app.New()

	pgRep := postgres.RepositoriesBuilder{}.Build(db,
		&postgres.CompanyTeamRepository{DB: db},
		&postgres.CompanyTournamentRepository{DB: db},
		&postgres.CompanyRepository{DB: db},
		&postgres.MatchPerfomanceRepository{DB: db},
		&postgres.MatchRepository{DB: db},
		&postgres.PlayerRepository{DB: db},
		&postgres.TeamPlayerRepository{DB: db},
		&postgres.TeamRepository{DB: db},
		&postgres.TournamentTeamRepository{DB: db},
		&postgres.TournamentRepository{DB: db},
		&postgres.UserRepository{DB: db})

	companyTeamController := controllers.CompanyTeamControllerBuilder{}.Build(&pgRep)
	companyTournamentController := controllers.CompanyTournamentControllerBuilder{}.Build(&pgRep)
	companyController := controllers.CompanyControllerBuilder{}.Build(&pgRep)
	matchPerfomanceController := controllers.MatchPerfomanceControllerBuilder{}.Build(&pgRep)
	matchController := controllers.MatchControllerBuilder{}.Build(&pgRep)
	playerController := controllers.PlayerControllerBuilder{}.Build(&pgRep)
	teamPlayerController := controllers.TeamPlayerControllerBuilder{}.Build(&pgRep)
	teamController := controllers.TeamControllerBuilder{}.Build(&pgRep)
	tournamentTeamController := controllers.TournamentTeamControllerBuilder{}.Build(&pgRep)
	tournamentController := controllers.TournamentControllerBuilder{}.Build(&pgRep)
	userController := controllers.UserControllerBuilder{}.Build(&pgRep)

	companyTeamManager := managers.CompanyTeamManagerBuilder{}.Build(companyTeamController)
	companyTournamentManager := managers.CompanyTournamentManagerBuilder{}.Build(companyTournamentController)
	companyManager := managers.CompanyManagerBuilder{}.Build(companyController)
	matchPerfomanceManager := managers.MatchPerfomanceManagerBuilder{}.Build(matchPerfomanceController)
	matchManager := managers.MatchManagerBuilder{}.Build(matchController)
	playerManager := managers.PlayerManagerBuilder{}.Build(playerController)
	teamPlayerManager := managers.TeamPlayerManagerBuilder{}.Build(teamPlayerController)
	teamManager := managers.TeamManagerBuilder{}.Build(teamController)
	tournamentTeamManager := managers.TournamentTeamManagerBuilder{}.Build(tournamentTeamController)
	tournamentManager := managers.TournamentManagerBuilder{}.Build(tournamentController)
	userManager := managers.UserManagerBuilder{}.Build(userController)

	managers := managers.ManagersBuilder{}.Build(companyTeamManager,
		companyTournamentManager,
		companyManager,
		matchPerfomanceManager,
		matchManager,
		playerManager,
		teamPlayerManager,
		teamManager,
		tournamentTeamManager,
		tournamentManager,
		userManager)

	if APP_TYPE == DESKTOP_APP {
		mainW := desktop.NewMainWindow(a, APP_TITLE, managers)
		loginW := login_widgets.NewLoginWindow(a, APP_TITLE, managers, mainW)
		loginW.Init()
		loginW.Show()
		a.Run()
	} else if APP_TYPE == WEB_APP {
		fmt.Println("TODO: WEB")
	}
}
