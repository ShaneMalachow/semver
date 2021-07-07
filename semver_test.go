package semver_test

import (
	"testing"

	"github.com/ShaneMalachow/semver"
)

func TestParseSemver(t *testing.T) {
	for _, testVer := range ValidSemvers {
		ver, err := semver.ParseSemVer(testVer.String)
		if err != nil {
			t.Errorf("Unable to parse valid test case: %s\n%s\n", testVer.String, err.Error())
		}
		if *ver != testVer.SemVer {
			t.Errorf("Result does not match test case\nresult:    %#v\ntest case:  %#v", ver, testVer.SemVer)
		}
	}
	for _, testVer := range InvalidSemvers {
		_, err := semver.ParseSemVer(testVer)
		if err == nil {
			t.Errorf("Failed to error on invalid test case: %s\n", testVer)
		}
	}
}

func TestVersion(t *testing.T) {
	for _, test := range ValidSemvers {
		s := test.SemVer.Version()
		if s != test.String {
			t.Errorf("Version did not print correctly:\nOutput: %s\nExpected: %s\n", s, test.String)
		}
	}
}

func TestCore(t *testing.T) {
	for _, test := range ValidSemvers {
		s := test.SemVer.Core()
		if s != test.Core {
			t.Errorf("Core did not print correctly:\nOutput: %s\nExpected: %s\n", s, test.Core)
		}
	}
}

func TestCompare(t *testing.T) {
	for _, test := range compareTests {
		result := semver.Compare(test.A, test.B)
		if result != test.Result {
			t.Errorf("Comparison failed:\n\tA: %#v\n\tB: %#v\n\tResult vs Expected: %d != %d", test.A, test.B, result, test.Result)
		}
	}
}

type SemverCompareTest struct {
	A, B   semver.SemVer
	Result int
}

var compareTests = []SemverCompareTest{
	{
		A:      semver.SemVer{Major: "1", Minor: "2", Patch: "3", Prerelease: "", BuildMetadata: ""},
		B:      semver.SemVer{Major: "1", Minor: "2", Patch: "3", Prerelease: "", BuildMetadata: ""},
		Result: 0,
	},
	{
		A:      semver.SemVer{Major: "1", Minor: "2", Patch: "4", Prerelease: "", BuildMetadata: ""},
		B:      semver.SemVer{Major: "1", Minor: "2", Patch: "3", Prerelease: "", BuildMetadata: ""},
		Result: 1,
	},
	{
		A:      semver.SemVer{Major: "1", Minor: "2", Patch: "3", Prerelease: "", BuildMetadata: ""},
		B:      semver.SemVer{Major: "1", Minor: "2", Patch: "4", Prerelease: "", BuildMetadata: ""},
		Result: -1,
	},
	{
		A:      semver.SemVer{Major: "1", Minor: "3", Patch: "3", Prerelease: "", BuildMetadata: ""},
		B:      semver.SemVer{Major: "1", Minor: "2", Patch: "3", Prerelease: "", BuildMetadata: ""},
		Result: 1,
	},
	{
		A:      semver.SemVer{Major: "1", Minor: "2", Patch: "3", Prerelease: "", BuildMetadata: ""},
		B:      semver.SemVer{Major: "1", Minor: "3", Patch: "3", Prerelease: "", BuildMetadata: ""},
		Result: -1,
	},
	{
		A:      semver.SemVer{Major: "2", Minor: "2", Patch: "3", Prerelease: "", BuildMetadata: ""},
		B:      semver.SemVer{Major: "1", Minor: "2", Patch: "3", Prerelease: "", BuildMetadata: ""},
		Result: 1,
	},
	{
		A:      semver.SemVer{Major: "1", Minor: "2", Patch: "3", Prerelease: "", BuildMetadata: ""},
		B:      semver.SemVer{Major: "2", Minor: "2", Patch: "3", Prerelease: "", BuildMetadata: ""},
		Result: -1,
	},
	{
		A:      semver.SemVer{Major: "1", Minor: "2", Patch: "3", Prerelease: "", BuildMetadata: ""},
		B:      semver.SemVer{Major: "2", Minor: "2", Patch: "3", Prerelease: "alpha", BuildMetadata: ""},
		Result: -1,
	},
	{
		A:      semver.SemVer{Major: "1", Minor: "2", Patch: "3", Prerelease: "alpha", BuildMetadata: ""},
		B:      semver.SemVer{Major: "2", Minor: "2", Patch: "3", Prerelease: "", BuildMetadata: ""},
		Result: -1,
	},
	{
		A:      semver.SemVer{Major: "1", Minor: "2", Patch: "4", Prerelease: "", BuildMetadata: ""},
		B:      semver.SemVer{Major: "1", Minor: "2", Patch: "3", Prerelease: "alpha", BuildMetadata: ""},
		Result: 1,
	},
}

type semverTestCase struct {
	String string
	Core   string
	SemVer semver.SemVer
}

var ValidSemvers = []semverTestCase{
	{String: "0.0.4", SemVer: semver.SemVer{Major: "0", Minor: "0", Patch: "4", Prerelease: "", BuildMetadata: ""}, Core: "0.0.4"},
	{String: "1.2.3", SemVer: semver.SemVer{Major: "1", Minor: "2", Patch: "3", Prerelease: "", BuildMetadata: ""}, Core: "1.2.3"},
	{String: "10.20.30", SemVer: semver.SemVer{Major: "10", Minor: "20", Patch: "30", Prerelease: "", BuildMetadata: ""}, Core: "10.20.30"},
	{String: "1.1.2-prerelease+meta", SemVer: semver.SemVer{Major: "1", Minor: "1", Patch: "2", Prerelease: "prerelease", BuildMetadata: "meta"}, Core: "1.1.2"},
	{String: "1.1.2+meta", SemVer: semver.SemVer{Major: "1", Minor: "1", Patch: "2", Prerelease: "", BuildMetadata: "meta"}, Core: "1.1.2"},
	{String: "1.1.2+meta-valid", SemVer: semver.SemVer{Major: "1", Minor: "1", Patch: "2", Prerelease: "", BuildMetadata: "meta-valid"}, Core: "1.1.2"},
	{String: "1.0.0-alpha", SemVer: semver.SemVer{Major: "1", Minor: "0", Patch: "0", Prerelease: "alpha", BuildMetadata: ""}, Core: "1.0.0"},
	{String: "1.0.0-beta", SemVer: semver.SemVer{Major: "1", Minor: "0", Patch: "0", Prerelease: "beta", BuildMetadata: ""}, Core: "1.0.0"},
	{String: "1.0.0-alpha.beta", SemVer: semver.SemVer{Major: "1", Minor: "0", Patch: "0", Prerelease: "alpha.beta", BuildMetadata: ""}, Core: "1.0.0"},
	{String: "1.0.0-alpha.beta.1", SemVer: semver.SemVer{Major: "1", Minor: "0", Patch: "0", Prerelease: "alpha.beta.1", BuildMetadata: ""}, Core: "1.0.0"},
	{String: "1.0.0-alpha.1", SemVer: semver.SemVer{Major: "1", Minor: "0", Patch: "0", Prerelease: "alpha.1", BuildMetadata: ""}, Core: "1.0.0"},
	{String: "1.0.0-alpha0.valid", SemVer: semver.SemVer{Major: "1", Minor: "0", Patch: "0", Prerelease: "alpha0.valid", BuildMetadata: ""}, Core: "1.0.0"},
	{String: "1.0.0-alpha.0valid", SemVer: semver.SemVer{Major: "1", Minor: "0", Patch: "0", Prerelease: "alpha.0valid", BuildMetadata: ""}, Core: "1.0.0"},
	{String: "1.0.0-alpha-a.b-c-somethinglong+build.1-aef.1-its-okay", SemVer: semver.SemVer{Major: "1", Minor: "0", Patch: "0", Prerelease: "alpha-a.b-c-somethinglong", BuildMetadata: "build.1-aef.1-its-okay"}, Core: "1.0.0"},
	{String: "1.0.0-rc.1+build.1", SemVer: semver.SemVer{Major: "1", Minor: "0", Patch: "0", Prerelease: "rc.1", BuildMetadata: "build.1"}, Core: "1.0.0"},
	{String: "2.0.0-rc.1+build.123", SemVer: semver.SemVer{Major: "2", Minor: "0", Patch: "0", Prerelease: "rc.1", BuildMetadata: "build.123"}, Core: "2.0.0"},
	{String: "1.2.3-beta", SemVer: semver.SemVer{Major: "1", Minor: "2", Patch: "3", Prerelease: "beta", BuildMetadata: ""}, Core: "1.2.3"},
	{String: "1.2.3-DEV-SNAPSHOT", SemVer: semver.SemVer{Major: "1", Minor: "2", Patch: "3", Prerelease: "DEV-SNAPSHOT", BuildMetadata: ""}, Core: "1.2.3"},
	{String: "1.2.3-SNAPSHOT-123", SemVer: semver.SemVer{Major: "1", Minor: "2", Patch: "3", Prerelease: "SNAPSHOT-123", BuildMetadata: ""}, Core: "1.2.3"},
	{String: "1.0.0", SemVer: semver.SemVer{Major: "1", Minor: "0", Patch: "0", Prerelease: "", BuildMetadata: ""}, Core: "1.0.0"},
	{String: "2.0.0", SemVer: semver.SemVer{Major: "2", Minor: "0", Patch: "0", Prerelease: "", BuildMetadata: ""}, Core: "2.0.0"},
	{String: "1.1.7", SemVer: semver.SemVer{Major: "1", Minor: "1", Patch: "7", Prerelease: "", BuildMetadata: ""}, Core: "1.1.7"},
	{String: "2.0.0+build.1848", SemVer: semver.SemVer{Major: "2", Minor: "0", Patch: "0", Prerelease: "", BuildMetadata: "build.1848"}, Core: "2.0.0"},
	{String: "2.0.1-alpha.1227", SemVer: semver.SemVer{Major: "2", Minor: "0", Patch: "1", Prerelease: "alpha.1227", BuildMetadata: ""}, Core: "2.0.1"},
	{String: "1.0.0-alpha+beta", SemVer: semver.SemVer{Major: "1", Minor: "0", Patch: "0", Prerelease: "alpha", BuildMetadata: "beta"}, Core: "1.0.0"},
	{String: "1.2.3----RC-SNAPSHOT.12.9.1--.12+788", SemVer: semver.SemVer{Major: "1", Minor: "2", Patch: "3", Prerelease: "---RC-SNAPSHOT.12.9.1--.12", BuildMetadata: "788"}, Core: "1.2.3"},
	{String: "1.2.3----R-S.12.9.1--.12+meta", SemVer: semver.SemVer{Major: "1", Minor: "2", Patch: "3", Prerelease: "---R-S.12.9.1--.12", BuildMetadata: "meta"}, Core: "1.2.3"},
	{String: "1.2.3----RC-SNAPSHOT.12.9.1--.12", SemVer: semver.SemVer{Major: "1", Minor: "2", Patch: "3", Prerelease: "---RC-SNAPSHOT.12.9.1--.12", BuildMetadata: ""}, Core: "1.2.3"},
	{String: "1.0.0+0.build.1-rc.10000aaa-kk-0.1", SemVer: semver.SemVer{Major: "1", Minor: "0", Patch: "0", Prerelease: "", BuildMetadata: "0.build.1-rc.10000aaa-kk-0.1"}, Core: "1.0.0"},
	{String: "99999999999999999999999.999999999999999999.99999999999999999", SemVer: semver.SemVer{Major: "99999999999999999999999", Minor: "999999999999999999", Patch: "99999999999999999", Prerelease: "", BuildMetadata: ""}, Core: "99999999999999999999999.999999999999999999.99999999999999999"},
	{String: "1.0.0-0A.is.legal", SemVer: semver.SemVer{Major: "1", Minor: "0", Patch: "0", Prerelease: "0A.is.legal", BuildMetadata: ""}, Core: "1.0.0"},
}

var InvalidSemvers = []string{
	"1",
	"1.2",
	"1.2.3-0123",
	"1.2.3-0123.0123",
	"1.1.2+.123",
	"+invalid",
	"-invalid",
	"-invalid+invalid",
	"-invalid.01",
	"alpha",
	"alpha.beta",
	"alpha.beta.1",
	"alpha.1",
	"alpha+beta",
	"alpha_beta",
	"alpha.",
	"alpha..",
	"beta",
	"1.0.0-alpha_beta",
	"-alpha.",
	"1.0.0-alpha..",
	"1.0.0-alpha..1",
	"1.0.0-alpha...1",
	"1.0.0-alpha....1",
	"1.0.0-alpha.....1",
	"1.0.0-alpha......1",
	"1.0.0-alpha.......1",
	"01.1.1",
	"1.01.1",
	"1.1.01",
	"1.2",
	"1.2.3.DEV",
	"1.2-SNAPSHOT",
	"1.2.31.2.3----RC-SNAPSHOT.12.09.1--..12+788",
	"1.2-RC-SNAPSHOT",
	"-1.0.3-gamma+b7718",
	"+justmeta",
	"9.8.7+meta+meta",
	"9.8.7-whatever+meta+meta",
	"99999999999999999999999.999999999999999999.99999999999999999----RC-SNAPSHOT.12.09.1--------------------------------..12",
}
