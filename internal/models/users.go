package models

import (
	"errors"
	"strconv"
	"time"
)

type User struct {
	Id             int
	Name           string
	Birthdate      string
	Login          string
	Password       string
	Email          string
	PrivilegeLevel int
}

type UserView struct {
	Id             string
	Name           string
	Birthdate      string
	Login          string
	Password       string
	Email          string
	PrivilegeLevel string
}

func (mod User) ToViewConv() *UserView {
	vmod := UserView{
		Id:             strconv.Itoa(mod.Id),
		Name:           mod.Name,
		Birthdate:      mod.Birthdate[:10],
		Login:          mod.Login,
		Password:       mod.Password,
		Email:          mod.Email,
		PrivilegeLevel: strconv.Itoa(mod.PrivilegeLevel),
	}
	return &vmod
}

func (mod User) Check() error {
	if mod.Id <= 0 {
		return errors.New("Некорректный ID")
	}
	_, err := time.Parse("2006-01-02", mod.Birthdate)
	if err != nil {
		return errors.New("Некорректная дата рождения")
	}
	if len(mod.Password) < 8 {
		return errors.New("Длина пароля >= 8")
	}
	if mod.PrivilegeLevel < 1 || mod.PrivilegeLevel > 3 {
		return errors.New("Некорректный уровень привилегий")
	}
	return nil
}

func (vmod UserView) FromViewConv() (User, error) {
	mod := User{}
	id, err := strconv.Atoi(vmod.Id)
	if err != nil {
		return mod, errors.New("Некорректный ID")
	}
	if len(vmod.Name) == 0 {
		return mod, errors.New("Некорректное имя")
	}
	if len(vmod.Login) == 0 {
		return mod, errors.New("Некорректный логин")
	}
	if len(vmod.Password) == 0 {
		return mod, errors.New("Некорректный пароль")
	}
	if len(vmod.Email) == 0 {
		return mod, errors.New("Некорректная почта")
	}
	privilegeLevel, err := strconv.Atoi(vmod.PrivilegeLevel)
	if err != nil {
		return mod, errors.New("Некорректный уровень привилегий")
	}

	mod.Id = id
	mod.Name = vmod.Name
	mod.Birthdate = vmod.Birthdate
	mod.Login = vmod.Login
	mod.Password = vmod.Password
	mod.Email = vmod.Email
	mod.PrivilegeLevel = privilegeLevel

	return mod, nil
}
