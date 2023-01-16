package models_test

import (
	"local/internal/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTournamentsModelsUnit(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		test_struct := models.TournamentView{
			Id:        "1",
			Name:      "name",
			Tier:      "1",
			PrizePool: "1",
			DateStart: "2000-01-01",
			Duration:  "1",
			DPCPoints: "1",
			Location:  "location"}
		_, err := test_struct.FromViewConv()
		require.Nil(t, err)
	})

	t.Run("incorrect_id", func(t *testing.T) {
		test_struct := models.TournamentView{
			Id:        "a",
			Name:      "name",
			Tier:      "1",
			PrizePool: "1",
			DateStart: "2000-01-01",
			Duration:  "1",
			DPCPoints: "1",
			Location:  "location"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_name", func(t *testing.T) {
		test_struct := models.TournamentView{
			Id:        "1",
			Name:      "",
			Tier:      "1",
			PrizePool: "1",
			DateStart: "2000-01-01",
			Duration:  "1",
			DPCPoints: "1",
			Location:  "location"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_tier", func(t *testing.T) {
		test_struct := models.TournamentView{
			Id:        "1",
			Name:      "name",
			Tier:      "",
			PrizePool: "1",
			DateStart: "2000-01-01",
			Duration:  "1",
			DPCPoints: "1",
			Location:  "location"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_prize_pool", func(t *testing.T) {
		test_struct := models.TournamentView{
			Id:        "1",
			Name:      "name",
			Tier:      "1",
			PrizePool: "a",
			DateStart: "2000-01-01",
			Duration:  "1",
			DPCPoints: "1",
			Location:  "location"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_duration", func(t *testing.T) {
		test_struct := models.TournamentView{
			Id:        "1",
			Name:      "name",
			Tier:      "1",
			PrizePool: "1",
			DateStart: "2000-01-01",
			Duration:  "a",
			DPCPoints: "1",
			Location:  "location"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_dpc_points", func(t *testing.T) {
		test_struct := models.TournamentView{
			Id:        "1",
			Name:      "name",
			Tier:      "1",
			PrizePool: "1",
			DateStart: "2000-01-01",
			Duration:  "1",
			DPCPoints: "a",
			Location:  "location"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_location", func(t *testing.T) {
		test_struct := models.TournamentView{
			Id:        "1",
			Name:      "name",
			Tier:      "1",
			PrizePool: "1",
			DateStart: "2000-01-01",
			Duration:  "1",
			DPCPoints: "1",
			Location:  ""}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

}
