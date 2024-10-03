package application_test

import (
	"testing"

	"github.com/diegodevtech/hexagonal-architecture/application"
	uuid "github.com/satori/go.uuid"
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

func TestProduct_Disable(t *testing.T){
	product := application.Product{}
	product.Name = "Produto 1"
	product.Status = application.ENABLED
	product.Price = 10.0

	err := product.Disable()
	require.Equal(t, "the price must be zero in order to disable a product" ,err.Error())

	product.Price = 0.0
	err = product.Disable()
	require.Nil(t, err)
	
}

func TestProduct_IsValid(t *testing.T){
	product := application.Product{}

	product.ID = uuid.NewV4().String()
	product.Name = "Produto 2"
	product.Status = application.DISABLED
	product.Price = 10.0

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"

	_, err = product.IsValid()
	require.Equal(t, "the status must be ENABLED or DISABLED", err.Error())

	product.Status = application.ENABLED
	product.Price = -1.0

	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal to zero", err.Error())
}

func TestProduct_GetID(t *testing.T){
	product := application.Product{}

	product.ID = "dddddddd-dddd-dddd-dddd-dddddddddddd"
	product.Name = "Produto 2"
	product.Status = application.DISABLED
	product.Price = 10.0

	id := product.GetID()
	require.Equal(t, "dddddddd-dddd-dddd-dddd-dddddddddddd", id)
}

func TestProduct_GetStatus(t *testing.T){
	product := application.Product{}

	product.ID = uuid.NewV4().String()
	product.Name = "Produto 2"
	product.Status = application.DISABLED
	product.Price = 10.0

	status := product.GetStatus()
	require.Equal(t, "disabled", status)

	product.Status = application.ENABLED

	status = product.GetStatus()
	require.Equal(t, "enabled", status)

	product.Status = ""
	product.IsValid()
	status = product.GetStatus()
	require.Equal(t, status, application.DISABLED)
}

func TestProduct_GetPrice(t *testing.T){
	product := application.Product{}

	product.ID = uuid.NewV4().String()
	product.Name = "Produto 2"
	product.Status = application.DISABLED
	product.Price = 10.0

	price := product.GetPrice()
	require.Equal(t, 10.0, price)
}

func TestProduct_GetName(t *testing.T){
	product := application.Product{}

	product.ID = uuid.NewV4().String()
	product.Name = "Produto 2"
	product.Status = application.DISABLED
	product.Price = 10.0

	name := product.GetName()
	require.Equal(t, "Produto 2", name)
}