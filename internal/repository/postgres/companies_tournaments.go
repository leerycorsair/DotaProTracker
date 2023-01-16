package postgres

import (
	"database/sql"
	"local/internal/models"
	"local/internal/repository"
)

type CompanyTournamentRepository struct {
	DB *sql.DB
}

func (r *CompanyTournamentRepository) Add(mod models.CompanyTournament) error {
	return r.DB.QueryRow("insert into companies_tournaments  (company_id, tournament_id, deposit) values ($1, $2, $3) returning id",
		mod.CompanyId, mod.TournamentId, mod.Deposit).Scan(&mod.Id)
}

func (r *CompanyTournamentRepository) Edit(mod models.CompanyTournament) error {
	_, err := r.DB.Exec("update companies_tournaments set company_id=$1, tournament_id=$2, deposit=$3 where id=$4",
		mod.CompanyId, mod.TournamentId, mod.Deposit, mod.Id)
	return err
}

func (r *CompanyTournamentRepository) Delete(id int) error {
	_, err := r.DB.Exec("delete from companies_tournaments where id=$1", id)
	return err
}

func (r *CompanyTournamentRepository) FindById(id int) ([]*models.CompanyTournament, error) {
	mod := models.CompanyTournament{}
	err := r.DB.QueryRow("select id, company_id, tournament_id, deposit from companies_tournaments where id=$1", id).Scan(
		&mod.Id,
		&mod.CompanyId,
		&mod.TournamentId,
		&mod.Deposit,
	)
	recs := []*models.CompanyTournament{}
	recs = append(recs, &mod)
	if err == sql.ErrNoRows {
		return nil, repository.ErrRecordNotFound
	}
	return recs, err
}

func (r *CompanyTournamentRepository) GetAll() ([]*models.CompanyTournament, error) {
	recs := []*models.CompanyTournament{}
	rows, err := r.DB.Query("select id, company_id, tournament_id, deposit from companies_tournaments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		mod := &models.CompanyTournament{}
		err = rows.Scan(
			&mod.Id,
			&mod.CompanyId,
			&mod.TournamentId,
			&mod.Deposit,
		)
		if err != nil {
			return nil, err
		}
		recs = append(recs, mod)
	}
	return recs, nil
}
