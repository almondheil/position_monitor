package main

import (
	"fmt"

	"github.com/almondheil/monitor_position"
)

func main() {
	err, config := monitor_position.ReadConfig("./example.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", config)
}
