package models

import (
	"errors"
	"strconv"
)

type TournamentTeam struct {
	Id                int
	TournamentId      int
	TeamId            int
	ParticipationType string
	IsWinner          bool
}

type TournamentTeamView struct {
	Id                string
	TournamentId      string
	TeamId            string
	ParticipationType string
	IsWinner          string
}

func (mod TournamentTeam) ToViewConv() *TournamentTeamView {
	vmod := TournamentTeamView{
		Id:                strconv.Itoa(mod.Id),
		TournamentId:      strconv.Itoa(mod.TournamentId),
		TeamId:            strconv.Itoa(mod.TeamId),
		ParticipationType: mod.ParticipationType,
		IsWinner:          strconv.FormatBool(mod.IsWinner),
	}
	return &vmod
}

func (mod TournamentTeam) Check() error {
	if mod.Id <= 0 {
		return errors.New("Некорректный ID")
	}
	if mod.TournamentId <= 0 {
		return errors.New("Некорректный ID Турнира")
	}
	if mod.TeamId <= 0 {
		return errors.New("Некорректный ID Команды")
	}
	if mod.ParticipationType != "invite" && mod.ParticipationType != "qualification" {
		return errors.New("Некорректный способ участия")
	}
	return nil
}

func (vmod TournamentTeamView) FromViewConv() (TournamentTeam, error) {
	mod := TournamentTeam{}
	id, err := strconv.Atoi(vmod.Id)
	if err != nil {
		return mod, errors.New("Некорректный ID")
	}
	tournamentId, err := strconv.Atoi(vmod.TournamentId)
	if err != nil {
		return mod, errors.New("Некорректный ID Турнира")
	}
	teamId, err := strconv.Atoi(vmod.TeamId)
	if err != nil {
		return mod, errors.New("Некорректный ID Команды")
	}
	isWinner, err := strconv.ParseBool(vmod.IsWinner)
	if err != nil {
		return mod, errors.New("Некорректный Победитель")
	}

	mod.Id = id
	mod.TournamentId = tournamentId
	mod.TeamId = teamId
	mod.ParticipationType = vmod.ParticipationType
	mod.IsWinner = isWinner
	return mod, nil
}
