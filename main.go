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

	order, err := libmonpos.FindMonitorOrder(config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", config)
	fmt.Printf("%v\n", order)
}
