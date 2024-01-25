package services

import repositores "go-redis/repositories/product"

type catalogServiceDb struct {
	productRepo repositores.ProductRepository
}

func NewCatalogServiceDB(productRepo repositores.ProductRepository) CatalogService{
	return catalogServiceDb{productRepo}
}

func (s catalogServiceDb)GetProducts() (products []Product,err error){

	productsDB, err := s.productRepo.GetProducts()
	if err != nil{
		return nil,err
	}

	for _,p := range productsDB {
		products = append(products, Product{
			ID: p.ID,
			Name: p.Name,
			Quantity: p.Quantity,
		})
	}


	return products,nil
}