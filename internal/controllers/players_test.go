package controllers_test

import (
	"database/sql"
	"errors"
	"local/internal/controllers"
	"local/internal/logger"
	"local/internal/models"
	"local/internal/repository/mocks"
	"local/internal/repository/postgres"
	"log"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type PlayerControllerTestSuite struct {
	suite.Suite
	C            controllers.PlayerController
	valid_data   models.Player
	invalid_data models.Player
	valid_id     int
	invalid_id   int
}

func (suite *PlayerControllerTestSuite) SetupSuite() {
	file, err := os.OpenFile("test_logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger.Logger = log.New(file, "LOGS: ", log.Ldate|log.Ltime|log.Lshortfile)

	ctl := gomock.NewController(suite.T())
	companyTeamRep := mocks.NewMockCompanyTeamRepository(ctl)
	companyTournamentRep := mocks.NewMockCompanyTournamentRepository(ctl)
	companyRep := mocks.NewMockCompanyRepository(ctl)
	matchPerfomanceRep := mocks.NewMockMatchPerfomanceRepository(ctl)
	matchRep := mocks.NewMockMatchRepository(ctl)
	playerRep := mocks.NewMockPlayerRepository(ctl)
	teamPlayerRep := mocks.NewMockTeamPlayerRepository(ctl)
	teamRep := mocks.NewMockTeamRepository(ctl)
	tournamentTeamRep := mocks.NewMockTournamentTeamRepository(ctl)
	tournamentRep := mocks.NewMockTournamentRepository(ctl)
	userRep := mocks.NewMockUserRepository(ctl)

	suite.valid_data, _ = models.PlayerView{
		Id:            "1",
		Nickname:      "nickname",
		Realname:      "realname",
		Birthdate:     "2000-01-01",
		Country:       "country",
		MMR:           "1",
		Role:          "role",
		SignatureHero: "pudge"}.FromViewConv()
	suite.invalid_data, _ = models.PlayerView{
		Id:            "a",
		Nickname:      "nickname",
		Realname:      "realname",
		Birthdate:     "2000-01-01",
		Country:       "country",
		MMR:           "1",
		Role:          "role",
		SignatureHero: "pudge"}.FromViewConv()

	suite.valid_id = 1
	suite.invalid_id = -1

	playerRep.EXPECT().Add(suite.valid_data).Return(nil).AnyTimes()
	playerRep.EXPECT().Edit(suite.valid_data).Return(nil).AnyTimes()
	playerRep.EXPECT().Delete(suite.valid_id).Return(nil).AnyTimes()
	playerRep.EXPECT().Delete(suite.invalid_id).Return(errors.New("")).AnyTimes()
	playerRep.EXPECT().FindById(suite.valid_id).Return(make([]*models.Player, 1), nil).AnyTimes()
	playerRep.EXPECT().FindById(suite.invalid_id).Return(nil, errors.New("")).AnyTimes()
	playerRep.EXPECT().GetAll().Return(make([]*models.Player, 1), nil).AnyTimes()

	mockReps := mocks.MockRepositoriesBuilder{}.Build(nil,
		companyTeamRep,
		companyTournamentRep,
		companyRep,
		matchPerfomanceRep,
		matchRep,
		playerRep,
		teamPlayerRep,
		teamRep,
		tournamentTeamRep,
		tournamentRep,
		userRep)

	suite.C = controllers.PlayerControllerBuilder{}.Build(&mockReps)
}

func (suite *PlayerControllerTestSuite) TearDownSuite() {
}

func TestPlayerControllerTestSuite(t *testing.T) {
	suite.Run(t, new(PlayerControllerTestSuite))
}

func (suite *PlayerControllerTestSuite) TestPlayerController() {
	suite.Run("add_valid", func() {

		err := suite.C.Add(suite.valid_data)
		suite.NoError(err)
	})

	suite.Run("add_not_valid", func() {
		err := suite.C.Add(suite.invalid_data)
		suite.Error(err)
	})

	suite.Run("edit_valid", func() {
		err := suite.C.Edit(suite.valid_data)
		suite.NoError(err)
	})

	suite.Run("edit_not_valid", func() {
		err := suite.C.Edit(suite.invalid_data)
		suite.Error(err)
	})

	suite.Run("delete_by_id", func() {
		err := suite.C.Delete(suite.valid_id)
		suite.NoError(err)
	})

	suite.Run("delete_by_id_err", func() {
		err := suite.C.Delete(suite.invalid_id)
		suite.Error(err)
	})

	suite.Run("find_by_id", func() {
		obj, err := suite.C.FindById(suite.valid_id)
		suite.NoError(err)
		suite.NotNil(obj)
	})

	suite.Run("find_by_id_err", func() {
		obj, err := suite.C.FindById(suite.invalid_id)
		suite.Error(err)
		suite.Nil(obj)
	})

	suite.Run("get_all", func() {
		obj, err := suite.C.GetAll()
		suite.NoError(err)
		suite.NotNil(obj)
	})
}

type PlayerControllerIntegrationTestSuite struct {
	suite.Suite
	C          controllers.PlayerController
	db         *sql.DB
	valid_id   int
	invalid_id int
}

func (suite *PlayerControllerIntegrationTestSuite) SetupSuite() {
	file, err := os.OpenFile("test_logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger.Logger = log.New(file, "LOGS: ", log.Ldate|log.Ltime|log.Lshortfile)

	suite.db, _ = postgres.NewDB(postgres.TEST)

	if err := suite.db.Ping(); err != nil {
		return
	}

	pgRep := postgres.RepositoriesBuilder{}.Build(suite.db,
		&postgres.CompanyTeamRepository{DB: suite.db},
		&postgres.CompanyTournamentRepository{DB: suite.db},
		&postgres.CompanyRepository{DB: suite.db},
		&postgres.MatchPerfomanceRepository{DB: suite.db},
		&postgres.MatchRepository{DB: suite.db},
		&postgres.PlayerRepository{DB: suite.db},
		&postgres.TeamPlayerRepository{DB: suite.db},
		&postgres.TeamRepository{DB: suite.db},
		&postgres.TournamentTeamRepository{DB: suite.db},
		&postgres.TournamentRepository{DB: suite.db},
		&postgres.UserRepository{DB: suite.db})

	suite.valid_id = 1
	suite.invalid_id = -1

	suite.C = controllers.PlayerControllerBuilder{}.Build(&pgRep)
}

func (suite *PlayerControllerIntegrationTestSuite) TearDownSuite() {
	defer suite.db.Close()
}

func TestPlayerControllerIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(PlayerControllerIntegrationTestSuite))
}

func (suite *PlayerControllerIntegrationTestSuite) TestPlayerControllerIntegration() {
	suite.Run("find_by_id", func() {
		obj, err := suite.C.FindById(suite.valid_id)
		suite.NoError(err)
		suite.NotNil(obj)
	})

	suite.Run("find_by_id_err", func() {
		obj, err := suite.C.FindById(suite.invalid_id)
		suite.Error(err)
		suite.Nil(obj)
	})

	suite.Run("get_all", func() {
		obj, err := suite.C.GetAll()
		suite.NoError(err)
		suite.NotNil(obj)
	})
}
