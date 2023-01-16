package repository

import "local/internal/models"

//go:generate mockgen -source=./interfaces.go -destination=./mocks/interfaces_mocks.go -package=mocks

type CompanyTeamRepository interface {
	Add(models.CompanyTeam) error
	Edit(models.CompanyTeam) error
	Delete(int) error
	FindById(int) ([]*models.CompanyTeam, error)
	GetAll() ([]*models.CompanyTeam, error)
}

type CompanyTournamentRepository interface {
	Add(models.CompanyTournament) error
	Edit(models.CompanyTournament) error
	Delete(int) error
	FindById(int) ([]*models.CompanyTournament, error)
	GetAll() ([]*models.CompanyTournament, error)
}

type CompanyRepository interface {
	Add(models.Company) error
	Edit(models.Company) error
	Delete(int) error
	FindById(int) ([]*models.Company, error)
	FindByName(string) ([]*models.Company, error)
	GetAll() ([]*models.Company, error)
}

type MatchPerfomanceRepository interface {
	Add(models.MatchPerfomance) error
	Edit(models.MatchPerfomance) error
	Delete(int) error
	FindById(int) ([]*models.MatchPerfomance, error)
	FindByHero(string) ([]*models.MatchPerfomance, error)
	GetAll() ([]*models.MatchPerfomance, error)
}

type MatchRepository interface {
	Add(models.Match) error
	Edit(models.Match) error
	Delete(int) error
	FindById(int) ([]*models.Match, error)
	GetAll() ([]*models.Match, error)
}

type PlayerRepository interface {
	Add(models.Player) error
	Edit(models.Player) error
	Delete(int) error
	FindById(int) ([]*models.Player, error)
	FindByBirthdate(int) ([]*models.Player, error)
	FindByName(string) ([]*models.Player, error)
	GetAll() ([]*models.Player, error)
}

type TeamPlayerRepository interface {
	Add(models.TeamPlayer) error
	Edit(models.TeamPlayer) error
	Delete(int) error
	FindById(int) ([]*models.TeamPlayer, error)
	GetAll() ([]*models.TeamPlayer, error)
}

type TeamRepository interface {
	Add(models.Team) error
	Edit(models.Team) error
	Delete(int) error
	FindById(int) ([]*models.Team, error)
	FindByName(string) ([]*models.Team, error)
	GetAll() ([]*models.Team, error)
}

type TournamentTeamRepository interface {
	Add(models.TournamentTeam) error
	Edit(models.TournamentTeam) error
	Delete(int) error
	FindById(int) ([]*models.TournamentTeam, error)
	GetAll() ([]*models.TournamentTeam, error)
}

type TournamentRepository interface {
	Add(models.Tournament) error
	Edit(models.Tournament) error
	Delete(int) error
	FindById(int) ([]*models.Tournament, error)
	FindByName(string) ([]*models.Tournament, error)
	FindByYear(int) ([]*models.Tournament, error)
	GetAll() ([]*models.Tournament, error)
}

type UserRepository interface {
	Add(models.User) error
	Edit(models.User) error
	Delete(int) error
	FindById(int) ([]*models.User, error)
	FindByLogin(string) ([]*models.User, error)
	GetAll() ([]*models.User, error)
}
