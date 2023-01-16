package models_test

import (
	"local/internal/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTournamentsTeamsModelsUnit(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		test_struct := models.TournamentTeamView{
			Id:                "1",
			TournamentId:      "1",
			TeamId:            "1",
			ParticipationType: "type",
			IsWinner:          "t"}
		_, err := test_struct.FromViewConv()
		require.Nil(t, err)
	})

	t.Run("incorrect_id", func(t *testing.T) {
		test_struct := models.TournamentTeamView{
			Id:                "a",
			TournamentId:      "1",
			TeamId:            "1",
			ParticipationType: "type",
			IsWinner:          "t"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_tournament_id", func(t *testing.T) {
		test_struct := models.TournamentTeamView{
			Id:                "1",
			TournamentId:      "a",
			TeamId:            "1",
			ParticipationType: "type",
			IsWinner:          "t"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_team_id", func(t *testing.T) {
		test_struct := models.TournamentTeamView{
			Id:                "1",
			TournamentId:      "1",
			TeamId:            "a",
			ParticipationType: "type",
			IsWinner:          "t"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_is_winner", func(t *testing.T) {
		test_struct := models.TournamentTeamView{
			Id:                "1",
			TournamentId:      "1",
			TeamId:            "1",
			ParticipationType: "type",
			IsWinner:          "a"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

}
