package main

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ReleaseTestSuite struct {
	suite.Suite
}

func (r *ReleaseTestSuite) TestRelease() {
	// -- Given
	//
	type test struct {
		Given           string
		ExpectedVersion string
		ExpectedEdition Edition
	}
	tests := []test{
		{Given: "Minecraft - 1.2.11 (Bedrock)", ExpectedEdition: EditionBedrock, ExpectedVersion: "1.2.11"},
		{Given: "Minecraft: Java Edition - 1.15.1", ExpectedEdition: EditionJava, ExpectedVersion: "1.15.1"},
		{Given: "Minecraft - Caves & Cliffs: Part II - 1.18.0 (Bedrock)", ExpectedEdition: EditionBedrock, ExpectedVersion: "1.18.0"},
		{Given: "Minecraft - Caves & Cliffs: Part 1 - 1.17 (Java)", ExpectedEdition: EditionJava, ExpectedVersion: "1.17"},
		{Given: "MCPE/WIN 10/Gear VR - 1.0.2", ExpectedVersion: "1.0.2"},
	}

	// -- When
	//
	for i, v := range tests {
		given := Release(v.Given)
		r.Equal(given.Version(), v.ExpectedVersion, "test: %d, given: %s", i, v.Given)
		r.Equal(given.Edition(), v.ExpectedEdition, "test: %d, given: %s", i, v.Given)
	}
}

func TestReleaseTestSuite(t *testing.T) {
	suite.Run(t, new(ReleaseTestSuite))
}
