package controllers

import (
	"local/internal/logger"
	"local/internal/models"
	"local/internal/repository"
)

type TeamController struct {
	Reps repository.Repositories
}

type TeamControllerBuilder struct{}

func (b TeamControllerBuilder) Build(Reps repository.Repositories) TeamController {
	return TeamController{Reps: Reps}
}

func (c TeamController) Add(mod models.Team) error {
	err := mod.Check()
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	err = c.Reps.TeamRep().Add(mod)
	return err
}

func (c TeamController) Edit(mod models.Team) error {
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
	return c.Reps.TeamRep().Edit(mod)
}

func (c TeamController) Delete(id int) error {
	_, err := c.FindById(id)
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	return c.Reps.TeamRep().Delete(id)
}

func (c TeamController) FindById(id int) ([]*models.Team, error) {
	return c.Reps.TeamRep().FindById(id)
}

func (c TeamController) GetAll() ([]*models.Team, error) {
	return c.Reps.TeamRep().GetAll()
}

func (c TeamController) FindByName(name string) ([]*models.Team, error) {
	return c.Reps.TeamRep().FindByName(name)
}
