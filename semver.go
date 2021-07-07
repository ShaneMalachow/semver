package semver

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func ParseSemVer(s string) (*SemVer, error) {
	semverRegex := regexp.MustCompile(`^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)
	if semverRegex.MatchString(s) {
		matches := semverRegex.FindStringSubmatch(s)
		major := matches[semverRegex.SubexpIndex("major")]
		minor := matches[semverRegex.SubexpIndex("minor")]
		patch := matches[semverRegex.SubexpIndex("patch")]
		prerelease := matches[semverRegex.SubexpIndex("prerelease")]
		buildmetadata := matches[semverRegex.SubexpIndex("buildmetadata")]

		return &SemVer{
			Major:         major,
			Minor:         minor,
			Patch:         patch,
			Prerelease:    prerelease,
			BuildMetadata: buildmetadata,
		}, nil
	}
	return &SemVer{}, errors.New("string does not match semantic version syntax")
}

type SemVer struct {
	Major         string
	Minor         string
	Patch         string
	Prerelease    string
	BuildMetadata string
}

func (s SemVer) Core() string {
	return fmt.Sprintf("%s.%s.%s", s.Major, s.Minor, s.Patch)
}

func (s SemVer) Version() string {
	version := s.Core()
	if s.Prerelease != "" {
		version = fmt.Sprintf("%s-%s", version, s.Prerelease)
	}
	if s.BuildMetadata != "" {
		version = fmt.Sprintf("%s+%s", version, s.BuildMetadata)
	}
	return version
}

func Compare(s1, s2 SemVer) int {
	if s1.Major != s2.Major {
		return strings.Compare(s1.Major, s2.Major)
	} else if s1.Minor != s2.Minor {
		return strings.Compare(s1.Minor, s2.Minor)
	} else if s1.Patch != s2.Patch {
		return strings.Compare(s1.Patch, s2.Patch)
	} else {
		if s1.Prerelease == "" && s2.Prerelease != "" {
			return 1
		} else if s1.Prerelease != "" && s2.Prerelease == "" {
			return -1
		} else {
			return strings.Compare(s2.Prerelease, s1.Prerelease)
		}
	}
}
