package mock_test

import (
	"testing"

	"github.com/seppone/goenv"
	"github.com/seppone/goenv/mock"
)

var versions = []goenv.Version{
	{
		Tag:            "1.8",
		Implementation: "go",
	},
	{
		Tag:            "1.8.1",
		Implementation: "go",
	},
	{
		Tag:            "1.8.2",
		Implementation: "go",
	},
	{
		Tag:            "1.9",
		Implementation: "go",
	},
	{
		Tag:            "1.10",
		Implementation: "go",
	},
	{
		Tag:            "1.11.1",
		Implementation: "go",
	},
	{
		Tag:            "1.13.5",
		Implementation: "go",
	},
}

func TestFinderMock(t *testing.T) {
	testCases := []struct {
		Description       string
		AvailableVersions []goenv.Version
		TagToFind         goenv.Tag
		ExpectedVersions  []goenv.Version
	}{
		{
			"find version with no available versions",
			[]goenv.Version{},
			"1.13",
			[]goenv.Version{},
		},
		{
			"try to find an unknown version",
			versions,
			"1.7",
			[]goenv.Version{},
		},
		{
			"find all versions tagged with 1",
			versions,
			"1",
			versions,
		},
		{
			"find a specific versions tagged with 1.13.5",
			versions,
			"1",
			versions,
		},
		{
			"find all versions tagged with 1.8",
			versions,
			"1.8",
			[]goenv.Version{
				{
					Tag:            "1.8",
					Implementation: "go",
				},
				{
					Tag:            "1.8.1",
					Implementation: "go",
				},
				{
					Tag:            "1.8.2",
					Implementation: "go",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			versionFinder := mock.NewVersionFinder(tc.AvailableVersions)
			actualVersions := versionFinder.Find(tc.TagToFind)
			SortedCompare(t, actualVersions, tc.ExpectedVersions)
		})
	}
}
