package models

import (
	"errors"
	"strconv"
)

type Match struct {
	Id           int
	TournamentId int
	RTeamId      int
	DTeamId      int
	Duration     int
	Winner       bool
}

type MatchView struct {
	Id           string
	TournamentId string
	RTeamId      string
	DTeamId      string
	Duration     string
	Winner       string
}

func (mod Match) ToViewConv() *MatchView {
	vmod := MatchView{
		Id:           strconv.Itoa(mod.Id),
		TournamentId: strconv.Itoa(mod.TournamentId),
		RTeamId:      strconv.Itoa(mod.RTeamId),
		DTeamId:      strconv.Itoa(mod.DTeamId),
		Duration:     strconv.Itoa(mod.Duration),
		Winner:       strconv.FormatBool(mod.Winner),
	}
	return &vmod
}

func (mod Match) Check() error {
	if mod.Id <= 0 {
		return errors.New("Некорректный ID")
	}
	if mod.TournamentId <= 0 {
		return errors.New("Некорректный ID Турнира")
	}
	if mod.RTeamId <= 0 {
		return errors.New("Некорректный ID Команды Света")
	}
	if mod.DTeamId <= 0 {
		return errors.New("Некорректный ID Команды Тьмы")
	}
	if mod.RTeamId == mod.DTeamId {
		return errors.New("Некорректный ID Команды Тьмы")
	}
	if mod.Duration <= 0 {
		return errors.New("Некорректная продолжительность")
	}
	return nil
}

func (vmod MatchView) FromViewConv() (Match, error) {
	mod := Match{}
	id, err := strconv.Atoi(vmod.Id)
	if err != nil {
		return mod, errors.New("Некорректный ID")
	}
	tournamentId, err := strconv.Atoi(vmod.TournamentId)
	if err != nil {
		return mod, errors.New("Некорректный ID Турнира")
	}
	rTeamId, err := strconv.Atoi(vmod.RTeamId)
	if err != nil {
		return mod, errors.New("Некорректный ID Команды Света")
	}
	dTeamId, err := strconv.Atoi(vmod.DTeamId)
	if err != nil {
		return mod, errors.New("Некорректный ID Команды Тьмы")
	}
	duration, err := strconv.Atoi(vmod.Duration)
	if err != nil {
		return mod, errors.New("Некорректная длительность")
	}
	winner, err := strconv.ParseBool(vmod.Winner)
	if err != nil {
		return mod, errors.New("Некорректный победитель")
	}
	mod.Id = id
	mod.TournamentId = tournamentId
	mod.RTeamId = rTeamId
	mod.DTeamId = dTeamId
	mod.Duration = duration
	mod.Winner = winner
	return mod, nil
}
