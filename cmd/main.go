package main

import (
	"fiber/pkg/setup"
	"fmt"
)

func main() {
	err := setup.Setup()
	if err != nil {
		fmt.Println(err)
	}

}
