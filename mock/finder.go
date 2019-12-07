package mock

import (
	"strings"

	"github.com/seppone/goenv"
)

type VersionFinder struct {
	versions []goenv.Version
}

func NewVersionFinder(versions []goenv.Version) goenv.VersionFinder {
	return &VersionFinder{versions}
}

func (vf VersionFinder) Find(tag goenv.Tag) []goenv.Version {
	var foundVersions []goenv.Version
	for _, version := range vf.versions {
		if strings.Contains(version.String(), string(tag)) {
			foundVersions = append(foundVersions, version)
		}
	}
	return foundVersions
}
