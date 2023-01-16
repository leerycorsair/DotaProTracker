package models

import (
	"errors"
	"strconv"
)

type CompanyTournament struct {
	Id           int
	CompanyId    int
	TournamentId int
	Deposit      int
}

type CompanyTournamentView struct {
	Id           string
	CompanyId    string
	TournamentId string
	Deposit      string
}

func (mod CompanyTournament) ToViewConv() *CompanyTournamentView {
	vmod := CompanyTournamentView{
		Id:           strconv.Itoa(mod.Id),
		CompanyId:    strconv.Itoa(mod.CompanyId),
		TournamentId: strconv.Itoa(mod.TournamentId),
		Deposit:      strconv.Itoa(mod.Deposit),
	}
	return &vmod
}

func (mod CompanyTournament) Check() error {
	if mod.Id <= 0 {
		return errors.New("Некорректный ID")
	}
	if mod.CompanyId <= 0 {
		return errors.New("Некорректный ID Компании")
	}
	if mod.TournamentId <= 0 {
		return errors.New("Некорректный ID Турнира")
	}
	if mod.Deposit <= 0 {
		return errors.New("Некорректный депозит")
	}
	return nil
}

func (vmod CompanyTournamentView) FromViewConv() (CompanyTournament, error) {
	mod := CompanyTournament{}
	id, err := strconv.Atoi(vmod.Id)
	if err != nil {
		return mod, errors.New("Некорректный ID")
	}
	companyId, err := strconv.Atoi(vmod.CompanyId)
	if err != nil {
		return mod, errors.New("Некорректный ID Компании")
	}
	tournamentId, err := strconv.Atoi(vmod.TournamentId)
	if err != nil {
		return mod, errors.New("Некорректный ID Турнира")
	}
	deposit, err := strconv.Atoi(vmod.Deposit)
	if err != nil {
		return mod, errors.New("Некорректный депозит")
	}
	mod.Id = id
	mod.CompanyId = companyId
	mod.TournamentId = tournamentId
	mod.Deposit = deposit
	return mod, nil
}
