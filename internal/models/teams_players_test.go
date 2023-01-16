package models_test

import (
	"local/internal/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTeamsPlayersModelsUnit(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		test_struct := models.TeamPlayerView{
			Id:           "1",
			TeamId:       "1",
			PlayerId:     "1",
			ContractDate: "2000-01-01",
			ContractTime: "1"}
		_, err := test_struct.FromViewConv()
		require.Nil(t, err)
	})

	t.Run("incorrect_id", func(t *testing.T) {
		test_struct := models.TeamPlayerView{
			Id:           "a",
			TeamId:       "1",
			PlayerId:     "1",
			ContractDate: "2000-01-01",
			ContractTime: "1"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_team_id", func(t *testing.T) {
		test_struct := models.TeamPlayerView{
			Id:           "1",
			TeamId:       "a",
			PlayerId:     "1",
			ContractDate: "2000-01-01",
			ContractTime: "1"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_player_id", func(t *testing.T) {
		test_struct := models.TeamPlayerView{
			Id:           "1",
			TeamId:       "1",
			PlayerId:     "a",
			ContractDate: "2000-01-01",
			ContractTime: "1"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_contract_time", func(t *testing.T) {
		test_struct := models.TeamPlayerView{
			Id:           "1",
			TeamId:       "1",
			PlayerId:     "1",
			ContractDate: "2000-01-01",
			ContractTime: "a"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

}
