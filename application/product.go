package application

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
	ENABLED = "enabled"
	DISABLED = "disabled"
)

type Product struct {
	ID string
	Name string
	Price float64
	Status string
}

func IsValid(p *Product) (bool, error){

}

func Enable(p *Product) (error){

}

func Disable(p *Product) (error){
	
}

func GetID(p *Product) string {

}

func GetName(p *Product) string {

}

func GetStatus(p *Product) string {

}

func GetPrice(p *Product) float64 {

}
