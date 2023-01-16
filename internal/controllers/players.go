package controllers

import (
	"local/internal/logger"
	"local/internal/models"
	"local/internal/repository"
)

type PlayerController struct {
	Reps repository.Repositories
}

type PlayerControllerBuilder struct{}

func (b PlayerControllerBuilder) Build(Reps repository.Repositories) PlayerController {
	return PlayerController{Reps: Reps}
}

func (c PlayerController) Add(mod models.Player) error {
	err := mod.Check()
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	err = c.Reps.PlayerRep().Add(mod)
	return err
}

func (c PlayerController) Edit(mod models.Player) error {
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
	return c.Reps.PlayerRep().Edit(mod)
}

func (c PlayerController) Delete(id int) error {
	_, err := c.FindById(id)
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	return c.Reps.PlayerRep().Delete(id)
}

func (c PlayerController) FindById(id int) ([]*models.Player, error) {
	return c.Reps.PlayerRep().FindById(id)
}

func (c PlayerController) FindByBirthdate(year int) ([]*models.Player, error) {
	return c.Reps.PlayerRep().FindByBirthdate(year)
}

func (c PlayerController) GetAll() ([]*models.Player, error) {
	return c.Reps.PlayerRep().GetAll()
}

func (c PlayerController) FindByName(name string) ([]*models.Player, error) {
	return c.Reps.PlayerRep().FindByName(name)
}
