package models

import (
	"errors"
	"strconv"
	"time"
)

type Team struct {
	Id            int
	Name          string
	CreatedAt     string
	Email         string
	TotalEarnings int
	Region        string
	Tier          int
}

type TeamView struct {
	Id            string
	Name          string
	CreatedAt     string
	Email         string
	TotalEarnings string
	Region        string
	Tier          string
}

func (mod Team) ToViewConv() *TeamView {
	vmod := TeamView{
		Id:            strconv.Itoa(mod.Id),
		Name:          mod.Name,
		CreatedAt:     mod.CreatedAt[:10],
		Email:         mod.Email,
		TotalEarnings: strconv.Itoa(mod.TotalEarnings),
		Region:        mod.Region,
		Tier:          strconv.Itoa(mod.Tier),
	}
	return &vmod
}

func (mod Team) Check() error {
	if mod.Id <= 0 {
		return errors.New("Некорректный ID")
	}
	_, err := time.Parse("2006-01-02", mod.CreatedAt)
	if err != nil {
		return errors.New("Некорректная дата создания")
	}
	if mod.TotalEarnings <= 0 {
		return errors.New("Некорректный общий заработок")
	}
	if mod.Tier <= 0 {
		return errors.New("Некорректный уровень")
	}
	return nil
}

func (vmod TeamView) FromViewConv() (Team, error) {
	mod := Team{}
	id, err := strconv.Atoi(vmod.Id)
	if err != nil {
		return mod, errors.New("Некорректный ID")
	}
	if len(vmod.Name) == 0 {
		return mod, errors.New("Некорректное название")
	}
	if len(vmod.Email) == 0 {
		return mod, errors.New("Некорректная почта")
	}
	totalEarnings, err := strconv.Atoi(vmod.TotalEarnings)
	if err != nil {
		return mod, errors.New("Некорректный общий заработок")
	}
	if len(vmod.Region) == 0 {
		return mod, errors.New("Некорректный регион")
	}
	tier, err := strconv.Atoi(vmod.Tier)
	if err != nil {
		return mod, errors.New("Некорректный уровень")
	}

	mod.Id = id
	mod.Name = vmod.Name
	mod.CreatedAt = vmod.CreatedAt
	mod.Email = vmod.Email
	mod.TotalEarnings = totalEarnings
	mod.Region = vmod.Region
	mod.Tier = tier
	return mod, nil
}
