package application

import (
	"errors"
	"github.com/asaskevich/govalidator"
)

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	ENABLED  = "enabled"
	DISABLED = "disabled"
)

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
	Status string  `valid:"required"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}
	if p.Status != DISABLED && p.Status != ENABLED {
		return false, errors.New("the status must be ENABLED or DISABLED")
	}
	if p.Price < 0 {
		return false, errors.New("the price must be greater or equal to zero")
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0.0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("the price must be grater than 0 to enable a product")
}

func (p *Product) Disable() error {
	if p.Price == 0.0 {
		p.Status = DISABLED
		return nil
	}
	return errors.New("the price must be zero in order to disable a product")
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
