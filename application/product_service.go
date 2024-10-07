package application

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// Get(id string) (ProductInterface, error)
// Create(name string, price float64) (ProductInterface, error)
// Enable(product ProductInterface) (ProductInterface, error)
// Disable(product ProductInterface) (ProductInterface, error)