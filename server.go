package main

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

var errorMessage =  "Invalid product ID"
var errorMessage2 =  "Product not found"
var path = "/products/:id"
var products = []Product{
	{ID: 1, Name: "Product 1", Description: "Lorem ipsum", Price: 2.50},
	{ID: 2, Name: "Product 2", Description: "Lorem ipsum", Price: 5.00},
}

func getProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest,errorMessage)
	}

	for _, product := range products {
		if product.ID == id {
			return c.JSON(http.StatusOK, product)
		}
	}

	return c.String(http.StatusNotFound, errorMessage)
}

func getProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func createProduct(c echo.Context) error {
	product := new(Product)
	if err := c.Bind(product); err != nil {
		return err
	}

	product.ID = len(products) + 1
	products = append(products, *product)

	return c.JSON(http.StatusCreated, product)
}

func updateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, errorMessage)
	}

	for i, product := range products {
		if product.ID == id {
			updatedProduct := new(Product)
			if err := c.Bind(updatedProduct); err != nil {
				return err
			}

			updatedProduct.ID = id
			products[i] = *updatedProduct

			return c.JSON(http.StatusOK, updatedProduct)
		}
	}

	return c.String(http.StatusNotFound, errorMessage2)
}

func deleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, errorMessage)
	}

	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}

	return c.String(http.StatusNotFound, errorMessage2)
}

func main() {
	e := echo.New()

	e.GET("/products", getProducts)
	e.GET(path, getProduct)
	e.POST("/products", createProduct)
	e.PUT(path, updateProduct)
	e.DELETE(path, deleteProduct)

	e.Logger.Fatal(e.Start(":8080"))
}