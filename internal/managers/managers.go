package managers

type Managers struct {
	СompanyTeamMan       CompanyTeamManager
	CompanyTournamentMan CompanyTournamentManager
	CompanyMan           CompanyManager
	MatchPerfomanceMan   MatchPerfomanceManager
	MatchMan             MatchManager
	PlayerMan            PlayerManager
	TeamPlayerMan        TeamPlayerManager
	TeamMan              TeamManager
	TournamentTeamMan    TournamentTeamManager
	TournamentMan        TournamentManager
	UserMan              UserManager
}

type ManagersBuilder struct {
}

func (b ManagersBuilder) Build(companyTeamMan CompanyTeamManager,
	companyTournamentMan CompanyTournamentManager,
	companyMan CompanyManager,
	matchPerfomanceMan MatchPerfomanceManager,
	matchMan MatchManager,
	playerMan PlayerManager,
	teamPlayerMan TeamPlayerManager,
	teamMan TeamManager,
	tournamentTeamMan TournamentTeamManager,
	tournamentMan TournamentManager,
	userMan UserManager) Managers {
	return Managers{СompanyTeamMan: companyTeamMan,
		CompanyTournamentMan: companyTournamentMan,
		CompanyMan:           companyMan,
		MatchPerfomanceMan:   matchPerfomanceMan,
		MatchMan:             matchMan,
		PlayerMan:            playerMan,
		TeamPlayerMan:        teamPlayerMan,
		TeamMan:              teamMan,
		TournamentTeamMan:    tournamentTeamMan,
		TournamentMan:        tournamentMan,
		UserMan:              userMan}
}
