package main

import "fmt"

type Ecommerce struct {
	users []User
}

type User struct {
	firstName string
	lastName  string
	email     string
	Products  []Product
}

type Product struct {
	name   string
	price  float64
	amount int
}

func main() {
	ecommerce := Ecommerce{}
	user := User{firstName: "John", lastName: "Doe", email: "johndoe@gmail.com"}

	ecommerce.AddUser(&user)
	firstProduct := newProduct("Product 1", 10.0, 1)
	secondProduct := newProduct("Product 2", 20.0, 2)

	ecommerce.AddProductToUser(&user, &firstProduct)
	ecommerce.AddProductToUser(&user, &secondProduct)

	ecommerce.PrintAllProducts(&user)
	ecommerce.DeleteAllProducts(&user)
	ecommerce.PrintAllProducts(&user)
}

func (e *Ecommerce) AddUser(u *User) {
	e.users = append(e.users, *u)
}

func newProduct(name string, price float64, amount int) Product {
	return Product{name: name, price: price, amount: amount}
}

func (e *Ecommerce) AddProductToUser(u *User, p *Product) {
	for i, user := range e.users {
		if user.email == u.email {
			e.users[i].Products = append(e.users[i].Products, *p)
			fmt.Printf("Product %s added to user %s\n", p.name, u.email)
		}
	}
}

func (e *Ecommerce) DeleteAllProducts(u *User) {
	for i, user := range e.users {
		if user.email == u.email {
			e.users[i].Products = nil
			fmt.Printf("All products deleted from user %s\n", u.email)
		}
	}
}

func (e *Ecommerce) PrintAllProducts(u *User) {
	for _, user := range e.users {
		if user.email == u.email {
			if len(user.Products) == 0 {
				fmt.Printf("User %s has no products\n", u.email)
				return
			}

			for _, product := range user.Products {
				fmt.Printf("Product %s, price %f, amount %d\n", product.name, product.price, product.amount)
			}
		}
	}
}
