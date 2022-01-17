package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"os"
	"os/exec"
	"path"
	"strings"
)

type JDKTag struct {
	Tag         string
	DisplayName string
	IsDefault   IsDefaultFunc
	Skip        SkipFunc
}

type IsDefaultFunc func(version, tag string) bool
type SkipFunc func(version, tag string) bool

func Skip(version, tag string) bool {
	return strings.HasPrefix(version, "1.18") && tag != "17-alpine"
}

func IsDefault(version, tag string) bool {
	return strings.HasPrefix(version, "1.18") && tag == "17-alpine"
}

var OpenJDKTags = []JDKTag{
	{Tag: "17-alpine", DisplayName: "java-17", Skip: Skip, IsDefault: IsDefault},
	{Tag: "15-alpine", DisplayName: "java-15", Skip: Skip, IsDefault: IsDefault},
	{Tag: "11-jre-slim", DisplayName: "java-11", Skip: Skip, IsDefault: IsDefault},
	{Tag: "8-jre-slim", DisplayName: "java-8", Skip: Skip, IsDefault: IsDefault},
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
		if path.Ext(link) != ".jar" {
			return
		}

		toPush := make([]string, 0, len(OpenJDKTags))
		for _, v := range OpenJDKTags {
			if v.Skip(version, v.Tag) {
				fmt.Println("Skipping JDK", v.Tag, "for version", version, ": Incompatible.")
				continue
			}
			fmt.Println("Building image for version", version, "jar path", link, "jdk tag", v)
			tag := fmt.Sprintf("%s:%s-%s", os.Getenv("GITHUB_REGISTRY"), version, v.DisplayName)
			if v.IsDefault(version, v.Tag) {
				tag = fmt.Sprintf("%s:%s", os.Getenv("GITHUB_REGISTRY"), version)
			}
			cmd := exec.Command("docker", "build", "-t",
				tag,
				"--build-arg", fmt.Sprintf("MINECRAFT_JAR_PATH=%s", link),
				"--build-arg", fmt.Sprintf("OPENJDK_TAG=%s", v.Tag),
				".",
			)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Failed to run", err.Error())
			}
			toPush = append(toPush, tag)
		}

		if len(toPush) == 0 {
			fmt.Println("Nothing to push. Skipping.")
			return
		}

		for _, v := range toPush {
			fmt.Println("Pushing", v)
			cmd := exec.Command("docker", "push", v)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Failed to push", err.Error())
			}
		}
	})

	err := c.Visit("https://minecraft.fandom.com/wiki/Java_Edition_version_history")
	if err != nil {
		panic(err.Error())
	}
}
