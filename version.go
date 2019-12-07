package goenv

type Version struct {
	Tag            string
	Implementation string
}

func (v Version) String() string {
	return v.Implementation + "-" + v.Tag
}

type Tag string

type VersionFinder interface {
	Find(Tag) []Version
}

type VersionInstaller interface {
	Install(Version) error
}

type VersionChooser interface {
	Determine() (Version, error)
	Choose(Version) error
}
