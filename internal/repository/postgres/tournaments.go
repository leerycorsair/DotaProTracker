package postgres

import (
	"database/sql"
	"local/internal/models"
	"local/internal/repository"
)

type TournamentRepository struct {
	DB *sql.DB
}

func (r *TournamentRepository) Add(mod models.Tournament) error {
	return r.DB.QueryRow("insert into tournaments (name, tier, prize_pool, date_start, duration, dpc_points, location) values ($1, $2, $3, $4, $5, $6, $7) returning id",
		mod.Name, mod.Tier, mod.PrizePool, mod.DateStart, mod.Duration, mod.DPCPoints, mod.Location).Scan(&mod.Id)
}

func (r *TournamentRepository) Edit(mod models.Tournament) error {
	_, err := r.DB.Exec("update tournaments set name=$1, tier=$2, prize_pool=$3, date_start=$4, duration=$5, dpc_points=$6, location=$7 where id=$8",
		mod.Name, mod.Tier, mod.PrizePool, mod.DateStart, mod.Duration, mod.DPCPoints, mod.Location, mod.Id)
	return err
}

func (r *TournamentRepository) Delete(id int) error {
	_, err := r.DB.Exec("delete from tournaments where id=$1", id)
	return err
}

func (r *TournamentRepository) FindById(id int) ([]*models.Tournament, error) {
	mod := models.Tournament{}
	err := r.DB.QueryRow("select id, name, tier, prize_pool, date_start, duration, dpc_points, location from tournaments where id=$1", id).Scan(
		&mod.Id,
		&mod.Name,
		&mod.Tier,
		&mod.PrizePool,
		&mod.DateStart,
		&mod.Duration,
		&mod.DPCPoints,
		&mod.Location,
	)
	recs := []*models.Tournament{}
	recs = append(recs, &mod)
	if err == sql.ErrNoRows {
		return nil, repository.ErrRecordNotFound
	}
	return recs, err
}

func (r *TournamentRepository) FindByName(name string) ([]*models.Tournament, error) {
	recs := []*models.Tournament{}
	rows, err := r.DB.Query("select id, name, tier, prize_pool, date_start, duration, dpc_points, location from tournaments where name like '%' || $1 || '%'", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		mod := &models.Tournament{}
		err = rows.Scan(
			&mod.Id,
			&mod.Name,
			&mod.Tier,
			&mod.PrizePool,
			&mod.DateStart,
			&mod.Duration,
			&mod.DPCPoints,
			&mod.Location,
		)
		if err != nil {
			return nil, err
		}
		recs = append(recs, mod)
	}
	return recs, nil
}

func (r *TournamentRepository) FindByYear(year int) ([]*models.Tournament, error) {
	recs := []*models.Tournament{}
	rows, err := r.DB.Query("select id, name, tier, prize_pool, date_start, duration, dpc_points, location from tournaments where date_part('year', date_start) = $1", year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		mod := &models.Tournament{}
		err = rows.Scan(
			&mod.Id,
			&mod.Name,
			&mod.Tier,
			&mod.PrizePool,
			&mod.DateStart,
			&mod.Duration,
			&mod.DPCPoints,
			&mod.Location,
		)
		if err != nil {
			return nil, err
		}
		recs = append(recs, mod)
	}
	return recs, nil
}

func (r *TournamentRepository) GetAll() ([]*models.Tournament, error) {
	recs := []*models.Tournament{}
	rows, err := r.DB.Query("select id, name, tier, prize_pool, date_start, duration, dpc_points, location from tournaments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		mod := &models.Tournament{}
		err = rows.Scan(
			&mod.Id,
			&mod.Name,
			&mod.Tier,
			&mod.PrizePool,
			&mod.DateStart,
			&mod.Duration,
			&mod.DPCPoints,
			&mod.Location,
		)
		if err != nil {
			return nil, err
		}
		recs = append(recs, mod)
	}
	return recs, nil
}
