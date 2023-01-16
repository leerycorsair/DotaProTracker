package models_test

import (
	"local/internal/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUsersModelsUnit(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		test_struct := models.UserView{
			Id:             "1",
			Name:           "name",
			Birthdate:      "2000-01-01",
			Login:          "login",
			Password:       "password",
			Email:          "email",
			PrivilegeLevel: "1"}
		_, err := test_struct.FromViewConv()
		require.Nil(t, err)
	})

	t.Run("incorrect_id", func(t *testing.T) {
		test_struct := models.UserView{
			Id:             "a",
			Name:           "name",
			Birthdate:      "2000-01-01",
			Login:          "login",
			Password:       "password",
			Email:          "email",
			PrivilegeLevel: "1"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_name", func(t *testing.T) {
		test_struct := models.UserView{
			Id:             "1",
			Name:           "",
			Birthdate:      "2000-01-01",
			Login:          "login",
			Password:       "password",
			Email:          "email",
			PrivilegeLevel: "1"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_login", func(t *testing.T) {
		test_struct := models.UserView{
			Id:             "1",
			Name:           "name",
			Birthdate:      "2000-01-01",
			Login:          "",
			Password:       "password",
			Email:          "email",
			PrivilegeLevel: "1"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_password", func(t *testing.T) {
		test_struct := models.UserView{
			Id:             "1",
			Name:           "name",
			Birthdate:      "2000-01-01",
			Login:          "login",
			Password:       "",
			Email:          "email",
			PrivilegeLevel: "1"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_email", func(t *testing.T) {
		test_struct := models.UserView{
			Id:             "1",
			Name:           "name",
			Birthdate:      "2000-01-01",
			Login:          "login",
			Password:       "password",
			Email:          "",
			PrivilegeLevel: "1"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_privilege_level", func(t *testing.T) {
		test_struct := models.UserView{
			Id:             "1",
			Name:           "name",
			Birthdate:      "2000-01-01",
			Login:          "login",
			Password:       "password",
			Email:          "email",
			PrivilegeLevel: "a"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

}
