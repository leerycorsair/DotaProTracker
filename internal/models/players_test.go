package models_test

import (
	"local/internal/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlayersModelsUnit(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		test_struct := models.PlayerView{
			Id:            "1",
			Nickname:      "nickname",
			Realname:      "realname",
			Birthdate:     "2000-01-01",
			Country:       "country",
			MMR:           "1",
			Role:          "role",
			SignatureHero: "pudge"}
		_, err := test_struct.FromViewConv()
		require.Nil(t, err)
	})

	t.Run("incorrect_id", func(t *testing.T) {
		test_struct := models.PlayerView{
			Id:            "a",
			Nickname:      "nickname",
			Realname:      "realname",
			Birthdate:     "2000-01-01",
			Country:       "country",
			MMR:           "1",
			Role:          "role",
			SignatureHero: "pudge"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_nickname", func(t *testing.T) {
		test_struct := models.PlayerView{
			Id:            "1",
			Nickname:      "",
			Realname:      "realname",
			Birthdate:     "2000-01-01",
			Country:       "country",
			MMR:           "1",
			Role:          "role",
			SignatureHero: "pudge"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})
	t.Run("incorrect_realname", func(t *testing.T) {
		test_struct := models.PlayerView{
			Id:            "1",
			Nickname:      "nickname",
			Realname:      "",
			Birthdate:     "2000-01-01",
			Country:       "country",
			MMR:           "1",
			Role:          "role",
			SignatureHero: "pudge"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})
	t.Run("incorrect_country", func(t *testing.T) {
		test_struct := models.PlayerView{
			Id:            "1",
			Nickname:      "nickname",
			Realname:      "realname",
			Birthdate:     "2000-01-01",
			Country:       "",
			MMR:           "1",
			Role:          "role",
			SignatureHero: "pudge"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})
	t.Run("incorrect_mmr", func(t *testing.T) {
		test_struct := models.PlayerView{
			Id:            "1",
			Nickname:      "nickname",
			Realname:      "realname",
			Birthdate:     "2000-01-01",
			Country:       "country",
			MMR:           "a",
			Role:          "role",
			SignatureHero: "pudge"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})
	t.Run("incorrect_role", func(t *testing.T) {
		test_struct := models.PlayerView{
			Id:            "1",
			Nickname:      "nickname",
			Realname:      "realname",
			Birthdate:     "2000-01-01",
			Country:       "country",
			MMR:           "1",
			Role:          "",
			SignatureHero: "pudge"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})
	t.Run("incorrect_hero", func(t *testing.T) {
		test_struct := models.PlayerView{
			Id:            "1",
			Nickname:      "nickname",
			Realname:      "realname",
			Birthdate:     "2000-01-01",
			Country:       "country",
			MMR:           "1",
			Role:          "role",
			SignatureHero: ""}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})
}
