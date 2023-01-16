package postgres

import (
	"database/sql"
	"local/internal/models"
	"local/internal/repository"
)

type CompanyTeamRepository struct {
	DB *sql.DB
}

func (r *CompanyTeamRepository) Add(mod models.CompanyTeam) error {
	return r.DB.QueryRow("insert into companies_teams (company_id, team_id, contract_date, contract_time) values ($1, $2, $3, $4) returning id",
		mod.CompanyId, mod.TeamId, mod.ContractDate, mod.ContractTime).Scan(&mod.Id)
}

func (r *CompanyTeamRepository) Edit(mod models.CompanyTeam) error {
	_, err := r.DB.Exec("update companies_teams set company_id=$1, team_id=$2, contract_date=$3, contract_time=$4 where id=$5",
		mod.CompanyId, mod.TeamId, mod.ContractDate, mod.ContractTime, mod.Id)
	return err
}

func (r *CompanyTeamRepository) Delete(id int) error {
	_, err := r.DB.Exec("delete from companies_teams where id=$1", id)
	return err
}

func (r *CompanyTeamRepository) FindById(id int) ([]*models.CompanyTeam, error) {
	mod := models.CompanyTeam{}
	err := r.DB.QueryRow("select id, company_id, team_id, contract_date, contract_time from companies_teams where id=$1", id).Scan(
		&mod.Id,
		&mod.CompanyId,
		&mod.TeamId,
		&mod.ContractDate,
		&mod.ContractTime,
	)
	recs := []*models.CompanyTeam{}
	recs = append(recs, &mod)
	if err == sql.ErrNoRows {
		return nil, repository.ErrRecordNotFound
	}
	return recs, err
}

func (r *CompanyTeamRepository) GetAll() ([]*models.CompanyTeam, error) {
	recs := []*models.CompanyTeam{}
	rows, err := r.DB.Query("select id, company_id, team_id, contract_date, contract_time from companies_teams")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		mod := &models.CompanyTeam{}
		err = rows.Scan(
			&mod.Id,
			&mod.CompanyId,
			&mod.TeamId,
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
