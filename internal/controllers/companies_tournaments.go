package controllers

import (
	"local/internal/logger"
	"local/internal/models"
	"local/internal/repository"
)

type CompanyTournamentController struct {
	Reps repository.Repositories
}

type CompanyTournamentControllerBuilder struct{}

func (b CompanyTournamentControllerBuilder) Build(Reps repository.Repositories) CompanyTournamentController {
	return CompanyTournamentController{Reps: Reps}
}

func (c CompanyTournamentController) Add(mod models.CompanyTournament) error {
	err := mod.Check()
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	err = c.Reps.CompanyTournamentRep().Add(mod)
	return err
}

func (c CompanyTournamentController) Edit(mod models.CompanyTournament) error {
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
	return c.Reps.CompanyTournamentRep().Edit(mod)
}

func (c CompanyTournamentController) Delete(id int) error {
	_, err := c.FindById(id)
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	return c.Reps.CompanyTournamentRep().Delete(id)
}

func (c CompanyTournamentController) FindById(id int) ([]*models.CompanyTournament, error) {
	return c.Reps.CompanyTournamentRep().FindById(id)
}

func (c CompanyTournamentController) GetAll() ([]*models.CompanyTournament, error) {
	return c.Reps.CompanyTournamentRep().GetAll()
}
