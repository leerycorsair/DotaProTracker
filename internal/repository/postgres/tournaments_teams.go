package postgres

import (
	"database/sql"
	"local/internal/models"
	"local/internal/repository"
)

type TournamentTeamRepository struct {
	DB *sql.DB
}

func (r *TournamentTeamRepository) Add(mod models.TournamentTeam) error {
	return r.DB.QueryRow("insert into tournaments_teams  (tournament_id, team_id, participation_type, is_winner) values ($1, $2, $3, $4) returning id",
		mod.TournamentId, mod.TeamId, mod.ParticipationType, mod.IsWinner).Scan(&mod.Id)
}

func (r *TournamentTeamRepository) Edit(mod models.TournamentTeam) error {
	_, err := r.DB.Exec("update tournaments_teams set tournament_id=$1, team_id=$2, participation_type=$3, is_winner=$4 where id=$5",
		mod.TournamentId, mod.TeamId, mod.ParticipationType, mod.IsWinner, mod.Id)
	return err
}

func (r *TournamentTeamRepository) Delete(id int) error {
	_, err := r.DB.Exec("delete from tournaments_teams where id=$1", id)
	return err
}

func (r *TournamentTeamRepository) FindById(id int) ([]*models.TournamentTeam, error) {
	mod := models.TournamentTeam{}
	err := r.DB.QueryRow("select id, tournament_id, team_id, participation_type, is_winner from tournaments_teams where id=$1", id).Scan(
		&mod.Id,
		&mod.TournamentId,
		&mod.TeamId,
		&mod.ParticipationType,
		&mod.IsWinner,
	)
	recs := []*models.TournamentTeam{}
	recs = append(recs, &mod)
	if err == sql.ErrNoRows {
		return nil, repository.ErrRecordNotFound
	}
	return recs, err
}

func (r *TournamentTeamRepository) GetAll() ([]*models.TournamentTeam, error) {
	recs := []*models.TournamentTeam{}
	rows, err := r.DB.Query("select id, tournament_id, team_id, participation_type, is_winner from tournaments_teams")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		mod := &models.TournamentTeam{}
		err = rows.Scan(
			&mod.Id,
			&mod.TournamentId,
			&mod.TeamId,
			&mod.ParticipationType,
			&mod.IsWinner,
		)
		if err != nil {
			return nil, err
		}
		recs = append(recs, mod)
	}
	return recs, nil
}
