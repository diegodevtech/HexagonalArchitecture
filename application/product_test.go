package application_test

import (
	"testing"
	"github.com/diegodevtech/hexagonal-architecture/application"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T){
	product := application.Product{}
	product.Name = "Produto 1"
	product.Status = application.DISABLED
	product.Price = 10.0

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0.0
	err = product.Enable()
	require.Equal(t, "the price must be grater than 0 to enable a product", err.Error())
	
}