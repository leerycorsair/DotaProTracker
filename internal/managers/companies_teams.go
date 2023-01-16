package managers

import (
	"errors"
	"local/internal/controllers"
	"local/internal/logger"
	"local/internal/models"
	"strconv"
)

type CompanyTeamManager struct {
	C controllers.CompanyTeamController
}

type CompanyTeamManagerBuilder struct {
}

func (b CompanyTeamManagerBuilder) Build(C controllers.CompanyTeamController) CompanyTeamManager {
	return CompanyTeamManager{C: C}
}

func (man *CompanyTeamManager) Add(vmod models.CompanyTeamView) error {
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

func (man *CompanyTeamManager) Edit(vmod models.CompanyTeamView) error {
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

func (man *CompanyTeamManager) Delete(strid string) error {
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

func (man *CompanyTeamManager) GetAll() ([]*models.CompanyTeamView, error) {
	vrecs := []*models.CompanyTeamView{}
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

func (man *CompanyTeamManager) FindById(strid string) ([]*models.CompanyTeamView, error) {
	id, err := strconv.Atoi(strid)
	if err != nil {
		logger.Logger.Println(err)
		return nil, errors.New("Некорректный ID")
	}
	vrecs := []*models.CompanyTeamView{}
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
