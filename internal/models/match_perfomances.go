package models

import (
	"errors"
	"strconv"
)

type MatchPerfomance struct {
	Id       int
	MatchId  int
	PlayerId int
	Team     bool
	Hero     string
	Kills    int
	Deaths   int
	Assists  int
	Networth int
	GPM      int
	XPM      int
	DMG      int
	Heal     int
	BLD      int
}

type MatchPerfomanceView struct {
	Id       string
	MatchId  string
	PlayerId string
	Team     string
	Hero     string
	Kills    string
	Deaths   string
	Assists  string
	Networth string
	GPM      string
	XPM      string
	DMG      string
	Heal     string
	BLD      string
}

func (mod MatchPerfomance) ToViewConv() *MatchPerfomanceView {
	vmod := MatchPerfomanceView{
		Id:       strconv.Itoa(mod.Id),
		MatchId:  strconv.Itoa(mod.MatchId),
		PlayerId: strconv.Itoa(mod.PlayerId),
		Team:     strconv.FormatBool(mod.Team),
		Hero:     mod.Hero,
		Kills:    strconv.Itoa(mod.Kills),
		Deaths:   strconv.Itoa(mod.Deaths),
		Assists:  strconv.Itoa(mod.Assists),
		Networth: strconv.Itoa(mod.Networth),
		GPM:      strconv.Itoa(mod.GPM),
		XPM:      strconv.Itoa(mod.XPM),
		DMG:      strconv.Itoa(mod.DMG),
		Heal:     strconv.Itoa(mod.Heal),
		BLD:      strconv.Itoa(mod.BLD),
	}
	return &vmod
}

func (mod MatchPerfomance) Check() error {
	if mod.Id <= 0 {
		return errors.New("Некорректный ID")
	}
	if mod.MatchId <= 0 {
		return errors.New("Некорректный ID Матча")
	}
	if mod.PlayerId <= 0 {
		return errors.New("Некорректный ID Игрока")
	}
	if mod.Kills < 0 {
		return errors.New("Некорректные убийства")
	}
	if mod.Deaths < 0 {
		return errors.New("Некорректные смерти")
	}
	if mod.Assists < 0 {
		return errors.New("Некорректные помощи")
	}
	if mod.Networth <= 0 {
		return errors.New("Некорректная ценность")
	}
	if mod.GPM <= 0 {
		return errors.New("Некорректный ЗВМ")
	}
	if mod.XPM <= 0 {
		return errors.New("Некорректный ОВМ")
	}
	if mod.DMG < 0 {
		return errors.New("Некорректный урон")
	}
	if mod.Heal < 0 {
		return errors.New("Некорретное лечение")
	}
	if mod.BLD < 0 {
		return errors.New("Некорретный урон по постройкам")
	}
	return nil
}

func (vmod MatchPerfomanceView) FromViewConv() (MatchPerfomance, error) {
	mod := MatchPerfomance{}
	id, err := strconv.Atoi(vmod.Id)
	if err != nil {
		return mod, errors.New("Некорректный ID")
	}
	matchId, err := strconv.Atoi(vmod.MatchId)
	if err != nil {
		return mod, errors.New("Некорректный ID Матча")
	}
	team, err := strconv.ParseBool(vmod.Team)
	if err != nil {
		return mod, errors.New("Некорректная команда")
	}
	if len(vmod.Hero) == 0 {
		return mod, errors.New("Некорректный герой")
	}
	playerId, err := strconv.Atoi(vmod.PlayerId)
	if err != nil {
		return mod, errors.New("Некорректный ID Игрока")
	}
	kills, err := strconv.Atoi(vmod.Kills)
	if err != nil {
		return mod, errors.New("Некорректные убийства")
	}
	deaths, err := strconv.Atoi(vmod.Deaths)
	if err != nil {
		return mod, errors.New("Некорректные смерти")
	}
	assists, err := strconv.Atoi(vmod.Assists)
	if err != nil {
		return mod, errors.New("Некорректные помощи")
	}
	networth, err := strconv.Atoi(vmod.Networth)
	if err != nil {
		return mod, errors.New("Некорректная ценность")
	}
	gpm, err := strconv.Atoi(vmod.GPM)
	if err != nil {
		return mod, errors.New("Некорректный ЗВМ")
	}
	xpm, err := strconv.Atoi(vmod.XPM)
	if err != nil {
		return mod, errors.New("Некорректный ОВМ")
	}
	dmg, err := strconv.Atoi(vmod.DMG)
	if err != nil {
		return mod, errors.New("Некорректный урон")
	}
	heal, err := strconv.Atoi(vmod.Heal)
	if err != nil {
		return mod, errors.New("Некорректное лечение")
	}
	bld, err := strconv.Atoi(vmod.BLD)
	if err != nil {
		return mod, errors.New("Некорректный урон по постройкам")
	}
	mod.Id = id
	mod.MatchId = matchId
	mod.PlayerId = playerId
	mod.Team = team
	mod.Hero = vmod.Hero
	mod.Kills = kills
	mod.Deaths = deaths
	mod.Assists = assists
	mod.Networth = networth
	mod.GPM = gpm
	mod.XPM = xpm
	mod.DMG = dmg
	mod.Heal = heal
	mod.BLD = bld
	return mod, nil
}
