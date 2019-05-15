package main

import ("fmt")
type Product struct{
	Sku 		string
	Name		string 
	Quantity	int
}

var products = map[string]Product{
	"3RV442S75": Product{Sku: "3RV442S75", Name:"Computer Desk", Quantity: 10},
	"3RV446R75": Product{Sku: "3RV446R75", Name:"Chair", Quantity: 20},
	"3RWK38R70": Product{Sku: "3RWK38R70", Name:"Desk Lamp", Quantity: 30},
	"3RWK36S70": Product{Sku: "3RWK36S70", Name:"Book Case", Quantity: 40},
	"3RV646L16": Product{Sku: "3RV646L16", Name:"Computer Monitor", Quantity: 50},
	"3H7J36S73": Product{Sku: "3H7J36S73", Name:"Computer Keyboard", Quantity: 60},
} 

func ResetQty() {
	i := 1
	for _, product := range products { 
		product.Quantity = i * 10
		i++
	}
	fmt.Printf("RESET QTY: \n %v", products )
}

func AllProducts() []Product {
	values := make([]Product, len(products))
	i := 0
	for _, product := range products {
		values[i] = product
		i++
	}
	return values
}

func GetProduct(sku string) (Product, bool) {
	product, found := products[sku]
	return product, found
}

func UpdateProduct(sku string, product Product) bool {
	_, exists := products[sku]
	if exists {
		products[sku] = product
	}
	return exists
}


