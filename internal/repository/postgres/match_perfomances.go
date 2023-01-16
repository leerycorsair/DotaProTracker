package postgres

import (
	"database/sql"
	"local/internal/logger"
	"local/internal/models"
	"local/internal/repository"
)

type MatchPerfomanceRepository struct {
	DB *sql.DB
}

func (r *MatchPerfomanceRepository) Add(mod models.MatchPerfomance) error {
	return r.DB.QueryRow("insert into match_perfomances  (match_id, player_id, team, hero, kills, deaths, assists, networth, gpm, xpm, dmg, heal, bld) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, 13) returning id",
		mod.MatchId, mod.PlayerId, mod.Team, mod.Hero, mod.Kills, mod.Deaths, mod.Assists, mod.Networth, mod.GPM, mod.XPM, mod.DMG, mod.Heal, mod.BLD).Scan(&mod.Id)
}

func (r *MatchPerfomanceRepository) Edit(mod models.MatchPerfomance) error {
	_, err := r.DB.Exec("update match_perfomances set match_id=$1, player_id=$2, team=$3, hero=$4, kills=$5, deaths=$6, assists=$7, networth=$8, gpm=$9, xpm=$10, dmg=$11, heal=$12, bld=$13 where id=$14",
		mod.MatchId, mod.PlayerId, mod.Team, mod.Hero, mod.Kills, mod.Deaths, mod.Assists, mod.Networth, mod.GPM, mod.XPM, mod.DMG, mod.Heal, mod.BLD, mod.Id)
	return err
}

func (r *MatchPerfomanceRepository) Delete(id int) error {
	_, err := r.DB.Exec("delete from match_perfomances where id=$1", id)
	return err
}

func (r *MatchPerfomanceRepository) FindById(id int) ([]*models.MatchPerfomance, error) {
	mod := models.MatchPerfomance{}
	err := r.DB.QueryRow("select id, match_id, player_id, team, hero, kills, deaths, assists, networth, gpm, xpm, dmg, heal, bld from match_perfomances where id=$1", id).Scan(
		&mod.Id,
		&mod.MatchId,
		&mod.PlayerId,
		&mod.Team,
		&mod.Hero,
		&mod.Kills,
		&mod.Deaths,
		&mod.Assists,
		&mod.Networth,
		&mod.GPM,
		&mod.XPM,
		&mod.DMG,
		&mod.Heal,
		&mod.BLD,
	)
	recs := []*models.MatchPerfomance{}
	recs = append(recs, &mod)
	if err == sql.ErrNoRows {
		return nil, repository.ErrRecordNotFound
	}
	return recs, err
}

func (r *MatchPerfomanceRepository) FindByHero(hero string) ([]*models.MatchPerfomance, error) {
	recs := []*models.MatchPerfomance{}
	rows, err := r.DB.Query("select id, match_id, player_id, team, hero, kills, deaths, assists, networth, gpm, xpm, dmg, heal, bld from match_perfomances where hero like '%' || $1 || '%'", hero)
	if err != nil {
		logger.Logger.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		mod := &models.MatchPerfomance{}
		err = rows.Scan(
			&mod.Id,
			&mod.MatchId,
			&mod.PlayerId,
			&mod.Team,
			&mod.Hero,
			&mod.Kills,
			&mod.Deaths,
			&mod.Assists,
			&mod.Networth,
			&mod.GPM,
			&mod.XPM,
			&mod.DMG,
			&mod.Heal,
			&mod.BLD,
		)
		if err != nil {
			logger.Logger.Println(err)
			return nil, err
		}
		recs = append(recs, mod)
	}
	return recs, nil
}

func (r *MatchPerfomanceRepository) GetAll() ([]*models.MatchPerfomance, error) {
	recs := []*models.MatchPerfomance{}
	rows, err := r.DB.Query("select id, match_id, player_id, team, hero, kills, deaths, assists, networth, gpm, xpm, dmg, heal, bld from match_perfomances")
	if err != nil {
		logger.Logger.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		mod := &models.MatchPerfomance{}
		err = rows.Scan(
			&mod.Id,
			&mod.MatchId,
			&mod.PlayerId,
			&mod.Team,
			&mod.Hero,
			&mod.Kills,
			&mod.Deaths,
			&mod.Assists,
			&mod.Networth,
			&mod.GPM,
			&mod.XPM,
			&mod.DMG,
			&mod.Heal,
			&mod.BLD,
		)
		if err != nil {
			logger.Logger.Println(err)
			return nil, err
		}
		recs = append(recs, mod)
	}
	return recs, nil
}
