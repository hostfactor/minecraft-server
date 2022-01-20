package main

import (
	"github.com/hostfactor/minecrafter"
	"github.com/hostfactor/minecrafter/edition"
	"os"
)

func main() {
	edBuilder := minecrafter.New([]string{os.Getenv("GITHUB_REGISTRY")})

	err := edBuilder.BuildEdition(new(edition.JavaEdition))
	if err != nil {
		panic(err.Error())
	}
}
