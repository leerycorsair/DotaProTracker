package models

import (
	"errors"
	"strconv"
)

type Company struct {
	Id       int
	Name     string
	Country  string
	Website  string
	Revenue  int
	Industry string
}

type CompanyView struct {
	Id       string
	Name     string
	Country  string
	Website  string
	Revenue  string
	Industry string
}

func (mod Company) ToViewConv() *CompanyView {
	vmod := CompanyView{
		Id:       strconv.Itoa(mod.Id),
		Name:     mod.Name,
		Country:  mod.Country,
		Website:  mod.Website,
		Revenue:  strconv.Itoa(mod.Revenue),
		Industry: mod.Industry,
	}
	return &vmod
}

func (mod Company) Check() error {
	if mod.Id <= 0 {
		return errors.New("Некорректный ID")
	}
	if mod.Revenue <= 0 {
		return errors.New("Некорректный оборот")
	}
	return nil
}

func (vmod CompanyView) FromViewConv() (Company, error) {
	mod := Company{}
	id, err := strconv.Atoi(vmod.Id)
	if err != nil {
		return mod, errors.New("Некорректный ID")
	}
	if len(vmod.Name) == 0 {
		return mod, errors.New("Некорректное имя")
	}
	if len(vmod.Country) == 0 {
		return mod, errors.New("Некорректная страна")
	}
	if len(vmod.Website) == 0 {
		return mod, errors.New("Некорректный сайт")
	}
	revenue, err := strconv.Atoi(vmod.Revenue)
	if err != nil {
		return mod, errors.New("Некорректный оборот")
	}
	if len(vmod.Industry) == 0 {
		return mod, errors.New("Некорректная индустрия")
	}
	mod.Id = id
	mod.Name = vmod.Name
	mod.Country = vmod.Country
	mod.Website = vmod.Website
	mod.Revenue = revenue
	mod.Industry = vmod.Industry
	return mod, nil
}
