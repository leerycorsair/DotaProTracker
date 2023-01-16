package managers

import (
	"errors"
	"local/internal/controllers"
	"local/internal/logger"
	"local/internal/models"
	"strconv"
)

type MatchPerfomanceManager struct {
	C controllers.MatchPerfomanceController
}

type MatchPerfomanceManagerBuilder struct {
}

func (b MatchPerfomanceManagerBuilder) Build(C controllers.MatchPerfomanceController) MatchPerfomanceManager {
	return MatchPerfomanceManager{C: C}
}

func (man *MatchPerfomanceManager) Add(vmod models.MatchPerfomanceView) error {
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

func (man *MatchPerfomanceManager) Edit(vmod models.MatchPerfomanceView) error {
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

func (man *MatchPerfomanceManager) Delete(strid string) error {
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

func (man *MatchPerfomanceManager) GetAll() ([]*models.MatchPerfomanceView, error) {
	vrecs := []*models.MatchPerfomanceView{}
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

func (man *MatchPerfomanceManager) FindById(strid string) ([]*models.MatchPerfomanceView, error) {
	id, err := strconv.Atoi(strid)
	if err != nil {
		logger.Logger.Println(err)
		return nil, errors.New("Некорректный ID")
	}
	vrecs := []*models.MatchPerfomanceView{}
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

func (man *MatchPerfomanceManager) FindByHero(hero string) ([]*models.MatchPerfomanceView, error) {
	if len(hero) == 0 {
		return nil, errors.New("Некорректное имя героя")
	}
	vrecs := []*models.MatchPerfomanceView{}
	recs, err := man.C.FindByHero(hero)
	if err != nil {
		logger.Logger.Println(err)
		return nil, err
	}
	for i := 0; i < len(recs); i++ {
		vrecs = append(vrecs, (recs[i].ToViewConv()))
	}
	return vrecs, nil
}
