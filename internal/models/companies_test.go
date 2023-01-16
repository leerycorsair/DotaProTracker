package models_test

import (
	"local/internal/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompaniesModelsUnit(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		test_struct := models.CompanyView{
			Id:       "1",
			Name:     "name",
			Country:  "country",
			Website:  "website.com",
			Revenue:  "1",
			Industry: "industry"}
		_, err := test_struct.FromViewConv()
		require.Nil(t, err)
	})

	t.Run("incorrect_id", func(t *testing.T) {
		test_struct := models.CompanyView{
			Id:       "a",
			Name:     "name",
			Country:  "country",
			Website:  "website.com",
			Revenue:  "1",
			Industry: "industry"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_name", func(t *testing.T) {
		test_struct := models.CompanyView{
			Id:       "1",
			Name:     "",
			Country:  "country",
			Website:  "website.com",
			Revenue:  "1",
			Industry: "industry"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_country", func(t *testing.T) {
		test_struct := models.CompanyView{
			Id:       "1",
			Name:     "name",
			Country:  "",
			Website:  "website.com",
			Revenue:  "1",
			Industry: "industry"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_website", func(t *testing.T) {
		test_struct := models.CompanyView{
			Id:       "1",
			Name:     "name",
			Country:  "country",
			Website:  "",
			Revenue:  "1",
			Industry: "industry"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_revenue", func(t *testing.T) {
		test_struct := models.CompanyView{
			Id:       "1",
			Name:     "name",
			Country:  "country",
			Website:  "website.com",
			Revenue:  "a",
			Industry: "industry"}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

	t.Run("incorrect_industry", func(t *testing.T) {
		test_struct := models.CompanyView{
			Id:       "1",
			Name:     "name",
			Country:  "country",
			Website:  "website.com",
			Revenue:  "1",
			Industry: ""}
		_, err := test_struct.FromViewConv()
		require.NotNil(t, err)
	})

}
