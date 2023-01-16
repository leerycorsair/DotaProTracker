package models_test

import (
	"local/internal/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompaniesTournamentsModels(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		test_struct := models.CompanyTournamentView{
			Id:           "1",
			CompanyId:    "1",
			TournamentId: "1",
			Deposit:      "1"}
		_, err := test_struct.FromViewConv()
		require.Nil(t, err)
	})

	t.Run("incorrect_id", func(t *testing.T) {
		test_struct := models.CompanyTournamentView{
			Id:           "a",
			CompanyId:    "1",
			TournamentId: "1",
			Deposit:      "1"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_company_id", func(t *testing.T) {
		test_struct := models.CompanyTournamentView{
			Id:           "1",
			CompanyId:    "a",
			TournamentId: "1",
			Deposit:      "1"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_tournament_id", func(t *testing.T) {
		test_struct := models.CompanyTournamentView{
			Id:           "1",
			CompanyId:    "1",
			TournamentId: "a",
			Deposit:      "1"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_deposit", func(t *testing.T) {
		test_struct := models.CompanyTournamentView{
			Id:           "1",
			CompanyId:    "1",
			TournamentId: "1",
			Deposit:      "a"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})
}
