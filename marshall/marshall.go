package marshall

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
)

type (
	Products struct {
		Items []Product
	}
	Product struct {
		ID            int
		Name          string
		Description   string
		FeaturedPhoto string
		GalleryPhotos []string
		RegularPrice  float32
		SalePrice     float32
		Vendor        string
		Quantity      int
		Attributes    []Attribute
		Dimensions    []Dimension
	}
	Attribute struct {
		Name  string
		Value any
	}
	Dimension struct {
		Name  string
		Value float32
	}
)

// Generate a random string of a given length
func randomString(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// List of possible values for FeaturedPhoto and GalleryPhotos
var photoURLs = []string{
	"https://example.com/photo1.jpg",
	"https://example.com/photo2.jpg",
	"https://example.com/photo3.jpg",
	"https://example.com/photo4.jpg",
	"https://example.com/photo5.jpg",
}

// List of possible values for Vendor
var vendors = []string{
	"VendorA",
	"VendorB",
	"VendorC",
	"VendorD",
	"VendorE",
}

// List of possible values for Attributes
var attributes = []Attribute{
	{Name: "Color", Value: "Red"},
	{Name: "Color", Value: "Blue"},
	{Name: "Size", Value: "Small"},
	{Name: "Size", Value: "Large"},
	{Name: "Material", Value: "Wood"},
	{Name: "Material", Value: "Metal"},
}

// List of possible values for Dimensions
var dimensions = []Dimension{
	{Name: "Height", Value: 10.5},
	{Name: "Height", Value: 20.0},
	{Name: "Width", Value: 5.0},
	{Name: "Width", Value: 15.0},
	{Name: "Depth", Value: 3.0},
	{Name: "Depth", Value: 8.0},
}

// Shuffle and select a random subset from a slice
func randomSubset[T any](slice []T, count int) []T {
	if count > len(slice) {
		count = len(slice)
	}
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
	return slice[:count]
}

// Generate a random Product
func randomProduct() Product {
	galleryPhotos := randomSubset(photoURLs, 2)
	attributes := randomSubset(attributes, 2)
	dimensions := randomSubset(dimensions, 2)

	return Product{
		ID:            rand.Intn(1000),
		Name:          randomString(10),
		Description:   randomString(50),
		FeaturedPhoto: photoURLs[rand.Intn(len(photoURLs))],
		GalleryPhotos: galleryPhotos,
		RegularPrice:  rand.Float32() * 100,
		SalePrice:     rand.Float32() * 100,
		Vendor:        vendors[rand.Intn(len(vendors))],
		Quantity:      rand.Intn(100),
		Attributes:    attributes,
		Dimensions:    dimensions,
	}
}

// Generate a random Products struct
func randomProducts(count int) Products {
	products := Products{
		Items: make([]Product, count),
	}
	for i := range products.Items {
		products.Items[i] = randomProduct()
	}
	return products
}

func TryMarshall() {
	products := randomProducts(5)
	body, err := json.MarshalIndent(products, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
