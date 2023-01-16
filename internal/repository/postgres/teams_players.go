package postgres

import (
	"database/sql"
	"local/internal/models"
	"local/internal/repository"
)

type TeamPlayerRepository struct {
	DB *sql.DB
}

func (r *TeamPlayerRepository) Add(mod models.TeamPlayer) error {
	return r.DB.QueryRow("insert into teams_players (team_id, player_id, contract_date, contract_time) values ($1, $2, $3, $4) returning id",
		mod.TeamId, mod.PlayerId, mod.ContractDate, mod.ContractTime).Scan(&mod.Id)
}

func (r *TeamPlayerRepository) Edit(mod models.TeamPlayer) error {
	_, err := r.DB.Exec("update teams_players set team_id=$1, player_id=$2, contract_date=$3, contract_time=$4 where id=$5",
		mod.TeamId, mod.PlayerId, mod.ContractDate, mod.ContractTime, mod.Id)
	return err
}

func (r *TeamPlayerRepository) Delete(id int) error {
	_, err := r.DB.Exec("delete from teams_players where id=$1", id)
	return err
}

func (r *TeamPlayerRepository) FindById(id int) ([]*models.TeamPlayer, error) {
	mod := models.TeamPlayer{}
	err := r.DB.QueryRow("select id, team_id, player_id, contract_date, contract_time from teams_players where id=$1", id).Scan(
		&mod.Id,
		&mod.TeamId,
		&mod.PlayerId,
		&mod.ContractDate,
		&mod.ContractTime,
	)
	recs := []*models.TeamPlayer{}
	recs = append(recs, &mod)
	if err == sql.ErrNoRows {
		return nil, repository.ErrRecordNotFound
	}
	return recs, err
}

func (r *TeamPlayerRepository) GetAll() ([]*models.TeamPlayer, error) {
	recs := []*models.TeamPlayer{}
	rows, err := r.DB.Query("select id, team_id, player_id, contract_date, contract_time from teams_players")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		mod := &models.TeamPlayer{}
		err = rows.Scan(
			&mod.Id,
			&mod.TeamId,
			&mod.PlayerId,
			&mod.ContractDate,
			&mod.ContractTime,
		)
		if err != nil {
			return nil, err
		}
		recs = append(recs, mod)
	}
	return recs, nil
}
