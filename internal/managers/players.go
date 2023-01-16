package managers

import (
	"errors"
	"local/internal/controllers"
	"local/internal/logger"
	"local/internal/models"
	"strconv"
)

type PlayerManager struct {
	C controllers.PlayerController
}

type PlayerManagerBuilder struct {
}

func (b PlayerManagerBuilder) Build(C controllers.PlayerController) PlayerManager {
	return PlayerManager{C: C}
}

func (man *PlayerManager) Add(vmod models.PlayerView) error {
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

func (man *PlayerManager) Edit(vmod models.PlayerView) error {
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

func (man *PlayerManager) Delete(strid string) error {
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

func (man *PlayerManager) GetAll() ([]*models.PlayerView, error) {
	vrecs := []*models.PlayerView{}
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

func (man *PlayerManager) FindById(strid string) ([]*models.PlayerView, error) {
	id, err := strconv.Atoi(strid)
	if err != nil {
		logger.Logger.Println(err)
		return nil, errors.New("Некорректный ID")
	}
	vrecs := []*models.PlayerView{}
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

func (man *PlayerManager) FindByBirthdate(stryear string) ([]*models.PlayerView, error) {
	year, err := strconv.Atoi(stryear)
	if err != nil {
		logger.Logger.Println(err)
		return nil, errors.New("Некорректный год рождения")
	}
	vrecs := []*models.PlayerView{}
	recs, err := man.C.FindByBirthdate(year)
	if err != nil {
		logger.Logger.Println(err)
		return nil, err
	}
	for i := 0; i < len(recs); i++ {
		vrecs = append(vrecs, (recs[i].ToViewConv()))
	}
	return vrecs, nil
}

func (man *PlayerManager) FindByName(name string) ([]*models.PlayerView, error) {
	if len(name) == 0 {
		return nil, errors.New("Некорректное имя")
	}
	vrecs := []*models.PlayerView{}
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
