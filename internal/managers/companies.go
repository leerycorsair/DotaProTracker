package managers

import (
	"errors"
	"local/internal/controllers"
	"local/internal/logger"
	"local/internal/models"
	"strconv"
)

type CompanyManager struct {
	C controllers.CompanyController
}

type CompanyManagerBuilder struct {
}

func (b CompanyManagerBuilder) Build(C controllers.CompanyController) CompanyManager {
	return CompanyManager{C: C}
}

func (man *CompanyManager) Add(vmod models.CompanyView) error {
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

func (man *CompanyManager) Edit(vmod models.CompanyView) error {
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

func (man *CompanyManager) Delete(strid string) error {
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

func (man *CompanyManager) GetAll() ([]*models.CompanyView, error) {
	vrecs := []*models.CompanyView{}
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

func (man *CompanyManager) FindById(strid string) ([]*models.CompanyView, error) {
	id, err := strconv.Atoi(strid)
	if err != nil {
		logger.Logger.Println(err)
		return nil, errors.New("Некорректный ID")
	}
	vrecs := []*models.CompanyView{}
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

func (man *CompanyManager) FindByName(name string) ([]*models.CompanyView, error) {
	if len(name) == 0 {
		return nil, errors.New("Некорректное название")
	}
	vrecs := []*models.CompanyView{}
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
