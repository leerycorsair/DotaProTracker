package postgres

import (
	"database/sql"
	"local/internal/logger"
	"local/internal/models"
	"local/internal/repository"
)

type TeamRepository struct {
	DB *sql.DB
}

func (r *TeamRepository) Add(mod models.Team) error {
	return r.DB.QueryRow("insert into teams (name, created_at, email, total_earnings, region, tier) values ($1, $2, $3, $4, $5, $6, $7) returning id",
		mod.Name, mod.CreatedAt, mod.Email, mod.TotalEarnings, mod.Region, mod.Tier).Scan(&mod.Id)
}

func (r *TeamRepository) Edit(mod models.Team) error {
	_, err := r.DB.Exec("update teams set name=$1, created_at=$2, email=$3, total_earnings=$4, region=$5, tier=$6 where id=$8",
		mod.Name, mod.CreatedAt, mod.Email, mod.TotalEarnings, mod.Region, mod.Tier, mod.Id)
	return err
}

func (r *TeamRepository) Delete(id int) error {
	_, err := r.DB.Exec("delete from teams where id=$1", id)
	return err
}

func (r *TeamRepository) FindById(id int) ([]*models.Team, error) {
	mod := models.Team{}
	err := r.DB.QueryRow("select id, name, created_at, email, total_earnings, region, tier from teams where id=$1", id).Scan(
		&mod.Id,
		&mod.Name,
		&mod.CreatedAt,
		&mod.Email,
		&mod.TotalEarnings,
		&mod.Region,
		&mod.Tier,
	)
	recs := []*models.Team{}
	recs = append(recs, &mod)
	if err == sql.ErrNoRows {
		return nil, repository.ErrRecordNotFound
	}
	return recs, err
}

func (r *TeamRepository) FindByName(name string) ([]*models.Team, error) {
	recs := []*models.Team{}
	rows, err := r.DB.Query("select id, name, created_at, email, total_earnings, region, tier from teams where name like '%' || $1 || '%'", name)
	if err != nil {
		logger.Logger.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		mod := &models.Team{}
		err = rows.Scan(
			&mod.Id,
			&mod.Name,
			&mod.CreatedAt,
			&mod.Email,
			&mod.TotalEarnings,
			&mod.Region,
			&mod.Tier,
		)
		if err != nil {
			logger.Logger.Println(err)
			return nil, err
		}
		recs = append(recs, mod)
	}
	return recs, nil
}

func (r *TeamRepository) GetAll() ([]*models.Team, error) {
	recs := []*models.Team{}
	rows, err := r.DB.Query("select id, name, created_at, email, total_earnings, region, tier from teams")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		mod := &models.Team{}
		err = rows.Scan(
			&mod.Id,
			&mod.Name,
			&mod.CreatedAt,
			&mod.Email,
			&mod.TotalEarnings,
			&mod.Region,
			&mod.Tier,
		)
		if err != nil {
			return nil, err
		}
		recs = append(recs, mod)
	}
	return recs, nil
}
