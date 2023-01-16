package models

import (
	"errors"
	"strconv"
	"time"
)

type TeamPlayer struct {
	Id           int
	TeamId       int
	PlayerId     int
	ContractDate string
	ContractTime int
}

type TeamPlayerView struct {
	Id           string
	TeamId       string
	PlayerId     string
	ContractDate string
	ContractTime string
}

func (mod TeamPlayer) ToViewConv() *TeamPlayerView {
	vmod := TeamPlayerView{
		Id:           strconv.Itoa(mod.Id),
		TeamId:       strconv.Itoa(mod.TeamId),
		PlayerId:     strconv.Itoa(mod.PlayerId),
		ContractDate: mod.ContractDate[:10],
		ContractTime: strconv.Itoa(mod.ContractTime),
	}
	return &vmod
}

func (mod TeamPlayer) Check() error {
	if mod.Id <= 0 {
		return errors.New("Некорректный ID")
	}
	if mod.TeamId <= 0 {
		return errors.New("Некорректный ID Команды")
	}
	if mod.PlayerId <= 0 {
		return errors.New("Некорректный ID Игрока")
	}
	_, err := time.Parse("2006-01-02", mod.ContractDate)
	if err != nil {
		return errors.New("Некорректная дата контракта")
	}
	if mod.ContractTime <= 0 {
		return errors.New("Некорректное время контракта")
	}
	return nil
}

func (vmod TeamPlayerView) FromViewConv() (TeamPlayer, error) {
	mod := TeamPlayer{}
	id, err := strconv.Atoi(vmod.Id)
	if err != nil {
		return mod, errors.New("Некорректный ID")
	}
	teamId, err := strconv.Atoi(vmod.TeamId)
	if err != nil {
		return mod, errors.New("Некорректный ID Команды")
	}
	playerId, err := strconv.Atoi(vmod.PlayerId)
	if err != nil {
		return mod, errors.New("Некорректный ID Игрока")
	}
	contractTime, err := strconv.Atoi(vmod.ContractTime)
	if err != nil {
		return mod, errors.New("Некорректное время контракта")
	}
	mod.Id = id
	mod.TeamId = teamId
	mod.PlayerId = playerId
	mod.ContractDate = vmod.ContractDate
	mod.ContractTime = contractTime
	return mod, nil
}
