package postgres

import (
	"database/sql"
	"local/internal/models"
	"local/internal/repository"
)

type MatchRepository struct {
	DB *sql.DB
}

func (r *MatchRepository) Add(mod models.Match) error {
	return r.DB.QueryRow("insert into matches  (tournament_id, r_team_id, d_team_id, duration, winner) values ($1, $2, $3, $4, $5) returning id",
		mod.TournamentId, mod.RTeamId, mod.DTeamId, mod.Duration, mod.Winner).Scan(&mod.Id)
}

func (r *MatchRepository) Edit(mod models.Match) error {
	_, err := r.DB.Exec("update matches set tournament_id=$1, r_team_id=$2, d_team_id=$3, duration=$4, winner=$5 where id=$6",
		mod.TournamentId, mod.RTeamId, mod.DTeamId, mod.Duration, mod.Winner, mod.Id)
	return err
}

func (r *MatchRepository) Delete(id int) error {
	_, err := r.DB.Exec("delete from matches where id=$1", id)
	return err
}

func (r *MatchRepository) FindById(id int) ([]*models.Match, error) {
	mod := models.Match{}
	err := r.DB.QueryRow("select id, tournament_id, r_team_id, d_team_id, duration, winner from matches where id=$1", id).Scan(
		&mod.Id,
		&mod.TournamentId,
		&mod.RTeamId,
		&mod.DTeamId,
		&mod.Duration,
		&mod.Winner,
	)
	recs := []*models.Match{}
	recs = append(recs, &mod)
	if err == sql.ErrNoRows {
		return nil, repository.ErrRecordNotFound
	}
	return recs, err
}

func (r *MatchRepository) GetAll() ([]*models.Match, error) {
	recs := []*models.Match{}
	rows, err := r.DB.Query("select id, tournament_id, r_team_id, d_team_id, duration, winner from matches")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		mod := &models.Match{}
		err = rows.Scan(
			&mod.Id,
			&mod.TournamentId,
			&mod.RTeamId,
			&mod.DTeamId,
			&mod.Duration,
			&mod.Winner,
		)
		if err != nil {
			return nil, err
		}
		recs = append(recs, mod)
	}
	return recs, nil
}
