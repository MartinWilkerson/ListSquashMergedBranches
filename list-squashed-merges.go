package main

import (
	"flag"
	"fmt"
)

func main() {
	projectURLPtr := flag.String("projecturl", "", "URL of the project to check")
	apiKeyPtr := flag.String("apikey", "", "API key to use")

	flag.Parse()

	fmt.Println("URL: ", *projectURLPtr)
	fmt.Println("Key: ", *apiKeyPtr)
}
