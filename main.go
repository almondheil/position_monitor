package main

import (
	"fmt"

	"github.com/almondheil/libmonpos"
)

func main() {
	config, err := libmonpos.LoadConfig("./example.yaml")
	if err != nil {
		panic(err)
	}

	graph, err := libmonpos.LoadGraph(config)
	if err != nil {
		panic(err)
	}

	positions, err := libmonpos.GeneratePositions(config, graph)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", config)
	fmt.Printf("%v\n", positions)
}
