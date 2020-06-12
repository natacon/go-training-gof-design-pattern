package factory_method

type Product interface {
	Use()
}

type IFactory interface {
	CreateProduct(owner string) Product
	RegisterProduct(product Product)
	Create(owner string) Product
}

type Factory struct {
	factory IFactory
}

func (f *Factory) Create(owner string) Product {
	product := f.factory.CreateProduct(owner)
	f.factory.RegisterProduct(product)
	return product
}
