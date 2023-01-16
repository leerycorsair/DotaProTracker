package postgres_test

import (
	"database/sql"
	"local/internal/repository/postgres"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type UsersRepUnitTestSuite struct {
	suite.Suite
	db    *sql.DB
	pgRep postgres.Repositories
}

func (suite *UsersRepUnitTestSuite) SetupSuite() {
	suite.db, _ = postgres.NewDB(postgres.TEST)

	if err := suite.db.Ping(); err != nil {
		return
	}

	suite.pgRep = postgres.RepositoriesBuilder{}.Build(suite.db,
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
}

func (suite *UsersRepUnitTestSuite) TearDownSuite() {
	defer suite.db.Close()
}

func TestUsersRepUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UsersRepUnitTestSuite))
}

func (suite *UsersRepUnitTestSuite) TestUsersRepUnit() {
	suite.Run("find_by_id", func() {
		obj, err := suite.pgRep.UserRep().FindById(1)
		suite.NoError(err)
		suite.NotNil(obj)
	})

	suite.Run("find_by_id_err", func() {
		obj, err := suite.pgRep.UserRep().FindById(-1)
		suite.Error(err)
		suite.Nil(obj)
	})

	suite.Run("get_all", func() {
		obj, err := suite.pgRep.UserRep().GetAll()
		suite.NoError(err)
		suite.NotNil(obj)
	})
}
