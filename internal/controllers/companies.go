package controllers

import (
	"local/internal/logger"
	"local/internal/models"
	"local/internal/repository"
)

type CompanyController struct {
	Reps repository.Repositories
}

type CompanyControllerBuilder struct{}

func (b CompanyControllerBuilder) Build(Reps repository.Repositories) CompanyController {
	return CompanyController{Reps: Reps}
}

func (c CompanyController) Add(mod models.Company) error {
	err := mod.Check()
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	err = c.Reps.CompanyRep().Add(mod)
	return err
}

func (c CompanyController) Edit(mod models.Company) error {
	err := mod.Check()
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	_, err = c.FindById(mod.Id)
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	return c.Reps.CompanyRep().Edit(mod)
}

func (c CompanyController) Delete(id int) error {
	_, err := c.FindById(id)
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	return c.Reps.CompanyRep().Delete(id)
}

func (c CompanyController) FindById(id int) ([]*models.Company, error) {
	return c.Reps.CompanyRep().FindById(id)
}

func (c CompanyController) GetAll() ([]*models.Company, error) {
	return c.Reps.CompanyRep().GetAll()
}

func (c CompanyController) FindByName(name string) ([]*models.Company, error) {
	return c.Reps.CompanyRep().FindByName(name)
}
