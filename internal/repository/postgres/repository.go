package postgres

import (
	"database/sql"
	"local/internal/repository"
)

type Repositories struct {
	DB                   *sql.DB
	companyTeamRep       *CompanyTeamRepository
	companyTournamentRep *CompanyTournamentRepository
	companyRep           *CompanyRepository
	matchPerfomanceRep   *MatchPerfomanceRepository
	matchRep             *MatchRepository
	playerRep            *PlayerRepository
	teamPlayerRep        *TeamPlayerRepository
	teamRep              *TeamRepository
	tournamentTeamRep    *TournamentTeamRepository
	tournamentRep        *TournamentRepository
	userRep              *UserRepository
}

type RepositoriesBuilder struct {
}

func (b RepositoriesBuilder) Build(DB *sql.DB,
	companyTeamRep *CompanyTeamRepository,
	companyTournamentRep *CompanyTournamentRepository,
	companyRep *CompanyRepository,
	matchPerfomanceRep *MatchPerfomanceRepository,
	matchRep *MatchRepository,
	playerRep *PlayerRepository,
	teamPlayerRep *TeamPlayerRepository,
	teamRep *TeamRepository,
	tournamentTeamRep *TournamentTeamRepository,
	tournamentRep *TournamentRepository,
	userRep *UserRepository) Repositories {
	return Repositories{DB: DB,
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

func (r *Repositories) CompanyTeamRep() repository.CompanyTeamRepository {
	if r.companyTeamRep != nil {
		return r.companyTeamRep
	}
	r.companyTeamRep = &CompanyTeamRepository{DB: r.DB}
	return r.companyTeamRep
}

func (r *Repositories) CompanyTournamentRep() repository.CompanyTournamentRepository {
	if r.companyTournamentRep != nil {
		return r.companyTournamentRep
	}
	r.companyTournamentRep = &CompanyTournamentRepository{DB: r.DB}
	return r.companyTournamentRep
}

func (r *Repositories) CompanyRep() repository.CompanyRepository {
	if r.companyRep != nil {
		return r.companyRep
	}
	r.companyRep = &CompanyRepository{DB: r.DB}
	return r.companyRep
}

func (r *Repositories) MatchPerfomanceRep() repository.MatchPerfomanceRepository {
	if r.matchPerfomanceRep != nil {
		return r.matchPerfomanceRep
	}
	r.matchPerfomanceRep = &MatchPerfomanceRepository{DB: r.DB}
	return r.matchPerfomanceRep
}

func (r *Repositories) MatchRep() repository.MatchRepository {
	if r.matchRep != nil {
		return r.matchRep
	}
	r.matchRep = &MatchRepository{DB: r.DB}
	return r.matchRep
}

func (r *Repositories) PlayerRep() repository.PlayerRepository {
	if r.playerRep != nil {
		return r.playerRep
	}
	r.playerRep = &PlayerRepository{DB: r.DB}
	return r.playerRep
}

func (r *Repositories) TeamPlayerRep() repository.TeamPlayerRepository {
	if r.teamPlayerRep != nil {
		return r.teamPlayerRep
	}
	r.teamPlayerRep = &TeamPlayerRepository{DB: r.DB}
	return r.teamPlayerRep
}

func (r *Repositories) TeamRep() repository.TeamRepository {
	if r.teamRep != nil {
		return r.teamRep
	}
	r.teamRep = &TeamRepository{DB: r.DB}
	return r.teamRep
}

func (r *Repositories) TournamentTeamRep() repository.TournamentTeamRepository {
	if r.tournamentTeamRep != nil {
		return r.tournamentTeamRep
	}
	r.tournamentTeamRep = &TournamentTeamRepository{DB: r.DB}
	return r.tournamentTeamRep
}

func (r *Repositories) TournamentRep() repository.TournamentRepository {
	if r.tournamentRep != nil {
		return r.tournamentRep
	}
	r.tournamentRep = &TournamentRepository{DB: r.DB}
	return r.tournamentRep
}

func (r *Repositories) UserRep() repository.UserRepository {
	if r.userRep != nil {
		return r.userRep
	}
	r.userRep = &UserRepository{DB: r.DB}
	return r.userRep
}
