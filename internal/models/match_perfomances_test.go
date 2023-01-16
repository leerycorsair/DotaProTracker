package models_test

import (
	"local/internal/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMatchPerfomancesModels(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		test_struct := models.MatchPerfomanceView{
			Id:       "1",
			MatchId:  "1",
			PlayerId: "1",
			Team:     "t",
			Hero:     "pudge",
			Kills:    "1",
			Deaths:   "1",
			Assists:  "1",
			Networth: "1",
			GPM:      "1",
			XPM:      "1",
			DMG:      "1",
			Heal:     "1",
			BLD:      "1"}

		_, err := test_struct.FromViewConv()
		require.Nil(t, err)
	})

	t.Run("incorrect_id", func(t *testing.T) {
		test_struct := models.MatchPerfomanceView{
			Id:       "a",
			MatchId:  "1",
			PlayerId: "1",
			Team:     "t",
			Hero:     "pudge",
			Kills:    "1",
			Deaths:   "1",
			Assists:  "1",
			Networth: "1",
			GPM:      "1",
			XPM:      "1",
			DMG:      "1",
			Heal:     "1",
			BLD:      "1"}

		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_match_id", func(t *testing.T) {
		test_struct := models.MatchPerfomanceView{
			Id:       "1",
			MatchId:  "a",
			PlayerId: "1",
			Team:     "t",
			Hero:     "pudge",
			Kills:    "1",
			Deaths:   "1",
			Assists:  "1",
			Networth: "1",
			GPM:      "1",
			XPM:      "1",
			DMG:      "1",
			Heal:     "1",
			BLD:      "1"}

		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_player_id", func(t *testing.T) {
		test_struct := models.MatchPerfomanceView{
			Id:       "1",
			MatchId:  "1",
			PlayerId: "a",
			Team:     "t",
			Hero:     "pudge",
			Kills:    "1",
			Deaths:   "1",
			Assists:  "1",
			Networth: "1",
			GPM:      "1",
			XPM:      "1",
			DMG:      "1",
			Heal:     "1",
			BLD:      "1"}

		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})
	t.Run("incorrect_team", func(t *testing.T) {
		test_struct := models.MatchPerfomanceView{
			Id:       "1",
			MatchId:  "1",
			PlayerId: "1",
			Team:     "ta",
			Hero:     "pudge",
			Kills:    "1",
			Deaths:   "1",
			Assists:  "1",
			Networth: "1",
			GPM:      "1",
			XPM:      "1",
			DMG:      "1",
			Heal:     "1",
			BLD:      "1"}

		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_hero", func(t *testing.T) {
		test_struct := models.MatchPerfomanceView{
			Id:       "1",
			MatchId:  "1",
			PlayerId: "1",
			Team:     "t",
			Hero:     "",
			Kills:    "1",
			Deaths:   "1",
			Assists:  "1",
			Networth: "1",
			GPM:      "1",
			XPM:      "1",
			DMG:      "1",
			Heal:     "1",
			BLD:      "1"}

		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_kills", func(t *testing.T) {
		test_struct := models.MatchPerfomanceView{
			Id:       "1",
			MatchId:  "1",
			PlayerId: "1",
			Team:     "t",
			Hero:     "pudge",
			Kills:    "a",
			Deaths:   "1",
			Assists:  "1",
			Networth: "1",
			GPM:      "1",
			XPM:      "1",
			DMG:      "1",
			Heal:     "1",
			BLD:      "1"}

		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_deaths", func(t *testing.T) {
		test_struct := models.MatchPerfomanceView{
			Id:       "1",
			MatchId:  "1",
			PlayerId: "1",
			Team:     "t",
			Hero:     "pudge",
			Kills:    "1",
			Deaths:   "a",
			Assists:  "1",
			Networth: "1",
			GPM:      "1",
			XPM:      "1",
			DMG:      "1",
			Heal:     "1",
			BLD:      "1"}

		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_assists", func(t *testing.T) {
		test_struct := models.MatchPerfomanceView{
			Id:       "1",
			MatchId:  "1",
			PlayerId: "1",
			Team:     "t",
			Hero:     "pudge",
			Kills:    "1",
			Deaths:   "1",
			Assists:  "a",
			Networth: "1",
			GPM:      "1",
			XPM:      "1",
			DMG:      "1",
			Heal:     "1",
			BLD:      "1"}

		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_networth", func(t *testing.T) {
		test_struct := models.MatchPerfomanceView{
			Id:       "1",
			MatchId:  "1",
			PlayerId: "1",
			Team:     "t",
			Hero:     "pudge",
			Kills:    "1",
			Deaths:   "1",
			Assists:  "1",
			Networth: "a",
			GPM:      "1",
			XPM:      "1",
			DMG:      "1",
			Heal:     "1",
			BLD:      "1"}

		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_gpm", func(t *testing.T) {
		test_struct := models.MatchPerfomanceView{
			Id:       "1",
			MatchId:  "1",
			PlayerId: "1",
			Team:     "t",
			Hero:     "pudge",
			Kills:    "1",
			Deaths:   "1",
			Assists:  "1",
			Networth: "1",
			GPM:      "a",
			XPM:      "1",
			DMG:      "1",
			Heal:     "1",
			BLD:      "1"}

		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_xpm", func(t *testing.T) {
		test_struct := models.MatchPerfomanceView{
			Id:       "1",
			MatchId:  "1",
			PlayerId: "1",
			Team:     "t",
			Hero:     "pudge",
			Kills:    "1",
			Deaths:   "1",
			Assists:  "1",
			Networth: "1",
			GPM:      "1",
			XPM:      "a",
			DMG:      "1",
			Heal:     "1",
			BLD:      "1"}

		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_dmg", func(t *testing.T) {
		test_struct := models.MatchPerfomanceView{
			Id:       "1",
			MatchId:  "1",
			PlayerId: "1",
			Team:     "t",
			Hero:     "pudge",
			Kills:    "1",
			Deaths:   "1",
			Assists:  "1",
			Networth: "1",
			GPM:      "1",
			XPM:      "1",
			DMG:      "a",
			Heal:     "1",
			BLD:      "1"}

		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_heal", func(t *testing.T) {
		test_struct := models.MatchPerfomanceView{
			Id:       "1",
			MatchId:  "1",
			PlayerId: "1",
			Team:     "t",
			Hero:     "pudge",
			Kills:    "1",
			Deaths:   "1",
			Assists:  "1",
			Networth: "1",
			GPM:      "1",
			XPM:      "1",
			DMG:      "1",
			Heal:     "a",
			BLD:      "1"}

		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_bld", func(t *testing.T) {
		test_struct := models.MatchPerfomanceView{
			Id:       "1",
			MatchId:  "1",
			PlayerId: "1",
			Team:     "t",
			Hero:     "pudge",
			Kills:    "1",
			Deaths:   "1",
			Assists:  "1",
			Networth: "1",
			GPM:      "1",
			XPM:      "1",
			DMG:      "1",
			Heal:     "1",
			BLD:      "a"}

		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

}
