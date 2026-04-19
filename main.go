package main

import (
	"fmt"
)

type Category struct {
	Name string
	Tax  float64
}

type Product struct {
	ID       int
	Name     string
	Category Category
	Price    float64
}

func (p Product) FinalPrice() float64 {
	return p.Price + p.Price*p.Category.Tax/100
}

func FindMostExpensiveByCategory(products []Product, categoryName string) (Product, bool) {
	var result Product
	found := false

	for _, product := range products {
		if product.Category.Name == categoryName {
			if !found || product.FinalPrice() > result.FinalPrice() {
				result = product
				found = true
			}
		}
	}

	return result, found
}

func PrintProducts(products []Product) {
	for _, p := range products {
		fmt.Printf("ID: %d | Name: %s | Category: %s | Base price: %.2f | Final price: %.2f\n",
			p.ID, p.Name, p.Category.Name, p.Price, p.FinalPrice())
	}
}

func main() {
	defer fmt.Println("\nProgram finished.")

	laptops := Category{Name: "Laptops", Tax: 20}
	smartphones := Category{Name: "Smartphones", Tax: 15}
	accessories := Category{Name: "Accessories", Tax: 10}

	categories := []Category{laptops, smartphones, accessories}

	products := []Product{
		{ID: 1, Name: "MacBook Air M3", Category: laptops, Price: 5200},
		{ID: 2, Name: "Lenovo Legion 5", Category: laptops, Price: 6100},
		{ID: 3, Name: "iPhone 15", Category: smartphones, Price: 4300},
		{ID: 4, Name: "Samsung Galaxy S24", Category: smartphones, Price: 3900},
		{ID: 5, Name: "Logitech MX Master 3S", Category: accessories, Price: 450},
		{ID: 6, Name: "Sony WH-1000XM5", Category: accessories, Price: 1700},
	}

	fmt.Println("List of all products:")
	PrintProducts(products)

	fmt.Println("\nAvailable categories:")
	for i, category := range categories {
		fmt.Printf("%d - %s\n", i+1, category.Name)
	}

	var choice int
	fmt.Print("\nEnter category number: ")
	fmt.Scan(&choice)

	if choice < 1 || choice > len(categories) {
		fmt.Println("Invalid category number.")
		return
	}

	selectedCategory := categories[choice-1].Name

	product, found := FindMostExpensiveByCategory(products, selectedCategory)
	if found {
		fmt.Printf("\nThe most expensive product in category \"%s\":\n", selectedCategory)
		fmt.Printf("ID: %d | Name: %s | Final price: %.2f\n",
			product.ID, product.Name, product.FinalPrice())
	} else {
		fmt.Println("No products found in this category.")
	}
}
