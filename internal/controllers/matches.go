package controllers

import (
	"local/internal/logger"
	"local/internal/models"
	"local/internal/repository"
)

type MatchController struct {
	Reps repository.Repositories
}

type MatchControllerBuilder struct{}

func (b MatchControllerBuilder) Build(Reps repository.Repositories) MatchController {
	return MatchController{Reps: Reps}
}

func (c MatchController) Add(mod models.Match) error {
	err := mod.Check()
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	err = c.Reps.MatchRep().Add(mod)
	return err
}

func (c MatchController) Edit(mod models.Match) error {
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
	return c.Reps.MatchRep().Edit(mod)
}

func (c MatchController) Delete(id int) error {
	_, err := c.FindById(id)
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	return c.Reps.MatchRep().Delete(id)
}

func (c MatchController) FindById(id int) ([]*models.Match, error) {
	return c.Reps.MatchRep().FindById(id)
}

func (c MatchController) GetAll() ([]*models.Match, error) {
	return c.Reps.MatchRep().GetAll()
}
