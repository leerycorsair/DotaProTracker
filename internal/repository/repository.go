package repository

type Repositories interface {
	CompanyTeamRep() CompanyTeamRepository
	CompanyTournamentRep() CompanyTournamentRepository
	CompanyRep() CompanyRepository
	MatchPerfomanceRep() MatchPerfomanceRepository
	MatchRep() MatchRepository
	PlayerRep() PlayerRepository
	TeamPlayerRep() TeamPlayerRepository
	TeamRep() TeamRepository
	TournamentTeamRep() TournamentTeamRepository
	TournamentRep() TournamentRepository
	UserRep() UserRepository
}

type RepositoriesBuilder interface {
	Build() Repositories
}
