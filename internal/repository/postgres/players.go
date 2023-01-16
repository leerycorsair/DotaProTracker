package postgres

import (
	"database/sql"
	"local/internal/models"
	"local/internal/repository"
)

type PlayerRepository struct {
	DB *sql.DB
}

func (r *PlayerRepository) Add(mod models.Player) error {
	return r.DB.QueryRow("insert into players (nickname, realname, birthdate, country, mmr, role, signature_hero) values ($1, $2, $3, $4, $5, $6, $7) returning id",
		mod.Nickname, mod.Realname, mod.Birthdate, mod.Country, mod.MMR, mod.Role, mod.SignatureHero).Scan(&mod.Id)
}

func (r *PlayerRepository) Edit(mod models.Player) error {
	_, err := r.DB.Exec("update players set nickname=$1, realname=$2, birthdate=$3, country=$4, mmr=$5, role=$6, signature_hero=$7 where id=$8",
		mod.Nickname, mod.Realname, mod.Birthdate, mod.Country, mod.MMR, mod.Role, mod.SignatureHero, mod.Id)
	return err
}

func (r *PlayerRepository) Delete(id int) error {
	_, err := r.DB.Exec("delete from players where id=$1", id)
	return err
}

func (r *PlayerRepository) FindById(id int) ([]*models.Player, error) {
	mod := models.Player{}
	err := r.DB.QueryRow("select id, nickname, realname, birthdate, country, mmr, role, signature_hero from players where id=$1", id).Scan(
		&mod.Id,
		&mod.Nickname,
		&mod.Realname,
		&mod.Birthdate,
		&mod.Country,
		&mod.MMR,
		&mod.Role,
		&mod.SignatureHero,
	)
	recs := []*models.Player{}
	recs = append(recs, &mod)
	if err == sql.ErrNoRows {
		return nil, repository.ErrRecordNotFound
	}
	return recs, err
}

func (r *PlayerRepository) FindByName(name string) ([]*models.Player, error) {
	recs := []*models.Player{}
	rows, err := r.DB.Query("select id, nickname, realname, birthdate, country, mmr, role, signature_hero from players where realname like '%' || $1 || '%'", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		mod := &models.Player{}
		err = rows.Scan(
			&mod.Id,
			&mod.Nickname,
			&mod.Realname,
			&mod.Birthdate,
			&mod.Country,
			&mod.MMR,
			&mod.Role,
			&mod.SignatureHero,
		)
		if err != nil {
			return nil, err
		}
		recs = append(recs, mod)
	}
	return recs, nil
}

func (r *PlayerRepository) FindByBirthdate(year int) ([]*models.Player, error) {
	recs := []*models.Player{}
	rows, err := r.DB.Query("select id, nickname, realname, birthdate, country, mmr, role, signature_hero from players where date_part('year', birthdate) = $1", year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		mod := &models.Player{}
		err = rows.Scan(
			&mod.Id,
			&mod.Nickname,
			&mod.Realname,
			&mod.Birthdate,
			&mod.Country,
			&mod.MMR,
			&mod.Role,
			&mod.SignatureHero,
		)
		if err != nil {
			return nil, err
		}
		recs = append(recs, mod)
	}
	return recs, nil
}

func (r *PlayerRepository) GetAll() ([]*models.Player, error) {
	recs := []*models.Player{}
	rows, err := r.DB.Query("select id, nickname, realname, birthdate, country, mmr, role, signature_hero from players")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		mod := &models.Player{}
		err = rows.Scan(
			&mod.Id,
			&mod.Nickname,
			&mod.Realname,
			&mod.Birthdate,
			&mod.Country,
			&mod.MMR,
			&mod.Role,
			&mod.SignatureHero,
		)
		if err != nil {
			return nil, err
		}
		recs = append(recs, mod)
	}
	return recs, nil
}
