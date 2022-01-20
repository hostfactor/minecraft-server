package main

import (
	"github.com/hostfactor/minecrafter"
	"github.com/hostfactor/minecrafter/edition"
	"os"
)

func main() {
	edBuilder := minecrafter.New([]string{os.Getenv("GITHUB_REGISTRY")})

	var err error
	if len(os.Args) > 1 {
		err = edBuilder.BuildRelease(new(edition.JavaEdition), os.Args[1])
	} else {
		err = edBuilder.BuildEdition(new(edition.JavaEdition))
	}
	if err != nil {
		panic(err.Error())
	}
}
