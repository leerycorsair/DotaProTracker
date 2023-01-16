package managers_test

import (
	"database/sql"
	"local/internal/controllers"
	"local/internal/logger"
	"local/internal/managers"
	"local/internal/models"
	"local/internal/repository/postgres"
	"log"
	"os"
	"strconv"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type CompanyManagerE2ETestSuite struct {
	suite.Suite
	M          managers.CompanyManager
	db         *sql.DB
	valid_data models.CompanyView
	valid_name string
}

func (suite *CompanyManagerE2ETestSuite) SetupSuite() {
	file, err := os.OpenFile("test_logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger.Logger = log.New(file, "LOGS: ", log.Ldate|log.Ltime|log.Lshortfile)

	suite.db, _ = postgres.NewDB(postgres.TEST)

	if err := suite.db.Ping(); err != nil {
		return
	}

	suite.valid_data = models.CompanyView{
		Id:       "2",
		Name:     "name",
		Country:  "country",
		Website:  "website.com",
		Revenue:  "1",
		Industry: "industry"}
	suite.valid_name = "name"

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

	companyController := controllers.CompanyControllerBuilder{}.Build(&pgRep)
	suite.M = managers.CompanyManagerBuilder{}.Build(companyController)
}

func (suite *CompanyManagerE2ETestSuite) TearDownSuite() {
	defer suite.db.Close()

}

func TestCompanyManagerE2ETestSuite(t *testing.T) {
	suite.Run(t, new(CompanyManagerE2ETestSuite))
}

func (suite *CompanyManagerE2ETestSuite) TestCompanyManagerE2E() {
	suite.Run("e2e test", func() {

		err := suite.M.Add(suite.valid_data)
		suite.NoError(err)

		obj, err := suite.M.FindByName(suite.valid_name)
		suite.NotNil(obj)
		suite.NotEmpty(obj)
		suite.NoError(err)

		for _, value := range obj {
			id, _ := strconv.Atoi(value.Id)
			err = suite.M.C.Delete(id)
			suite.NoError(err)
		}

		obj, err = suite.M.FindByName(suite.valid_name)
		suite.Empty(obj)
		suite.NoError(err)
	})
}
