package main

import (
	"flag"
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/hostfactor/minecrafter"
	"github.com/hostfactor/minecrafter/edition"
	"os"
)

func main() {
	edBuilder := minecrafter.New([]string{os.Getenv("GITHUB_REGISTRY")})
	all := flag.Bool("all", false, "Build every Minecraft version")
	flag.Parse()
	args := flag.Args()

	con, _ := semver.NewConstraint(">= 1.18")
	opts := []minecrafter.BuildOpt{
		minecrafter.WithSemverConstraint(con),
	}
	var err error
	if all != nil && *all {
		fmt.Println("Building every Minecraft version.")
		opts = []minecrafter.BuildOpt{}
	}
	if len(args) > 1 {
		err = edBuilder.BuildRelease(new(edition.Java), os.Args[1], opts...)
	} else {
		err = edBuilder.BuildEdition(new(edition.Java), opts...)
	}
	if err != nil {
		panic(err.Error())
	}
}
