package models

import (
	"errors"
	"strconv"
	"time"
)

type CompanyTeam struct {
	Id           int
	CompanyId    int
	TeamId       int
	ContractDate string
	ContractTime int
}

type CompanyTeamView struct {
	Id           string
	CompanyId    string
	TeamId       string
	ContractDate string
	ContractTime string
}

func (mod CompanyTeam) Check() error {
	if mod.Id <= 0 {
		return errors.New("Некорректный ID")
	}
	if mod.CompanyId <= 0 {
		return errors.New("Некорректный ID Компании")
	}
	if mod.TeamId <= 0 {
		return errors.New("Некорректный ID Команды")
	}
	_, err := time.Parse("2006-01-02", mod.ContractDate)
	if err != nil {
		return errors.New("Некорректная дата контракта")
	}
	if mod.ContractTime <= 0 {
		return errors.New("Некорректная продолжительность")
	}
	return nil
}

func (mod CompanyTeam) ToViewConv() *CompanyTeamView {
	vmod := CompanyTeamView{
		Id:           strconv.Itoa(mod.Id),
		CompanyId:    strconv.Itoa(mod.CompanyId),
		TeamId:       strconv.Itoa(mod.TeamId),
		ContractDate: mod.ContractDate[:10],
		ContractTime: strconv.Itoa(mod.ContractTime),
	}
	return &vmod
}

func (vmod CompanyTeamView) FromViewConv() (CompanyTeam, error) {
	mod := CompanyTeam{}
	id, err := strconv.Atoi(vmod.Id)
	if err != nil {
		return mod, errors.New("Некорректный ID")
	}
	companyId, err := strconv.Atoi(vmod.CompanyId)
	if err != nil {
		return mod, errors.New("Некорректный ID Компании")
	}
	teamId, err := strconv.Atoi(vmod.TeamId)
	if err != nil {
		return mod, errors.New("Некорректный ID Команды")
	}
	contractTime, err := strconv.Atoi(vmod.ContractTime)
	if err != nil {
		return mod, errors.New("Некорректная продолжительность")
	}
	mod.Id = id
	mod.CompanyId = companyId
	mod.TeamId = teamId
	mod.ContractDate = vmod.ContractDate
	mod.ContractTime = contractTime
	return mod, nil
}
