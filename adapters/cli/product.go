package cli

import (
	"fmt"

	"github.com/diegodevtech/hexagonal-architecture/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with name %s has been created with price %f and status %s", 
		product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		productResult, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been enabled.", productResult.GetName())

	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		productResult, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been disabled.", productResult.GetName())

	case "all":
		productArr, err := service.GetAll()
		if err != nil {
			return result, err
		}
		for i := 0 ; i < len(productArr) ; i++ {
			if i >= 1 {
				result += fmt.Sprintf("\n==========\nProduct ID: %s\nName: %s\nPrice: %.2f\nStatus: %s", productArr[i].GetID(), productArr[i].GetName(), productArr[i].GetPrice(), productArr[i].GetStatus())
			} else {
				result += fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %.2f\nStatus: %s", productArr[i].GetID(), productArr[i].GetName(), productArr[i].GetPrice(), productArr[i].GetStatus())
			}
		}
		
	default:
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	}
	return result, nil
}