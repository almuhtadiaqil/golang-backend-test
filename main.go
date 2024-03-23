package main

import (
	"backend-test/config"
	"backend-test/src"
)

func main() {
	err := config.Init()
	if err != nil {
		panic(err)
	}

	dep := src.Dependencies()

	r := src.SetupRouter(dep)

	r.Run()
}
