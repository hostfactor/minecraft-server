package main

import (
	"regexp"
	"strings"
)

type Release string

type Edition int

var semverRegex = regexp.MustCompile("[0-9]+(\\.[0-9]+)+")

const (
	EditionUnknown Edition = iota
	EditionJava
	EditionBedrock
)

func (r Release) Edition() Edition {
	lower := strings.ToLower(string(r))
	if strings.Contains(lower, "java") {
		return EditionJava
	} else if strings.Contains(lower, "bedrock") {
		return EditionBedrock
	}
	return EditionUnknown
}

func (r Release) Version() string {
	return semverRegex.FindString(string(r))
}
