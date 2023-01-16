package postgres

import (
	"database/sql"
	"local/internal/models"
	"local/internal/repository"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) Add(mod models.User) error {
	return r.DB.QueryRow("insert into users (name, birthdate, login, password, email, privilege_level) values ($1, $2, $3, $4, $5, $6) returning id",
		mod.Name, mod.Birthdate, mod.Login, mod.Password, mod.Email, mod.PrivilegeLevel).Scan(&mod.Id)
}

func (r *UserRepository) Edit(mod models.User) error {
	_, err := r.DB.Exec("update users set name=$1, birthdate=$2, login=$3, password=$4, email=$5, privilege_level=$6 where id=$7",
		mod.Name, mod.Birthdate, mod.Login, mod.Password, mod.Email, mod.PrivilegeLevel, mod.Id)
	return err
}

func (r *UserRepository) Delete(id int) error {
	_, err := r.DB.Exec("delete from users where id=$1", id)
	return err
}

func (r *UserRepository) FindById(id int) ([]*models.User, error) {
	mod := models.User{}
	err := r.DB.QueryRow("select id, name, birthdate, login, password, email, privilege_level from users where id=$1", id).Scan(
		&mod.Id,
		&mod.Name,
		&mod.Birthdate,
		&mod.Login,
		&mod.Password,
		&mod.Email,
		&mod.PrivilegeLevel,
	)
	recs := []*models.User{}
	recs = append(recs, &mod)
	if err == sql.ErrNoRows {
		return nil, repository.ErrRecordNotFound
	}
	return recs, err
}

func (r *UserRepository) FindByLogin(login string) ([]*models.User, error) {
	mod := models.User{}
	err := r.DB.QueryRow("select id, name, birthdate, login, password, email, privilege_level from users where login=$1", login).Scan(
		&mod.Id,
		&mod.Name,
		&mod.Birthdate,
		&mod.Login,
		&mod.Password,
		&mod.Email,
		&mod.PrivilegeLevel,
	)
	recs := []*models.User{}
	recs = append(recs, &mod)
	if err == sql.ErrNoRows {
		return nil, repository.ErrRecordNotFound
	}
	return recs, err
}

func (r *UserRepository) GetAll() ([]*models.User, error) {
	recs := []*models.User{}
	rows, err := r.DB.Query("select id, name, birthdate, login, password, email, privilege_level from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		mod := &models.User{}
		err = rows.Scan(
			&mod.Id,
			&mod.Name,
			&mod.Birthdate,
			&mod.Login,
			&mod.Password,
			&mod.Email,
			&mod.PrivilegeLevel,
		)
		if err != nil {
			return nil, err
		}
		recs = append(recs, mod)
	}
	return recs, nil
}
