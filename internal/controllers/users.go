package controllers

import (
	"errors"
	"local/internal/logger"
	"local/internal/models"
	"local/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	Reps repository.Repositories
}

type UserControllerBuilder struct{}

func (b UserControllerBuilder) Build(Reps repository.Repositories) UserController {
	return UserController{Reps: Reps}
}

func (c UserController) Add(mod models.User) error {
	err := mod.Check()
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	mod.Password, err = EncryptString(mod.Password)
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	err = c.Reps.UserRep().Add(mod)
	return err
}

func (c UserController) Edit(mod models.User) error {
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
	mod.Password, err = EncryptString(mod.Password)
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	return c.Reps.UserRep().Edit(mod)
}

func (c UserController) Delete(id int) error {
	_, err := c.FindById(id)
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	return c.Reps.UserRep().Delete(id)
}

func (c UserController) FindById(id int) ([]*models.User, error) {
	return c.Reps.UserRep().FindById(id)
}

func (c UserController) GetAll() ([]*models.User, error) {
	return c.Reps.UserRep().GetAll()
}

func (c UserController) Login(mod models.User) (int, error) {
	usr, err := c.Reps.UserRep().FindByLogin(mod.Login)
	if err != nil {
		logger.Logger.Println(err)
		return 1, err
	}
	if ComparePassword(usr[0].Password, mod.Password) != nil {
		return 1, errors.New("Incorrect Password")
	}
	return usr[0].PrivilegeLevel, nil

}

func ComparePassword(enc_password string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(enc_password), []byte(password))
}

func EncryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		logger.Logger.Println(err)
		return "", err
	}
	return string(b), nil
}
