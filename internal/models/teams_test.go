package models_test

import (
	"local/internal/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTeamsModelsUnit(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		test_struct := models.TeamView{
			Id:            "1",
			Name:          "name",
			CreatedAt:     "2000-01-01",
			Email:         "email",
			TotalEarnings: "1",
			Region:        "region",
			Tier:          "1"}
		_, err := test_struct.FromViewConv()
		require.Nil(t, err)
	})

	t.Run("incorrect_id", func(t *testing.T) {
		test_struct := models.TeamView{
			Id:            "a",
			Name:          "name",
			CreatedAt:     "2000-01-01",
			Email:         "email",
			TotalEarnings: "1",
			Region:        "region",
			Tier:          "1"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_name", func(t *testing.T) {
		test_struct := models.TeamView{
			Id:            "1",
			Name:          "",
			CreatedAt:     "2000-01-01",
			Email:         "email",
			TotalEarnings: "1",
			Region:        "region",
			Tier:          "1"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})
	t.Run("incorrect_email", func(t *testing.T) {
		test_struct := models.TeamView{
			Id:            "1",
			Name:          "name",
			CreatedAt:     "2000-01-01",
			Email:         "",
			TotalEarnings: "1",
			Region:        "region",
			Tier:          "1"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})
	t.Run("incorrect_total_earnings", func(t *testing.T) {
		test_struct := models.TeamView{
			Id:            "1",
			Name:          "name",
			CreatedAt:     "2000-01-01",
			Email:         "email",
			TotalEarnings: "a",
			Region:        "region",
			Tier:          "1"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})
	t.Run("incorrect_region", func(t *testing.T) {
		test_struct := models.TeamView{
			Id:            "1",
			Name:          "name",
			CreatedAt:     "2000-01-01",
			Email:         "email",
			TotalEarnings: "1",
			Region:        "",
			Tier:          "1"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})
	t.Run("incorrect_tier", func(t *testing.T) {
		test_struct := models.TeamView{
			Id:            "1",
			Name:          "name",
			CreatedAt:     "2000-01-01",
			Email:         "email",
			TotalEarnings: "1",
			Region:        "region",
			Tier:          "a"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})
}
