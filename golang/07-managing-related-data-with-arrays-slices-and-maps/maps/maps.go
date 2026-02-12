package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	// Maps:
	// A map in Go is a built‑in data structure that stores unordered key–value pairs,
	// where each key is unique and is used to quickly look up its associated value.

	websites := map[string]string{
		"Amazon Web Services": "https://www.aws.com",
		"Google":              "https://www.google.com",
	}

	fmt.Println(websites)

	fmt.Println(websites["Amazon Web Services"])
	fmt.Println(websites["Google"])

	websites["LinkedIn"] = "https://www.linkedin.com"

	fmt.Println(websites)

	delete(websites, "Amazon Web Services")
	fmt.Println(websites)
}
