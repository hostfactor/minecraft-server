package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"os"
	"os/exec"
	"path"
	"strings"
)

var OpenJDKTags = []string{
	"17-alpine",
	"15-alpine",
	"11-jre-slim",
	"8-jre-slim",
}

func main() {
	c := colly.NewCollector(colly.UserAgent("Mozilla/5.0"))

	version := ""
	link := ""
	c.OnHTML("table.wikitable td a[href]", func(element *colly.HTMLElement) {
		if !semverRegex.MatchString(element.Text) {
			return
		}

		version = element.Text
		err := element.Request.Visit(element.Attr("href"))
		if err != nil {
			panic(err.Error())
		}
	})

	c.OnHTML(`a[href].external.text`, func(element *colly.HTMLElement) {
		if element.Text != "Server" {
			return
		}
		link = element.Attr("href")
		if path.Ext(link) == ".jar" {
			for _, v := range OpenJDKTags {
				if strings.HasPrefix(version, "1.18") && v != "17-alpine" {
					fmt.Println("Skipping JDK", v, "for version", version, ": Incompatible.")
				}
				fmt.Println("Building image for version", version, "jar path", link, "jdk tag")
				cmd := exec.Command("docker", "build", "-t",
					fmt.Sprintf("%s:%s", os.Getenv("GITHUB_REGISTRY"), version),
					"--build-arg", fmt.Sprintf("MINECRAFT_JAR_PATH=%s", link),
					"--build-arg", fmt.Sprintf("OPENJDK_TAG=%s", v),
				)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					fmt.Println("Failed to run", err.Error())
				}
			}
		}
	})

	err := c.Visit("https://minecraft.fandom.com/wiki/Java_Edition_version_history")
	if err != nil {
		panic(err.Error())
	}
}
