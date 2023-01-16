package models_test

import (
	"local/internal/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompaniesTeamsModelsUnit(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		test_struct := models.CompanyTeamView{
			Id:           "1",
			CompanyId:    "1",
			TeamId:       "1",
			ContractDate: "2000-01-01",
			ContractTime: "1"}
		_, err := test_struct.FromViewConv()
		require.Nil(t, err)
	})

	t.Run("incorrect_id", func(t *testing.T) {
		test_struct := models.CompanyTeamView{
			Id:           "fsdfsd",
			CompanyId:    "1",
			TeamId:       "1",
			ContractDate: "2000-01-01",
			ContractTime: "1"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_company_id", func(t *testing.T) {
		test_struct := models.CompanyTeamView{
			Id:           "1",
			CompanyId:    "sfsd",
			TeamId:       "1",
			ContractDate: "2000-01-01",
			ContractTime: "1"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_team_id", func(t *testing.T) {
		test_struct := models.CompanyTeamView{
			Id:           "1",
			CompanyId:    "1",
			TeamId:       "fskfsfkd",
			ContractDate: "2000-01-01",
			ContractTime: "1"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_contract_time", func(t *testing.T) {
		test_struct := models.CompanyTeamView{
			Id:           "1",
			CompanyId:    "1",
			TeamId:       "1",
			ContractDate: "2000-01-01",
			ContractTime: "fsdfsdkfms"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})
}
