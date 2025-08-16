package main

import (
	"fmt"

	"github.com/almondheil/libmonpos"
)

func main() {
	err, config := libmonpos.ReadConfig("./example.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", config)
}
