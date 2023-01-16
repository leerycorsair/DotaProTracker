package managers

import (
	"errors"
	"local/internal/controllers"
	"local/internal/logger"
	"local/internal/models"
	"strconv"
)

type UserManager struct {
	C controllers.UserController
}

type UserManagerBuilder struct {
}

func (b UserManagerBuilder) Build(C controllers.UserController) UserManager {
	return UserManager{C: C}
}

func (man *UserManager) Add(vmod models.UserView) error {
	mod, err := vmod.FromViewConv()
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	err = man.C.Add(mod)
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	return nil
}

func (man *UserManager) Edit(vmod models.UserView) error {
	mod, err := vmod.FromViewConv()
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	err = man.C.Edit(mod)
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	return nil
}

func (man *UserManager) Delete(strid string) error {
	id, err := strconv.Atoi(strid)
	if err != nil {
		logger.Logger.Println(err)
		return errors.New("Некорректный ID")
	}
	err = man.C.Delete(id)
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	return nil
}

func (man *UserManager) GetAll() ([]*models.UserView, error) {
	vrecs := []*models.UserView{}
	recs, err := man.C.GetAll()
	if err != nil {
		logger.Logger.Println(err)
		return nil, err
	}
	for i := 0; i < len(recs); i++ {
		vrecs = append(vrecs, (recs[i].ToViewConv()))
	}
	return vrecs, nil
}

func (man *UserManager) FindById(strid string) ([]*models.UserView, error) {
	id, err := strconv.Atoi(strid)
	if err != nil {
		logger.Logger.Println(err)
		return nil, errors.New("Некорректный ID")
	}
	vrecs := []*models.UserView{}
	recs, err := man.C.FindById(id)
	if err != nil {
		logger.Logger.Println(err)
		return nil, err
	}
	for i := 0; i < len(recs); i++ {
		vrecs = append(vrecs, (recs[i].ToViewConv()))
	}
	return vrecs, nil
}

func (man *UserManager) Login(vmod models.UserView) (int, error) {
	mod, err := vmod.FromViewConv()
	if err != nil {
		logger.Logger.Println(err)
		return 1, err
	}
	priv_level, err := man.C.Login(mod)
	if err != nil {
		logger.Logger.Println(err)
		return 1, err
	}
	return priv_level, nil

}
