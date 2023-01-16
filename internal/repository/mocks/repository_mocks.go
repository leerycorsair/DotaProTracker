package mocks

import (
	"database/sql"
	"local/internal/repository"
)

type MockRepositories struct {
	DB                   *sql.DB
	companyTeamRep       *MockCompanyTeamRepository
	companyTournamentRep *MockCompanyTournamentRepository
	companyRep           *MockCompanyRepository
	matchPerfomanceRep   *MockMatchPerfomanceRepository
	matchRep             *MockMatchRepository
	playerRep            *MockPlayerRepository
	teamPlayerRep        *MockTeamPlayerRepository
	teamRep              *MockTeamRepository
	tournamentTeamRep    *MockTournamentTeamRepository
	tournamentRep        *MockTournamentRepository
	userRep              *MockUserRepository
}

type MockRepositoriesBuilder struct {
}

func (b MockRepositoriesBuilder) Build(DB *sql.DB,
	companyTeamRep *MockCompanyTeamRepository,
	companyTournamentRep *MockCompanyTournamentRepository,
	companyRep *MockCompanyRepository,
	matchPerfomanceRep *MockMatchPerfomanceRepository,
	matchRep *MockMatchRepository,
	playerRep *MockPlayerRepository,
	teamPlayerRep *MockTeamPlayerRepository,
	teamRep *MockTeamRepository,
	tournamentTeamRep *MockTournamentTeamRepository,
	tournamentRep *MockTournamentRepository,
	userRep *MockUserRepository) MockRepositories {
	return MockRepositories{DB: DB,
		companyTeamRep:       companyTeamRep,
		companyTournamentRep: companyTournamentRep,
		companyRep:           companyRep,
		matchPerfomanceRep:   matchPerfomanceRep,
		matchRep:             matchRep,
		playerRep:            playerRep,
		teamPlayerRep:        teamPlayerRep,
		teamRep:              teamRep,
		tournamentTeamRep:    tournamentTeamRep,
		tournamentRep:        tournamentRep,
		userRep:              userRep}
}

func (r *MockRepositories) CompanyTeamRep() repository.CompanyTeamRepository {
	if r.companyTeamRep != nil {
		return r.companyTeamRep
	}
	r.companyTeamRep = &MockCompanyTeamRepository{}
	return r.companyTeamRep
}

func (r *MockRepositories) CompanyTournamentRep() repository.CompanyTournamentRepository {
	if r.companyTournamentRep != nil {
		return r.companyTournamentRep
	}
	r.companyTournamentRep = &MockCompanyTournamentRepository{}
	return r.companyTournamentRep
}

func (r *MockRepositories) CompanyRep() repository.CompanyRepository {
	if r.companyRep != nil {
		return r.companyRep
	}
	r.companyRep = &MockCompanyRepository{}
	return r.companyRep
}

func (r *MockRepositories) MatchPerfomanceRep() repository.MatchPerfomanceRepository {
	if r.matchPerfomanceRep != nil {
		return r.matchPerfomanceRep
	}
	r.matchPerfomanceRep = &MockMatchPerfomanceRepository{}
	return r.matchPerfomanceRep
}

func (r *MockRepositories) MatchRep() repository.MatchRepository {
	if r.matchRep != nil {
		return r.matchRep
	}
	r.matchRep = &MockMatchRepository{}
	return r.matchRep
}

func (r *MockRepositories) PlayerRep() repository.PlayerRepository {
	if r.playerRep != nil {
		return r.playerRep
	}
	r.playerRep = &MockPlayerRepository{}
	return r.playerRep
}

func (r *MockRepositories) TeamPlayerRep() repository.TeamPlayerRepository {
	if r.teamPlayerRep != nil {
		return r.teamPlayerRep
	}
	r.teamPlayerRep = &MockTeamPlayerRepository{}
	return r.teamPlayerRep
}

func (r *MockRepositories) TeamRep() repository.TeamRepository {
	if r.teamRep != nil {
		return r.teamRep
	}
	r.teamRep = &MockTeamRepository{}
	return r.teamRep
}

func (r *MockRepositories) TournamentTeamRep() repository.TournamentTeamRepository {
	if r.tournamentTeamRep != nil {
		return r.tournamentTeamRep
	}
	r.tournamentTeamRep = &MockTournamentTeamRepository{}
	return r.tournamentTeamRep
}

func (r *MockRepositories) TournamentRep() repository.TournamentRepository {
	if r.tournamentRep != nil {
		return r.tournamentRep
	}
	r.tournamentRep = &MockTournamentRepository{}
	return r.tournamentRep
}

func (r *MockRepositories) UserRep() repository.UserRepository {
	if r.userRep != nil {
		return r.userRep
	}
	r.userRep = &MockUserRepository{}
	return r.userRep
}
