package models_test

import (
	"local/internal/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMatchesModels(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		test_struct := models.MatchView{
			Id:           "1",
			TournamentId: "1",
			RTeamId:      "1",
			DTeamId:      "1",
			Duration:     "1",
			Winner:       "t"}
		_, err := test_struct.FromViewConv()
		require.Nil(t, err)
	})

	t.Run("incorrect_id", func(t *testing.T) {
		test_struct := models.MatchView{
			Id:           "a",
			TournamentId: "1",
			RTeamId:      "1",
			DTeamId:      "1",
			Duration:     "1",
			Winner:       "t"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_tournament_id", func(t *testing.T) {
		test_struct := models.MatchView{
			Id:           "1",
			TournamentId: "a",
			RTeamId:      "1",
			DTeamId:      "1",
			Duration:     "1",
			Winner:       "t"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})
	t.Run("incorrect_rteam_id", func(t *testing.T) {
		test_struct := models.MatchView{
			Id:           "1",
			TournamentId: "1",
			RTeamId:      "a",
			DTeamId:      "1",
			Duration:     "1",
			Winner:       "t"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})
	t.Run("incorrect_dteam_id", func(t *testing.T) {
		test_struct := models.MatchView{
			Id:           "1",
			TournamentId: "1",
			RTeamId:      "1",
			DTeamId:      "a",
			Duration:     "1",
			Winner:       "t"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})
	t.Run("incorrect_duration", func(t *testing.T) {
		test_struct := models.MatchView{
			Id:           "1",
			TournamentId: "1",
			RTeamId:      "1",
			DTeamId:      "1",
			Duration:     "a",
			Winner:       "t"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})
	t.Run("incorrect_winner", func(t *testing.T) {
		test_struct := models.MatchView{
			Id:           "1",
			TournamentId: "1",
			RTeamId:      "1",
			DTeamId:      "1",
			Duration:     "1",
			Winner:       "fdt"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})
}
