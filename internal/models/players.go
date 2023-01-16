package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Player struct {
	Id            int
	Nickname      string
	Realname      string
	Birthdate     string
	Country       string
	MMR           int
	Role          string
	SignatureHero string
}

type PlayerView struct {
	Id            string
	Nickname      string
	Realname      string
	Birthdate     string
	Country       string
	MMR           string
	Role          string
	SignatureHero string
}

func (mod Player) Check() error {
	if mod.Id <= 0 {
		return errors.New("Некорректный ID")
	}
	_, err := time.Parse("2006-01-02", mod.Birthdate)
	if err != nil {
		return errors.New("Некорректная дата рождения")
	}
	if mod.MMR <= 0 {
		return errors.New("Некорректный Рейтинг")
	}
	return nil
}

func date(date string) string {
	str, _ := time.Parse("2006-01-02", date)
	fmt.Print(date)
	return str.Format("2006-01-02")
}

func (mod Player) ToViewConv() *PlayerView {
	vmod := PlayerView{
		Id:            strconv.Itoa(mod.Id),
		Nickname:      mod.Nickname,
		Realname:      mod.Realname,
		Birthdate:     mod.Birthdate[:10],
		Country:       mod.Country,
		MMR:           strconv.Itoa(mod.MMR),
		Role:          mod.Role,
		SignatureHero: mod.SignatureHero,
	}
	return &vmod
}

func (vmod PlayerView) FromViewConv() (Player, error) {
	mod := Player{}
	id, err := strconv.Atoi(vmod.Id)
	if err != nil {
		return mod, errors.New("Некорректный ID")
	}
	if len(vmod.Nickname) == 0 {
		return mod, errors.New("Некорректное прозвище")
	}
	if len(vmod.Realname) == 0 {
		return mod, errors.New("Некорректное имя")
	}
	if len(vmod.Country) == 0 {
		return mod, errors.New("Некорректная страна")
	}
	mmr, err := strconv.Atoi(vmod.MMR)
	if err != nil {
		return mod, errors.New("Некорректный рейтинг")
	}
	if len(vmod.Role) == 0 {
		return mod, errors.New("Некорректная роль")
	}
	if len(vmod.SignatureHero) == 0 {
		return mod, errors.New("Некорректный сигнатурный герой")
	}

	mod.Id = id
	mod.Nickname = vmod.Nickname
	mod.Realname = vmod.Realname
	mod.Birthdate = vmod.Birthdate
	mod.Country = vmod.Country
	mod.MMR = mmr
	mod.Role = vmod.Role
	mod.SignatureHero = vmod.SignatureHero
	return mod, nil
}
