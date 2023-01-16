package controllers

import (
	"local/internal/logger"
	"local/internal/models"
	"local/internal/repository"
)

type TournamentController struct {
	Reps repository.Repositories
}

type TournamentControllerBuilder struct{}

func (b TournamentControllerBuilder) Build(Reps repository.Repositories) TournamentController {
	return TournamentController{Reps: Reps}
}

func (c TournamentController) Add(mod models.Tournament) error {
	err := mod.Check()
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	err = c.Reps.TournamentRep().Add(mod)
	return err
}

func (c TournamentController) Edit(mod models.Tournament) error {
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
	return c.Reps.TournamentRep().Edit(mod)
}

func (c TournamentController) Delete(id int) error {
	_, err := c.FindById(id)
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	return c.Reps.TournamentRep().Delete(id)
}

func (c TournamentController) FindById(id int) ([]*models.Tournament, error) {
	return c.Reps.TournamentRep().FindById(id)
}

func (c TournamentController) FindByYear(year int) ([]*models.Tournament, error) {
	return c.Reps.TournamentRep().FindByYear(year)
}

func (c TournamentController) GetAll() ([]*models.Tournament, error) {
	return c.Reps.TournamentRep().GetAll()
}

func (c TournamentController) FindByName(name string) ([]*models.Tournament, error) {
	return c.Reps.TournamentRep().FindByName(name)
}
