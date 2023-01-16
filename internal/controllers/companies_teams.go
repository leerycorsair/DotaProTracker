package controllers

import (
	"local/internal/logger"
	"local/internal/models"
	"local/internal/repository"
)

type CompanyTeamController struct {
	Reps repository.Repositories
}

type CompanyTeamControllerBuilder struct{}

func (b CompanyTeamControllerBuilder) Build(Reps repository.Repositories) CompanyTeamController {
	return CompanyTeamController{Reps: Reps}
}

func (c CompanyTeamController) Add(mod models.CompanyTeam) error {
	err := mod.Check()
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	err = c.Reps.CompanyTeamRep().Add(mod)
	return err
}

func (c CompanyTeamController) Edit(mod models.CompanyTeam) error {
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
	return c.Reps.CompanyTeamRep().Edit(mod)
}

func (c CompanyTeamController) Delete(id int) error {
	_, err := c.FindById(id)
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	return c.Reps.CompanyTeamRep().Delete(id)
}

func (c CompanyTeamController) FindById(id int) ([]*models.CompanyTeam, error) {
	return c.Reps.CompanyTeamRep().FindById(id)
}

func (c CompanyTeamController) GetAll() ([]*models.CompanyTeam, error) {
	return c.Reps.CompanyTeamRep().GetAll()
}
