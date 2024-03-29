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

type UserControllerUnitTestSuite struct {
	suite.Suite
	C            controllers.UserController
	valid_data   models.User
	invalid_data models.User
	valid_id     int
	invalid_id   int
}

func (suite *UserControllerUnitTestSuite) SetupSuite() {
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

	suite.valid_data, _ = models.UserView{
		Id:             "1",
		Name:           "name",
		Birthdate:      "2000-01-01",
		Login:          "login",
		Password:       "password",
		Email:          "email",
		PrivilegeLevel: "1"}.FromViewConv()

	suite.invalid_data, _ = models.UserView{
		Id:             "a",
		Name:           "name",
		Birthdate:      "2000-01-01",
		Login:          "login",
		Password:       "password",
		Email:          "email",
		PrivilegeLevel: "1"}.FromViewConv()

	suite.valid_id = 1
	suite.invalid_id = -1

	userRep.EXPECT().Delete(suite.valid_id).Return(nil).AnyTimes()
	userRep.EXPECT().Delete(suite.invalid_id).Return(errors.New("")).AnyTimes()
	userRep.EXPECT().FindById(suite.valid_id).Return(make([]*models.User, 1), nil).AnyTimes()
	userRep.EXPECT().FindById(suite.invalid_id).Return(nil, errors.New("")).AnyTimes()
	userRep.EXPECT().GetAll().Return(make([]*models.User, 1), nil).AnyTimes()

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

	suite.C = controllers.UserControllerBuilder{}.Build(&mockReps)
}

func (suite *UserControllerUnitTestSuite) TearDownSuite() {
}

func TestUserControllerUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerUnitTestSuite))
}

func (suite *UserControllerUnitTestSuite) TestUserControllerTestSuiteUnit() {
	suite.Run("add_not_valid", func() {
		err := suite.C.Add(suite.invalid_data)
		suite.Error(err)
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

type UserControllerIntegrationTestSuite struct {
	suite.Suite
	C          controllers.UserController
	db         *sql.DB
	valid_id   int
	invalid_id int
}

func (suite *UserControllerIntegrationTestSuite) SetupSuite() {
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

	suite.C = controllers.UserControllerBuilder{}.Build(&pgRep)
}

func (suite *UserControllerIntegrationTestSuite) TearDownSuite() {
	defer suite.db.Close()
}

func TestUserControllerIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerIntegrationTestSuite))
}

func (suite *UserControllerIntegrationTestSuite) TestUserControllerIntegration() {
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
