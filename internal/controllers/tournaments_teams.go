package controllers

import (
	"local/internal/logger"
	"local/internal/models"
	"local/internal/repository"
)

type TournamentTeamController struct {
	Reps repository.Repositories
}

type TournamentTeamControllerBuilder struct{}

func (b TournamentTeamControllerBuilder) Build(Reps repository.Repositories) TournamentTeamController {
	return TournamentTeamController{Reps: Reps}
}

func (c TournamentTeamController) Add(mod models.TournamentTeam) error {
	err := mod.Check()
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	err = c.Reps.TournamentTeamRep().Add(mod)
	return err
}

func (c TournamentTeamController) Edit(mod models.TournamentTeam) error {
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
	return c.Reps.TournamentTeamRep().Edit(mod)
}

func (c TournamentTeamController) Delete(id int) error {
	_, err := c.FindById(id)
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	return c.Reps.TournamentTeamRep().Delete(id)
}

func (c TournamentTeamController) FindById(id int) ([]*models.TournamentTeam, error) {
	return c.Reps.TournamentTeamRep().FindById(id)
}

func (c TournamentTeamController) GetAll() ([]*models.TournamentTeam, error) {
	return c.Reps.TournamentTeamRep().GetAll()
}
