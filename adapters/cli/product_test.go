package cli_test

import (
	"fmt"
	"testing"

	"github.com/diegodevtech/hexagonal-architecture/adapters/cli"
	"github.com/diegodevtech/hexagonal-architecture/application"
	mock_application "github.com/diegodevtech/hexagonal-architecture/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    productName := "Product Test"
    productPrice := 25.99
    productStatus := "enabled"
    productId := "abc"

    productName2 := "X"
    productPrice2 := 10.50
    productStatus2 := "disabled"
    productId2 := "x"

    // Mock product 1
    productMock := mock_application.NewMockProductInterface(ctrl)
    productMock.EXPECT().GetID().Return(productId).AnyTimes()
    productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
    productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
    productMock.EXPECT().GetName().Return(productName).AnyTimes()

    // Mock product 2
    productMock2 := mock_application.NewMockProductInterface(ctrl)
    productMock2.EXPECT().GetID().Return(productId2).AnyTimes()
    productMock2.EXPECT().GetStatus().Return(productStatus2).AnyTimes()
    productMock2.EXPECT().GetPrice().Return(productPrice2).AnyTimes()
    productMock2.EXPECT().GetName().Return(productName2).AnyTimes()

    // Mock service behavior
    service := mock_application.NewMockProductServiceInterface(ctrl)
    service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
    service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
    service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
    service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()
    service.EXPECT().GetAll().Return([]application.ProductInterface{productMock, productMock2}, nil).AnyTimes()

    // Test case for "create"
    resultExpected := fmt.Sprintf("Product ID %s with name %s has been created with price %f and status %s", 
        productId, productName, productPrice, productStatus)
    result, err := cli.Run(service, "create", "", productName, productPrice)
    require.Nil(t, err)
    require.Equal(t, resultExpected, result)

    // Test case for "enable"
    resultExpected = fmt.Sprintf("Product %s has been enabled.", productName)
    result, err = cli.Run(service, "enable", productId, "", 0.0)
    require.Nil(t, err)
    require.Equal(t, resultExpected, result)

    // Test case for "disable"
    resultExpected = fmt.Sprintf("Product %s has been disabled.", productName)
    result, err = cli.Run(service, "disable", productId, "", 0.0)
    require.Nil(t, err)
    require.Equal(t, resultExpected, result)

    // Test case for "get"
    resultExpected = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s", productId, productName, productPrice, productStatus)
    result, err = cli.Run(service, "get", productId, "", 0.0)
    require.Nil(t, err)
    require.Equal(t, resultExpected, result)

    // Test case for "all"
    resultExpected = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %.2f\nStatus: %s\n==========\nProduct ID: %s\nName: %s\nPrice: %.2f\nStatus: %s", 
        productId, productName, productPrice, productStatus, productId2, productName2, productPrice2, productStatus2)
    result, err = cli.Run(service, "all", "", "", 0.0)
    require.Nil(t, err)
    require.Equal(t, resultExpected, result)
}