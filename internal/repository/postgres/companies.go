package postgres

import (
	"database/sql"
	"local/internal/models"
	"local/internal/repository"
)

type CompanyRepository struct {
	DB *sql.DB
}

func (r *CompanyRepository) Add(mod models.Company) error {
	return r.DB.QueryRow("insert into companies  (name, country, website, revenue, industry) values ($1, $2, $3, $4, $5) returning id",
		mod.Name, mod.Country, mod.Website, mod.Revenue, mod.Industry).Scan(&mod.Id)
}

func (r *CompanyRepository) Edit(mod models.Company) error {
	_, err := r.DB.Exec("update companies set name=$1, country=$2, website=$3, revenue=$4, industry=$5 where id=$6",
		mod.Name, mod.Country, mod.Website, mod.Revenue, mod.Industry, mod.Id)
	return err
}

func (r *CompanyRepository) Delete(id int) error {
	_, err := r.DB.Exec("delete from companies where id=$1", id)
	return err
}

func (r *CompanyRepository) FindById(id int) ([]*models.Company, error) {
	mod := models.Company{}
	err := r.DB.QueryRow("select id, name, country, website, revenue, industry from companies where id=$1", id).Scan(
		&mod.Id,
		&mod.Name,
		&mod.Country,
		&mod.Website,
		&mod.Revenue,
		&mod.Industry,
	)
	recs := []*models.Company{}
	recs = append(recs, &mod)
	if err == sql.ErrNoRows {
		return nil, repository.ErrRecordNotFound
	}
	return recs, err
}

func (r *CompanyRepository) FindByName(name string) ([]*models.Company, error) {
	recs := []*models.Company{}
	rows, err := r.DB.Query("select id, name, country, website, revenue, industry from companies where name like '%' || $1 || '%'", name)
	if err != nil {

		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		mod := &models.Company{}
		err = rows.Scan(
			&mod.Id,
			&mod.Name,
			&mod.Country,
			&mod.Website,
			&mod.Revenue,
			&mod.Industry,
		)
		if err != nil {

			return nil, err
		}
		recs = append(recs, mod)
	}
	return recs, nil
}

func (r *CompanyRepository) GetAll() ([]*models.Company, error) {
	recs := []*models.Company{}
	rows, err := r.DB.Query("select id, name, country, website, revenue, industry from companies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		mod := &models.Company{}
		err = rows.Scan(
			&mod.Id,
			&mod.Name,
			&mod.Country,
			&mod.Website,
			&mod.Revenue,
			&mod.Industry,
		)
		if err != nil {
			return nil, err
		}
		recs = append(recs, mod)
	}
	return recs, nil
}
