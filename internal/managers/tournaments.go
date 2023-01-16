package managers

import (
	"errors"
	"local/internal/controllers"
	"local/internal/logger"
	"local/internal/models"
	"strconv"
)

type TournamentManager struct {
	C controllers.TournamentController
}

type TournamentManagerBuilder struct {
}

func (b TournamentManagerBuilder) Build(C controllers.TournamentController) TournamentManager {
	return TournamentManager{C: C}
}

func (man *TournamentManager) Add(vmod models.TournamentView) error {
	mod, err := vmod.FromViewConv()
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	err = man.C.Add(mod)
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	return nil
}

func (man *TournamentManager) Edit(vmod models.TournamentView) error {
	mod, err := vmod.FromViewConv()
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	err = man.C.Edit(mod)
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	return nil
}

func (man *TournamentManager) Delete(strid string) error {
	id, err := strconv.Atoi(strid)
	if err != nil {
		logger.Logger.Println(err)
		return errors.New("Некорректный ID")
	}
	err = man.C.Delete(id)
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	return nil
}

func (man *TournamentManager) GetAll() ([]*models.TournamentView, error) {
	vrecs := []*models.TournamentView{}
	recs, err := man.C.GetAll()
	if err != nil {
		logger.Logger.Println(err)
		return nil, err
	}
	for i := 0; i < len(recs); i++ {
		vrecs = append(vrecs, (recs[i].ToViewConv()))
	}
	return vrecs, nil
}

func (man *TournamentManager) FindById(strid string) ([]*models.TournamentView, error) {
	id, err := strconv.Atoi(strid)
	if err != nil {
		logger.Logger.Println(err)
		return nil, errors.New("Некорректный ID")
	}
	vrecs := []*models.TournamentView{}
	recs, err := man.C.FindById(id)
	if err != nil {
		logger.Logger.Println(err)
		return nil, err
	}
	for i := 0; i < len(recs); i++ {
		vrecs = append(vrecs, (recs[i].ToViewConv()))
	}
	return vrecs, nil
}

func (man *TournamentManager) FindByYear(stryear string) ([]*models.TournamentView, error) {
	year, err := strconv.Atoi(stryear)
	if err != nil {
		logger.Logger.Println(err)
		return nil, errors.New("Некорректный год")
	}
	vrecs := []*models.TournamentView{}
	recs, err := man.C.FindByYear(year)
	if err != nil {
		logger.Logger.Println(err)
		return nil, err
	}
	for i := 0; i < len(recs); i++ {
		vrecs = append(vrecs, (recs[i].ToViewConv()))
	}
	return vrecs, nil
}

func (man *TournamentManager) FindByName(name string) ([]*models.TournamentView, error) {
	if len(name) == 0 {
		return nil, errors.New("Некорректное название")
	}
	vrecs := []*models.TournamentView{}
	recs, err := man.C.FindByName(name)
	if err != nil {
		logger.Logger.Println(err)
		return nil, err
	}
	for i := 0; i < len(recs); i++ {
		vrecs = append(vrecs, (recs[i].ToViewConv()))
	}
	return vrecs, nil
}
