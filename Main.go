package main

import (
	"fmt"

	"tudai2021.com/model"
)

func main() {
	Result, err := model.Parser("TX03ABC")
	fmt.Println(Result, err)
}
