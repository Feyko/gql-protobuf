package main

import (
	"flag"
	"fmt"
	"gql-proto-gen/gen"
	"log"
	"os"
)

func main() {
	filepath := flag.String("f", "", "GraphQL file to parse")
	flag.Parse()

	data, err := os.ReadFile(*filepath)
	if err != nil {
		log.Fatalf("Could not read the GraphQL file: %v", err)
	}

	doc, err := gen.Parse(string(data))
	if err != nil {
		log.Fatalf("Could not parse the GraphQL schema: %v", err)
	}

	fmt.Println(doc)
}
