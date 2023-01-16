package controllers

import (
	"local/internal/logger"
	"local/internal/models"
	"local/internal/repository"
)

type MatchPerfomanceController struct {
	Reps repository.Repositories
}

type MatchPerfomanceControllerBuilder struct{}

func (b MatchPerfomanceControllerBuilder) Build(Reps repository.Repositories) MatchPerfomanceController {
	return MatchPerfomanceController{Reps: Reps}
}

func (c MatchPerfomanceController) Add(mod models.MatchPerfomance) error {
	err := mod.Check()
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	err = c.Reps.MatchPerfomanceRep().Add(mod)
	return err
}

func (c MatchPerfomanceController) Edit(mod models.MatchPerfomance) error {
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
	return c.Reps.MatchPerfomanceRep().Edit(mod)
}

func (c MatchPerfomanceController) Delete(id int) error {
	_, err := c.FindById(id)
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	return c.Reps.MatchPerfomanceRep().Delete(id)
}

func (c MatchPerfomanceController) FindById(id int) ([]*models.MatchPerfomance, error) {
	return c.Reps.MatchPerfomanceRep().FindById(id)
}

func (c MatchPerfomanceController) GetAll() ([]*models.MatchPerfomance, error) {
	return c.Reps.MatchPerfomanceRep().GetAll()
}

func (c MatchPerfomanceController) FindByHero(hero string) ([]*models.MatchPerfomance, error) {
	return c.Reps.MatchPerfomanceRep().FindByHero(hero)
}
