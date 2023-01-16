package models

import (
	"errors"
	"strconv"
	"time"
)

type Tournament struct {
	Id        int
	Name      string
	Tier      int
	PrizePool int
	DateStart string
	Duration  int
	DPCPoints int
	Location  string
}

type TournamentView struct {
	Id        string
	Name      string
	Tier      string
	PrizePool string
	DateStart string
	Duration  string
	DPCPoints string
	Location  string
}

func (mod Tournament) ToViewConv() *TournamentView {
	vmod := TournamentView{
		Id:        strconv.Itoa(mod.Id),
		Name:      mod.Name,
		Tier:      strconv.Itoa(mod.Tier),
		PrizePool: strconv.Itoa(mod.PrizePool),
		DateStart: mod.DateStart[:10],
		Duration:  strconv.Itoa(mod.Duration),
		DPCPoints: strconv.Itoa(mod.DPCPoints),
		Location:  mod.Location,
	}
	return &vmod
}

func (mod Tournament) Check() error {
	if mod.Id <= 0 {
		return errors.New("Некорректный ID")
	}
	if mod.Tier <= 0 {
		return errors.New("Некорректный уровень")
	}
	if mod.PrizePool <= 0 {
		return errors.New("Некорректный призовой фонд")
	}
	_, err := time.Parse("2006-01-02", mod.DateStart)
	if err != nil {
		return errors.New("Некорректная дата проведения")
	}
	if mod.Duration <= 0 {
		return errors.New("Некорректная продолжительность")
	}
	if mod.DPCPoints <= 0 {
		return errors.New("Некорректные DPC-очки")
	}
	return nil
}

func (vmod TournamentView) FromViewConv() (Tournament, error) {
	mod := Tournament{}
	id, err := strconv.Atoi(vmod.Id)
	if err != nil {
		return mod, errors.New("Некорректный ID")
	}
	if len(vmod.Name) == 0 {
		return mod, errors.New("Некорректное название")
	}
	tier, err := strconv.Atoi(vmod.Tier)
	if err != nil {
		return mod, errors.New("Некорректный уровень")
	}
	prizePool, err := strconv.Atoi(vmod.PrizePool)
	if err != nil {
		return mod, errors.New("Некорректный призовой фонд")
	}
	duration, err := strconv.Atoi(vmod.Duration)
	if err != nil {
		return mod, errors.New("Некорректная продолжительность")
	}
	dpcPoints, err := strconv.Atoi(vmod.DPCPoints)
	if err != nil {
		return mod, errors.New("Некорректные DPC-очки")
	}
	if len(vmod.Location) == 0 {
		return mod, errors.New("Некорректное место проведения")
	}

	mod.Id = id
	mod.Name = vmod.Name
	mod.Tier = tier
	mod.PrizePool = prizePool
	mod.DateStart = vmod.DateStart
	mod.Duration = duration
	mod.DPCPoints = dpcPoints
	mod.Location = vmod.Location

	return mod, nil
}
