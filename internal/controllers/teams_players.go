package controllers

import (
	"local/internal/logger"
	"local/internal/models"
	"local/internal/repository"
)

type TeamPlayerController struct {
	Reps repository.Repositories
}

type TeamPlayerControllerBuilder struct{}

func (b TeamPlayerControllerBuilder) Build(Reps repository.Repositories) TeamPlayerController {
	return TeamPlayerController{Reps: Reps}
}

func (c TeamPlayerController) Add(mod models.TeamPlayer) error {
	err := mod.Check()
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	err = c.Reps.TeamPlayerRep().Add(mod)
	return err
}

func (c TeamPlayerController) Edit(mod models.TeamPlayer) error {
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
	return c.Reps.TeamPlayerRep().Edit(mod)
}

func (c TeamPlayerController) Delete(id int) error {
	_, err := c.FindById(id)
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	return c.Reps.TeamPlayerRep().Delete(id)
}

func (c TeamPlayerController) FindById(id int) ([]*models.TeamPlayer, error) {
	return c.Reps.TeamPlayerRep().FindById(id)
}

func (c TeamPlayerController) GetAll() ([]*models.TeamPlayer, error) {
	return c.Reps.TeamPlayerRep().GetAll()
}
