package main

import (
	"fmt"
)

type Product struct {
	name   string
	price  float64
	amount int
}

type Services struct {
	name    string
	price   float64
	minutes int
}

type Maintenance struct {
	name  string
	price float64
}

func main() {
	firstProduct := createProduct("Product 1", 10.0, 2)
	secondProduct := createProduct("Product 2", 20.0, 3)
	thirdProduct := createProduct("Product 3", 30.0, 4)

	firstService := createService("Service 1", 10.0, 60)
	secondService := createService("Service 2", 20.0, 20)

	firstMaintenance := createMaintenance("Maintenance 1", 10.0)

	products := []Product{firstProduct, secondProduct, thirdProduct}
	services := []Services{firstService, secondService}
	maintenance := []Maintenance{firstMaintenance}

	results := make(chan float64)

	go func() {
		totalProducts := sumProducts(products)
		fmt.Printf("Total products: %.2f\n", totalProducts)
		results <- totalProducts
	}()

	go func() {
		totalServices := sumServices(services)
		fmt.Printf("Total services: %.2f\n", totalServices)
		results <- totalServices
	}()

	go func() {
		totalMaintenance := sumMaintenance(maintenance)
		fmt.Printf("Total maintenance: %.2f\n", totalMaintenance)
		results <- totalMaintenance
	}()

	totalSum := 0.0
	for i := 0; i < 3; i++ {
		totalSum += <-results
	}

	fmt.Printf("Total sum: %.2f\n", totalSum)
}

func sumProducts(p []Product) float64 {
	var sum float64
	for _, v := range p {
		sum += v.price * float64(v.amount)
	}
	return sum
}

func sumServices(s []Services) float64 {
	var sum float64
	for _, v := range s {
		hours := float64(v.minutes) / 60.0
		if hours < 0.5 {
			hours = 0.5
		}
		sum += v.price * hours
	}
	return sum
}

func sumMaintenance(m []Maintenance) float64 {
	var sum float64
	for _, v := range m {
		sum += v.price
	}
	return sum
}

func createProduct(name string, price float64, amount int) Product {
	return Product{name, price, amount}
}

func createService(name string, price float64, minutes int) Services {
	return Services{name, price, minutes}
}

func createMaintenance(name string, price float64) Maintenance {
	return Maintenance{name, price}
}
